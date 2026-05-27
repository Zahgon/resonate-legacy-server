package sqlite

import (
	"database/sql"
	"math/rand" // nosemgrep
	"time"

	"github.com/go-viper/mapstructure/v2"
	"github.com/resonatehq/resonate/internal/aio"
	"github.com/resonatehq/resonate/internal/kernel/bus"
	"github.com/resonatehq/resonate/internal/kernel/t_aio"
	"github.com/resonatehq/resonate/internal/metrics"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	_ "github.com/mattn/go-sqlite3"
)

const (
	// Schema is now managed by migrations (see internal/migrationfiles/migrations/sqlite/)
	// We create the migrations table to handle
	CREATE_TABLE_STATEMENT = `
	CREATE TABLE IF NOT EXISTS migrations (
		id INTEGER PRIMARY KEY
	);`

	PROMISE_SELECT_STATEMENT = `
	SELECT
		id, state, param_headers, param_data, value_headers, value_data, timeout, tags, created_on, completed_on
	FROM
		promises
	WHERE
		id = ?`

	PROMISE_SELECT_ALL_STATEMENT = `
	SELECT
		id, state, param_headers, param_data, value_headers, value_data, timeout, tags, created_on, completed_on, sort_id
	FROM
		promises
	WHERE
		state = 1 AND timeout <= ?
	LIMIT
		?`

	PROMISE_SEARCH_STATEMENT = `
	SELECT
		id, state, param_headers, param_data, value_headers, value_data, timeout, tags, created_on, completed_on, sort_id
	FROM
		promises
	WHERE
		(? IS NULL OR sort_id < ?) AND
		id LIKE ? AND
		state & ? != 0
		%s
	ORDER BY
		sort_id DESC
	LIMIT
		?`

	PROMISE_INSERT_STATEMENT = `
	INSERT INTO promises
		(id, state, param_headers, param_data, timeout, idempotency_key_for_create, tags, created_on)
	VALUES
		(?, ?, ?, ?, ?, ?, ?, ?)
	ON CONFLICT(id) DO NOTHING -- idempotency_key must be equal to id for this insert`

	PROMISE_UPDATE_STATEMENT = `
	UPDATE
		promises
	SET
		state = ?, value_headers = ?, value_data = ?, idempotency_key_for_complete = ?, completed_on = ?
	WHERE
		id = ? AND state = 1 -- idempotency_key must be equal to id for this insert`

	CALLBACK_INSERT_STATEMENT = `
	INSERT INTO callbacks
		(id, promise_id, root_promise_id, recv, mesg, timeout, created_on)
	SELECT
		?, ?, ?, ?, ?, ?, ?
	WHERE EXISTS
		(SELECT 1 FROM promises WHERE id = ? AND state = 1)
	AND NOT EXISTS
		(SELECT 1 FROM callbacks WHERE id = ?)`

	CALLBACK_DELETE_STATEMENT = `
	DELETE FROM callbacks WHERE promise_id = ?`

	SCHEDULE_SELECT_STATEMENT = `
	SELECT
		id, description, cron, tags, promise_id, promise_timeout, promise_param_headers, promise_param_data, promise_tags, last_run_time, next_run_time, created_on
	FROM
		schedules
	WHERE
		id = ?`

	SCHEDULE_SELECT_ALL_STATEMENT = `
	SELECT
		id, cron, promise_id, promise_timeout, promise_param_headers, promise_param_data, promise_tags, last_run_time, next_run_time
	FROM
		schedules
	WHERE
		next_run_time <= ?
	ORDER BY
		next_run_time ASC, sort_id ASC
	LIMIT
		?`

	SCHEDULE_SEARCH_STATEMENT = `
	SELECT
		id, cron, tags, last_run_time, next_run_time, created_on, sort_id
	FROM
		schedules
	WHERE
		(? IS NULL OR sort_id < ?) AND
		id LIKE ?
		%s
	ORDER BY
		sort_id DESC
	LIMIT
		?`

	SCHEDULE_INSERT_STATEMENT = `
	INSERT INTO schedules
		(id, description, cron, tags, promise_id, promise_timeout, promise_param_headers, promise_param_data, promise_tags, next_run_time, idempotency_key, created_on)
	VALUES
		(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	ON CONFLICT(id) DO NOTHING -- idempotency_key must be equal to id for this stmt`

	SCHEDULE_UPDATE_STATEMENT = `
	UPDATE
		schedules
	SET
		last_run_time = next_run_time, next_run_time = ?
	WHERE
		id = ? AND next_run_time = ?`

	SCHEDULE_DELETE_STATEMENT = `
	DELETE FROM schedules WHERE id = ?`

	TASK_SELECT_STATEMENT = `
	SELECT
		id, process_id, state, root_promise_id, recv, mesg, timeout, counter, attempt, ttl, expires_at, created_on, completed_on
	FROM
		tasks
	WHERE
		id = ?`

	TASK_VALIDATE_STATEMENT = `
	SELECT
		COUNT(*)
	FROM
		tasks
	WHERE
		id = ? AND counter  = ? AND state = 4 -- state == Claimed`

	TASK_SELECT_ALL_STATEMENT = `
	SELECT
		id,
		process_id,
		state,
		root_promise_id,
		recv,
		mesg,
		timeout,
		counter,
		attempt,
		ttl,
		expires_at,
		created_on,
		completed_on
	FROM tasks
	WHERE
		state & ? != 0 AND ((expires_at != 0 AND expires_at <= ?) OR timeout <= ?)
	ORDER BY root_promise_id, sort_id ASC
	LIMIT ?`

	TASK_SELECT_ENQUEUEABLE_STATEMENT = `
	SELECT
		id,
		process_id,
		state,
		root_promise_id,
		recv,
		mesg,
		timeout,
		counter,
		attempt,
		ttl,
		expires_at,
		created_on,
		completed_on,
		min(sort_id)
	FROM tasks t1
	WHERE
		state = 1 AND expires_at <= ? -- State = 1 -> Init
	AND NOT EXISTS (
		SELECT 1
		FROM tasks t2
		WHERE t2.root_promise_id = t1.root_promise_id
		AND t2.state in (2, 4) -- 2 -> Enqueue, 4 -> Claimed
	)
	GROUP BY root_promise_id
	ORDER BY root_promise_id, sort_id ASC
	LIMIT ?`

	TASK_INSERT_STATEMENT = `
	INSERT INTO tasks
		(id, recv, mesg, timeout, process_id, state, root_promise_id, ttl, expires_at, created_on)
	VALUES
		(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	ON CONFLICT(id) DO NOTHING`

	TASK_INSERT_ALL_STATEMENT = `
	INSERT INTO tasks
		(id, recv, mesg, timeout, root_promise_id, created_on)
	SELECT
		id, recv, mesg, timeout, root_promise_id, ?
	FROM
		callbacks
	WHERE
		promise_id = ?
	ORDER BY
		id`

	TASK_UPDATE_STATEMENT = `
	UPDATE
		tasks
	SET
		process_id = ?, state = ?, counter = ?, attempt = ?, ttl = ?, expires_at = ?, completed_on = ?
	WHERE
		id = ? AND state & ? != 0 AND counter = ?`

	TASK_COMPLETE_BY_ROOT_ID_STATEMENT = `
	UPDATE
		tasks
	SET
		state = 8, completed_on = ? -- State 8 -> Completed
	WHERE
		root_promise_id = ? AND state in (1, 2, 4) -- State in (Init, Enqueued, Claimed)`

	TASK_HEARTBEAT_STATEMENT = `
	UPDATE
		tasks
	SET
		expires_at = MIN(? + ttl, 9223372036854775807) -- max int64
	WHERE
		process_id = ? AND state = 4`
)

