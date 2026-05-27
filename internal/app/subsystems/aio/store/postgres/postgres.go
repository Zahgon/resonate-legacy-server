package postgres

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

	_ "github.com/lib/pq"
)

const (
	// Schema is now managed by migrations (see internal/migrationfiles/migrations/postgres/)
	CREATE_TABLE_STATEMENT = `
	CREATE TABLE IF NOT EXISTS migrations (
		id INTEGER,
		PRIMARY KEY(id)
	);`

	DROP_TABLE_STATEMENT = `
	DROP TABLE promises;
	DROP TABLE callbacks;
	DROP TABLE schedules;
	DROP TABLE locks;
	DROP TABLE tasks;
	DROP TABLE migrations;`

	PROMISE_SELECT_STATEMENT = `
	SELECT
		id, state, param_headers, param_data, value_headers, value_data, timeout, tags, created_on, completed_on
	FROM
		promises
	WHERE
		id = $1`

	PROMISE_SELECT_ALL_STATEMENT = `
	SELECT
		id, state, param_headers, param_data, value_headers, value_data, timeout, tags, created_on, completed_on, sort_id
	FROM
		promises
	WHERE
		state = 1 AND timeout <= $1
	LIMIT
		$2`

	PROMISE_SEARCH_STATEMENT = `
	SELECT
		id, state, param_headers, param_data, value_headers, value_data, timeout, tags, created_on, completed_on, sort_id
	FROM
		promises
	WHERE
		($1::int IS NULL OR sort_id < $1) AND
		id LIKE $2 AND
		state & $3 != 0 AND
		($4::jsonb IS NULL OR tags @> $4)
	ORDER BY
		sort_id DESC
	LIMIT
		$5`

	PROMISE_INSERT_STATEMENT = `
	INSERT INTO promises
		(id, state, param_headers, param_data, timeout, idempotency_key_for_create, tags, created_on)
	VALUES
		($1, $2, $3, $4, $5, $6, $7, $8)
	ON CONFLICT(id) DO NOTHING -- idempotency_key must be equal to id for this stmt`

	PROMISE_UPDATE_STATEMENT = `
	UPDATE
		promises
	SET
		state = $1, value_headers = $2, value_data = $3, idempotency_key_for_complete = $4, completed_on = $5
	WHERE
		id = $6 AND state = 1 -- idempotency_key must be equal to id for this stmt`

	CALLBACK_INSERT_STATEMENT = `
	INSERT INTO callbacks
		(id, promise_id, root_promise_id, recv, mesg, timeout, created_on)
	SELECT
		$1, $2, $3, $4, $5, $6, $7
	WHERE EXISTS
		(SELECT 1 FROM promises WHERE id = $2 AND state = 1)
	AND NOT EXISTS
		(SELECT 1 FROM callbacks WHERE id = $1)`

	CALLBACK_DELETE_STATEMENT = `
	DELETE FROM callbacks WHERE promise_id = $1`

	SCHEDULE_SELECT_STATEMENT = `
	SELECT
		id, description, cron, tags, promise_id, promise_timeout, promise_param_headers, promise_param_data, promise_tags, last_run_time, next_run_time, created_on
	FROM
		schedules
	WHERE
		id = $1`

	SCHEDULE_SELECT_ALL_STATEMENT = `
	SELECT
		id, cron, promise_id, promise_timeout, promise_param_headers, promise_param_data, promise_tags, last_run_time, next_run_time
	FROM
		schedules
	WHERE
		next_run_time <= $1
	ORDER BY
		next_run_time ASC, sort_id ASC
	LIMIT
		$2`

	SCHEDULE_SEARCH_STATEMENT = `
	SELECT
		id, cron, tags, last_run_time, next_run_time, created_on, sort_id
	FROM
		schedules
	WHERE
		($1::int IS NULL OR sort_id < $1) AND
		id LIKE $2 AND
		($3::jsonb IS NULL OR tags @> $3)
	ORDER BY
		sort_id DESC
	LIMIT
		$4`

	SCHEDULE_INSERT_STATEMENT = `
	INSERT INTO schedules
		(id, description, cron, tags, promise_id, promise_timeout, promise_param_headers, promise_param_data, promise_tags, next_run_time, idempotency_key, created_on)
	VALUES
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	ON CONFLICT(id) DO NOTHING -- idempotency_key must be equal to id for this stmt`

	SCHEDULE_UPDATE_STATEMENT = `
	UPDATE
		schedules
	SET
		last_run_time = next_run_time, next_run_time = $1
	WHERE
		id = $2 AND next_run_time = $3`

	SCHEDULE_DELETE_STATEMENT = `
	DELETE FROM schedules WHERE id = $1`

	TASK_SELECT_STATEMENT = `
	SELECT
		id, process_id, state, root_promise_id, recv, mesg, timeout, counter, attempt, ttl, expires_at, created_on, completed_on
	FROM
		tasks
	WHERE
		id = $1`

	TASK_VALIDATE_STATEMENT = `
	SELECT
		COUNT(id)
	FROM
		tasks
	WHERE
		id = $1 AND counter = $2 AND state = 4 -- state == Claimed`

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
		state & $1 != 0 AND ((expires_at != 0 AND expires_at <= $2) OR timeout <= $2)
	ORDER BY root_promise_id, sort_id ASC
	LIMIT $3`

	TASK_SELECT_ENQUEUEABLE_STATEMENT = `
	SELECT DISTINCT ON (root_promise_id)
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
	FROM tasks t1
	WHERE
		state = 1 AND expires_at <= $1 -- State = 1 -> Init
	AND NOT EXISTS (
		SELECT 1
		FROM tasks t2
		WHERE t2.root_promise_id = t1.root_promise_id
		AND t2.state in (2, 4) -- 2 -> Enqueue, 4 -> Claimed
	)
	ORDER BY root_promise_id, sort_id ASC
	LIMIT $2`

	TASK_INSERT_STATEMENT = `
	INSERT INTO tasks
		(id, recv, mesg, timeout, process_id, state, root_promise_id, ttl, expires_at, created_on)
	VALUES
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	ON CONFLICT(id) DO NOTHING`

	TASK_INSERT_ALL_STATEMENT = `
	INSERT INTO tasks
		(id, recv, mesg, timeout, root_promise_id, created_on)
	SELECT
		id, recv, mesg, timeout, root_promise_id, $1
	FROM
		callbacks
	WHERE
		promise_id = $2
	ORDER BY
		id`

	TASK_UPDATE_STATEMENT = `
	UPDATE
		tasks
	SET
		process_id = $1, state = $2, counter = $3, attempt = $4, ttl = $5, expires_at = $6, completed_on = $7
	WHERE
		id = $8 AND state & $9 != 0 AND counter = $10`

	TASK_COMPLETE_BY_ROOT_ID_STATEMENT = `
	UPDATE
		tasks
	SET
		state = 8, completed_on = $1 -- State = 8 -> Completed
	WHERE
		root_promise_id = $2 AND state in (1, 2, 4) -- State in (Init, Enqueued, Claimed)`

	TASK_HEARTBEAT_STATEMENT = `
	UPDATE
		tasks
	SET
		expires_at =
			CASE
				WHEN $1 > 9223372036854775807 - ttl THEN 9223372036854775807 -- max int64
				ELSE $1 + ttl
			END
	WHERE
		process_id = $2 AND state = 4`
)

// Config

type Config struct {
	Size      int               `flag:"size" desc:"submission buffered channel size" default:"1000"`
	BatchSize int               `flag:"batch-size" desc:"max submissions processed per iteration" default:"1000"`
	Workers   int               `flag:"workers" desc:"number of workers" default:"1" dst:"1"`
	Host      string            `flag:"host" desc:"postgres host" default:"localhost"`
	Port      string            `flag:"port" desc:"postgres port" default:"5432"`
	Username  string            `flag:"username" desc:"postgres username"`
	Password  string            `flag:"password" desc:"postgres password"`
	Database  string            `flag:"database" desc:"postgres database" default:"resonate" dst:"resonate_dst"`
	Query     map[string]string `flag:"query" desc:"postgres query options" dst:"{\"sslmode\":\"disable\"}" dev:"{\"sslmode\":\"disable\"}"`
	TxTimeout time.Duration     `flag:"tx-timeout" desc:"postgres transaction timeout" default:"10s"`
	Reset     bool              `flag:"reset" desc:"reset postgres db on shutdown" default:"false" dst:"true"`
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

type PostgresStore struct {
	config  *Config
	sq      chan<- *bus.SQE[t_aio.Submission, t_aio.Completion]
	db      *sql.DB
	workers []*PostgresStoreWorker
}

func (s *PostgresStore) DB() *sql.DB { _ = "STUB: not implemented"; return nil }

type ConnConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	Query    map[string]string
}

func NewConn(config *ConnConfig) (*sql.DB, error) { _ = "STUB: not implemented"; return nil, nil }

func New(aio aio.AIO, metrics *metrics.Metrics, config *Config) (*PostgresStore, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *PostgresStore) String() string { _ = "STUB: not implemented"; return "" }

func (s *PostgresStore) Kind() t_aio.Kind { _ = "STUB: not implemented"; return *new(t_aio.Kind) }

func (s *PostgresStore) Start(chan<- error) error { _ = "STUB: not implemented"; return nil }

// Get pending migrations

// If version == 0, the db is fresh and we can apply all migrations automatically

// Validate migration sequence

// Apply all migrations

// For existing databases, check for pending migrations and error if any exist

func (s *PostgresStore) Stop() error { _ = "STUB: not implemented"; return nil }

func (s *PostgresStore) Reset() error { _ = "STUB: not implemented"; return nil }

func (s *PostgresStore) Enqueue(sqe *bus.SQE[t_aio.Submission, t_aio.Completion]) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *PostgresStore) Flush(t int64) { _ = "STUB: not implemented"; return }

func (s *PostgresStore) Process(sqes []*bus.SQE[t_aio.Submission, t_aio.Completion]) []*bus.CQE[t_aio.Submission, t_aio.Completion] {
	_ = "STUB: not implemented"
	return nil
}

// Worker

type PostgresStoreWorker struct {
	config  *Config
	i       int
	db      *sql.DB
	sq      <-chan *bus.SQE[t_aio.Submission, t_aio.Completion]
	flush   chan int64
	aio     aio.AIO
	metrics *metrics.Metrics
}

func (w *PostgresStoreWorker) String() string { _ = "STUB: not implemented"; return "" }

func (w *PostgresStoreWorker) Start() { _ = "STUB: not implemented"; return }

func (w *PostgresStoreWorker) Flush(t int64) {
	_ = "STUB: not implemented"
	// ignore case where flush channel is full,
	// this means the flush is waiting on the cq
	return
}

func (w *PostgresStoreWorker) Process(sqes []*bus.SQE[t_aio.Submission, t_aio.Completion]) []*bus.CQE[t_aio.Submission, t_aio.Completion] {
	_ = "STUB: not implemented"
	return nil
}

func (w *PostgresStoreWorker) Execute(transactions []*t_aio.Transaction) ([]*t_aio.StoreCompletion, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *PostgresStoreWorker) performCommands(tx *sql.Tx, transactions []*t_aio.Transaction) ([]*t_aio.StoreCompletion, error) {
	_ = "STUB: not implemented"
	// Lazily defined prepared statements
	return nil, nil
}

// Results

// Promises

// Callbacks

// Schedules

// Tasks

// Promises

func (w *PostgresStoreWorker) readPromise(tx *sql.Tx, cmd *t_aio.ReadPromiseCommand) (*t_aio.QueryPromisesResult, error) {
	_ = "STUB: not implemented"
	// select
	return nil, nil
}

func (w *PostgresStoreWorker) readPromises(tx *sql.Tx, cmd *t_aio.ReadPromisesCommand) (*t_aio.QueryPromisesResult, error) {
	_ = "STUB: not implemented"
	// select
	return nil, nil
}

func (w *PostgresStoreWorker) searchPromises(tx *sql.Tx, cmd *t_aio.SearchPromisesCommand) (*t_aio.QueryPromisesResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// convert query

// convert list of state to bit mask

// tags

// select

func (w *PostgresStoreWorker) createPromise(_ *sql.Tx, stmt *sql.Stmt, cmd *t_aio.CreatePromiseCommand) (*t_aio.AlterPromisesResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// insert

func (w *PostgresStoreWorker) createPromiseAndTask(tx *sql.Tx, promiseStmt *sql.Stmt, taskStmt *sql.Stmt, cmd *t_aio.CreatePromiseAndTaskCommand) (*t_aio.AlterPromisesResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *PostgresStoreWorker) updatePromise(tx *sql.Tx, stmt *sql.Stmt, cmd *t_aio.UpdatePromiseCommand) (*t_aio.AlterPromisesResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// update

// Callbacks

func (w *PostgresStoreWorker) createCallback(tx *sql.Tx, stmt *sql.Stmt, cmd *t_aio.CreateCallbackCommand) (*t_aio.AlterCallbacksResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *PostgresStoreWorker) deleteCallbacks(tx *sql.Tx, stmt *sql.Stmt, cmd *t_aio.DeleteCallbacksCommand) (*t_aio.AlterCallbacksResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Schedules

func (w *PostgresStoreWorker) readSchedule(tx *sql.Tx, cmd *t_aio.ReadScheduleCommand) (*t_aio.QuerySchedulesResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *PostgresStoreWorker) readSchedules(tx *sql.Tx, cmd *t_aio.ReadSchedulesCommand) (*t_aio.QuerySchedulesResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *PostgresStoreWorker) searchSchedules(tx *sql.Tx, cmd *t_aio.SearchSchedulesCommand) (*t_aio.QuerySchedulesResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// convert query

// tags

// select

func (w *PostgresStoreWorker) createSchedule(tx *sql.Tx, stmt *sql.Stmt, cmd *t_aio.CreateScheduleCommand) (*t_aio.AlterSchedulesResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *PostgresStoreWorker) updateSchedule(tx *sql.Tx, stmt *sql.Stmt, cmd *t_aio.UpdateScheduleCommand) (*t_aio.AlterSchedulesResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *PostgresStoreWorker) deleteSchedule(tx *sql.Tx, stmt *sql.Stmt, cmd *t_aio.DeleteScheduleCommand) (*t_aio.AlterSchedulesResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Tasks

func (w *PostgresStoreWorker) readTask(tx *sql.Tx, cmd *t_aio.ReadTaskCommand) (*t_aio.QueryTasksResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *PostgresStoreWorker) readTasks(tx *sql.Tx, cmd *t_aio.ReadTasksCommand) (*t_aio.QueryTasksResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *PostgresStoreWorker) readEnqueueableTasks(tx *sql.Tx, cmd *t_aio.ReadEnqueueableTasksCommand) (*t_aio.QueryTasksResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *PostgresStoreWorker) createTask(tx *sql.Tx, stmt *sql.Stmt, cmd *t_aio.CreateTaskCommand) (*t_aio.AlterTasksResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// insert

func (w *PostgresStoreWorker) createTasks(tx *sql.Tx, stmt *sql.Stmt, cmd *t_aio.CreateTasksCommand) (*t_aio.AlterTasksResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *PostgresStoreWorker) completeTasks(tx *sql.Tx, stmt *sql.Stmt, cmd *t_aio.CompleteTasksCommand) (*t_aio.AlterTasksResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *PostgresStoreWorker) updateTask(tx *sql.Tx, stmt *sql.Stmt, cmd *t_aio.UpdateTaskCommand) (*t_aio.AlterTasksResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *PostgresStoreWorker) heartbeatTasks(tx *sql.Tx, stmt *sql.Stmt, cmd *t_aio.HeartbeatTasksCommand) (*t_aio.AlterTasksResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *PostgresStoreWorker) validFencingToken(tx *sql.Tx, transaction *t_aio.Transaction) (bool, error) {
	_ = "STUB: not implemented"
	// if the task is not provided continue with the operation
	return false, nil
}