// Config

type Config struct {
	Size      int           `flag:"size" desc:"submission buffered channel size" default:"1000"`
	BatchSize int           `flag:"batch-size" desc:"max submissions processed per iteration" default:"1000"`
	Path      string        `flag:"path" desc:"sqlite database path" default:"resonate.db" dst:":memory:" dev:":memory:"`
	TxTimeout time.Duration `flag:"tx-timeout" desc:"sqlite transaction timeout" default:"10s"`
	Reset     bool          `flag:"reset" desc:"reset sqlite db on shutdown" default:"false" dst:"true"`
}

func (c *Config) Bind(cmd *cobra.Command, flg *pflag.FlagSet, vip *viper.Viper, name string, prefix string, keyPrefix string) {
	_ = "STUB: not implemented"
	return
}

func (c *Config) Decode(value any, decodeHook mapstructure.DecodeHookFunc) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Config) New(aio aio.AIO, metrics *metrics.Metrics) (aio.Subsystem, error) {
	_ = "STUB: not implemented"
	return *new(aio.Subsystem), nil
}

func (c *Config) NewDST(aio aio.AIO, metrics *metrics.Metrics, _ *rand.Rand, _ chan any) (aio.SubsystemDST, error) {
	_ = "STUB: not implemented"
	return *new(aio.SubsystemDST), nil
}

// Subsystem

type SqliteStore struct {
	config *Config
	sq     chan<- *bus.SQE[t_aio.Submission, t_aio.Completion]
	db     *sql.DB
	worker *SqliteStoreWorker
}

func (s *SqliteStore) DB() *sql.DB { _ = "STUB: not implemented"; return nil }

func NewConn(path string) (*sql.DB, error) { _ = "STUB: not implemented"; return nil, nil }

func New(aio aio.AIO, metrics *metrics.Metrics, config *Config) (*SqliteStore, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SqliteStore) String() string { _ = "STUB: not implemented"; return "" }

func (s *SqliteStore) Kind() t_aio.Kind { _ = "STUB: not implemented"; return *new(t_aio.Kind) }

func (s *SqliteStore) Start(chan<- error) error { _ = "STUB: not implemented"; return nil }

// Get pending migrations

// If version == 0, the db is fresh and we can apply all migrations automatically

// Validate migration sequence

// Apply all migrations

// For existing databases, check for pending migrations and error if any exist

// start worker on a goroutine

func (s *SqliteStore) Stop() error { _ = "STUB: not implemented"; return nil }

func (s *SqliteStore) Reset() error { _ = "STUB: not implemented"; return nil }

func (s *SqliteStore) Enqueue(sqe *bus.SQE[t_aio.Submission, t_aio.Completion]) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *SqliteStore) Flush(t int64) { _ = "STUB: not implemented"; return }

func (s *SqliteStore) Process(sqes []*bus.SQE[t_aio.Submission, t_aio.Completion]) []*bus.CQE[t_aio.Submission, t_aio.Completion] {
	_ = "STUB: not implemented"
	return nil
}

// Worker

type SqliteStoreWorker struct {
	config  *Config
	db      *sql.DB
	sq      <-chan *bus.SQE[t_aio.Submission, t_aio.Completion]
	flush   chan int64
	aio     aio.AIO
	metrics *metrics.Metrics
}

func (w *SqliteStoreWorker) String() string { _ = "STUB: not implemented"; return "" }

func (w *SqliteStoreWorker) Start() { _ = "STUB: not implemented"; return }

func (w *SqliteStoreWorker) Flush(t int64) {
	_ = "STUB: not implemented"
	// ignore case where flush channel is full,
	// this means the flush is waiting on the cq
	return
}

func (w *SqliteStoreWorker) Process(sqes []*bus.SQE[t_aio.Submission, t_aio.Completion]) []*bus.CQE[t_aio.Submission, t_aio.Completion] {
	_ = "STUB: not implemented"
	return nil
}

func (w *SqliteStoreWorker) Execute(transactions []*t_aio.Transaction) ([]*t_aio.StoreCompletion, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *SqliteStoreWorker) performCommands(tx *sql.Tx, transactions []*t_aio.Transaction) ([]*t_aio.StoreCompletion, error) {
	_ = "STUB: not implemented"
	// lazily instantiate prepared statements
	return nil, nil
}

// Results

// Promises

// Callbacks

// Schedules

// Tasks

// Promises

func (w *SqliteStoreWorker) readPromise(tx *sql.Tx, cmd *t_aio.ReadPromiseCommand) (*t_aio.QueryPromisesResult, error) {
	_ = "STUB: not implemented"
	// select
	return nil, nil
}

func (w *SqliteStoreWorker) readPromises(tx *sql.Tx, cmd *t_aio.ReadPromisesCommand) (*t_aio.QueryPromisesResult, error) {
	_ = "STUB: not implemented"
	// select
	return nil, nil
}

func (w *SqliteStoreWorker) searchPromises(tx *sql.Tx, cmd *t_aio.SearchPromisesCommand) (*t_aio.QueryPromisesResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// convert query

// convert list of state to bit mask

// tags

// Dynamic placeholders for tags.

// select

func (w *SqliteStoreWorker) createPromise(tx *sql.Tx, stmt *sql.Stmt, cmd *t_aio.CreatePromiseCommand) (*t_aio.AlterPromisesResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *SqliteStoreWorker) createPromiseAndTask(tx *sql.Tx, promiseStmt *sql.Stmt, taskStmt *sql.Stmt, cmd *t_aio.CreatePromiseAndTaskCommand) (*t_aio.AlterPromisesResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *SqliteStoreWorker) updatePromise(tx *sql.Tx, stmt *sql.Stmt, cmd *t_aio.UpdatePromiseCommand) (*t_aio.AlterPromisesResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// update

// Callbacks

func (w *SqliteStoreWorker) createCallback(tx *sql.Tx, stmt *sql.Stmt, cmd *t_aio.CreateCallbackCommand) (*t_aio.AlterCallbacksResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *SqliteStoreWorker) deleteCallbacks(tx *sql.Tx, stmt *sql.Stmt, cmd *t_aio.DeleteCallbacksCommand) (*t_aio.AlterCallbacksResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Schedules

func (w *SqliteStoreWorker) readSchedule(tx *sql.Tx, cmd *t_aio.ReadScheduleCommand) (*t_aio.QuerySchedulesResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *SqliteStoreWorker) readSchedules(tx *sql.Tx, cmd *t_aio.ReadSchedulesCommand) (*t_aio.QuerySchedulesResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *SqliteStoreWorker) searchSchedules(tx *sql.Tx, cmd *t_aio.SearchSchedulesCommand) (*t_aio.QuerySchedulesResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// convert query

// tags

// select

func (w *SqliteStoreWorker) createSchedule(tx *sql.Tx, stmt *sql.Stmt, cmd *t_aio.CreateScheduleCommand) (*t_aio.AlterSchedulesResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *SqliteStoreWorker) updateSchedule(tx *sql.Tx, stmt *sql.Stmt, cmd *t_aio.UpdateScheduleCommand) (*t_aio.AlterSchedulesResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *SqliteStoreWorker) deleteSchedule(tx *sql.Tx, stmt *sql.Stmt, cmd *t_aio.DeleteScheduleCommand) (*t_aio.AlterSchedulesResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Tasks

func (w *SqliteStoreWorker) readTask(tx *sql.Tx, cmd *t_aio.ReadTaskCommand) (*t_aio.QueryTasksResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *SqliteStoreWorker) readTasks(tx *sql.Tx, cmd *t_aio.ReadTasksCommand) (*t_aio.QueryTasksResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *SqliteStoreWorker) readEnqueueableTasks(tx *sql.Tx, cmd *t_aio.ReadEnqueueableTasksCommand) (*t_aio.QueryTasksResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *SqliteStoreWorker) createTask(tx *sql.Tx, stmt *sql.Stmt, cmd *t_aio.CreateTaskCommand) (*t_aio.AlterTasksResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *SqliteStoreWorker) createTasks(tx *sql.Tx, stmt *sql.Stmt, cmd *t_aio.CreateTasksCommand) (*t_aio.AlterTasksResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *SqliteStoreWorker) completeTasks(tx *sql.Tx, stmt *sql.Stmt, cmd *t_aio.CompleteTasksCommand) (*t_aio.AlterTasksResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *SqliteStoreWorker) updateTask(tx *sql.Tx, stmt *sql.Stmt, cmd *t_aio.UpdateTaskCommand) (*t_aio.AlterTasksResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *SqliteStoreWorker) heartbeatTasks(tx *sql.Tx, stmt *sql.Stmt, cmd *t_aio.HeartbeatTasksCommand) (*t_aio.AlterTasksResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *SqliteStoreWorker) validFencingToken(tx *sql.Tx, transaction *t_aio.Transaction) (bool, error) {
	_ = "STUB: not implemented"
	// if the task is not provided continue with the operation
	return false, nil
}
