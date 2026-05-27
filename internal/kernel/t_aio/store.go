package t_aio

import (
	"github.com/resonatehq/resonate/pkg/message"
	"github.com/resonatehq/resonate/pkg/promise"
	"github.com/resonatehq/resonate/pkg/schedule"
	"github.com/resonatehq/resonate/pkg/task"
)

type StoreSubmission struct {
	Transaction *Transaction
}

func (s *StoreSubmission) String() string { _ = "STUB: not implemented"; return "" }

type Transaction struct {
	Fence    *task.FencingToken
	Commands []Command
}

type Command interface {
	String() string
	isCommand()
}

type ReadPromiseCommand struct {
	Id string
}

func (c *ReadPromiseCommand) String() string { _ = "STUB: not implemented"; return "" }

type ReadPromisesCommand struct {
	Time  int64
	Limit int
}

func (c *ReadPromisesCommand) String() string { _ = "STUB: not implemented"; return "" }

type SearchPromisesCommand struct {
	Id     string
	States []promise.State
	Tags   map[string]string
	Limit  int
	SortId *int64
}

func (c *SearchPromisesCommand) String() string { _ = "STUB: not implemented"; return "" }

type CreatePromiseCommand struct {
	Id        string
	State     promise.State // used to create timedout promise in non-pending state
	Param     promise.Value
	Timeout   int64
	Tags      map[string]string
	CreatedOn int64
}

func (c *CreatePromiseCommand) String() string { _ = "STUB: not implemented"; return "" }

type UpdatePromiseCommand struct {
	Id          string
	State       promise.State
	Value       promise.Value
	CompletedOn int64
}

func (c *UpdatePromiseCommand) String() string { _ = "STUB: not implemented"; return "" }

type CreateCallbackCommand struct {
	Id        string
	PromiseId string
	Recv      []byte
	Mesg      *message.Mesg
	Timeout   int64
	CreatedOn int64
}

func (c *CreateCallbackCommand) String() string { _ = "STUB: not implemented"; return "" }

type DeleteCallbacksCommand struct {
	PromiseId string
}

func (c *DeleteCallbacksCommand) String() string { _ = "STUB: not implemented"; return "" }

type ReadScheduleCommand struct {
	Id string
}

func (c *ReadScheduleCommand) String() string { _ = "STUB: not implemented"; return "" }

type ReadSchedulesCommand struct {
	NextRunTime int64
	Limit       int
}

func (c *ReadSchedulesCommand) String() string { _ = "STUB: not implemented"; return "" }

type SearchSchedulesCommand struct {
	Id     string
	Tags   map[string]string
	Limit  int
	SortId *int64
}

func (c *SearchSchedulesCommand) String() string { _ = "STUB: not implemented"; return "" }

type CreateScheduleCommand struct {
	Id             string
	Description    string
	Cron           string
	Tags           map[string]string
	PromiseId      string
	PromiseTimeout int64
	PromiseParam   promise.Value
	PromiseTags    map[string]string
	NextRunTime    int64
	CreatedOn      int64
}

func (c *CreateScheduleCommand) String() string { _ = "STUB: not implemented"; return "" }

type UpdateScheduleCommand struct {
	Id          string
	LastRunTime *int64
	NextRunTime int64
}

func (c *UpdateScheduleCommand) String() string { _ = "STUB: not implemented"; return "" }

type DeleteScheduleCommand struct {
	Id string
}

func (c *DeleteScheduleCommand) String() string { _ = "STUB: not implemented"; return "" }

type ReadTaskCommand struct {
	Id string
}

func (c *ReadTaskCommand) String() string { _ = "STUB: not implemented"; return "" }

type ReadTasksCommand struct {
	States []task.State
	Time   int64
	Limit  int
}

func (c *ReadTasksCommand) String() string { _ = "STUB: not implemented"; return "" }

type ReadEnqueueableTasksCommand struct {
	Time  int64
	Limit int
}

func (c *ReadEnqueueableTasksCommand) String() string { _ = "STUB: not implemented"; return "" }

type CreateTaskCommand struct {
	Id        string
	Recv      []byte
	Mesg      *message.Mesg
	Timeout   int64
	ProcessId *string
	State     task.State
	Ttl       int64
	ExpiresAt int64
	CreatedOn int64
}

func (c *CreateTaskCommand) String() string { _ = "STUB: not implemented"; return "" }

type CreateTasksCommand struct {
	PromiseId string
	CreatedOn int64
}

func (c *CreateTasksCommand) String() string { _ = "STUB: not implemented"; return "" }

type CompleteTasksCommand struct {
	RootPromiseId string
	CompletedOn   int64
}

func (c *CompleteTasksCommand) String() string { _ = "STUB: not implemented"; return "" }

type UpdateTaskCommand struct {
	Id             string
	ProcessId      *string
	State          task.State
	Counter        int
	Attempt        int
	Ttl            int64
	ExpiresAt      int64
	CompletedOn    *int64
	CurrentStates  []task.State
	CurrentCounter int
}

func (c *UpdateTaskCommand) String() string { _ = "STUB: not implemented"; return "" }

type HeartbeatTasksCommand struct {
	ProcessId string
	Time      int64
}

func (c *HeartbeatTasksCommand) String() string { _ = "STUB: not implemented"; return "" }

type CreatePromiseAndTaskCommand struct {
	PromiseCommand *CreatePromiseCommand
	TaskCommand    *CreateTaskCommand
}

func (c *CreatePromiseAndTaskCommand) String() string { _ = "STUB: not implemented"; return "" }

func (*ReadPromiseCommand) isCommand()          { _ = "STUB: not implemented"; return }
func (*ReadPromisesCommand) isCommand()         { _ = "STUB: not implemented"; return }
func (*SearchPromisesCommand) isCommand()       { _ = "STUB: not implemented"; return }
func (*CreatePromiseCommand) isCommand()        { _ = "STUB: not implemented"; return }
func (*UpdatePromiseCommand) isCommand()        { _ = "STUB: not implemented"; return }
func (*CreateCallbackCommand) isCommand()       { _ = "STUB: not implemented"; return }
func (*DeleteCallbacksCommand) isCommand()      { _ = "STUB: not implemented"; return }
func (*ReadScheduleCommand) isCommand()         { _ = "STUB: not implemented"; return }
func (*ReadSchedulesCommand) isCommand()        { _ = "STUB: not implemented"; return }
func (*SearchSchedulesCommand) isCommand()      { _ = "STUB: not implemented"; return }
func (*CreateScheduleCommand) isCommand()       { _ = "STUB: not implemented"; return }
func (*UpdateScheduleCommand) isCommand()       { _ = "STUB: not implemented"; return }
func (*DeleteScheduleCommand) isCommand()       { _ = "STUB: not implemented"; return }
func (*ReadTaskCommand) isCommand()             { _ = "STUB: not implemented"; return }
func (*ReadTasksCommand) isCommand()            { _ = "STUB: not implemented"; return }
func (*ReadEnqueueableTasksCommand) isCommand() { _ = "STUB: not implemented"; return }
func (*CreateTaskCommand) isCommand()           { _ = "STUB: not implemented"; return }
func (*CreateTasksCommand) isCommand()          { _ = "STUB: not implemented"; return }
func (*CompleteTasksCommand) isCommand()        { _ = "STUB: not implemented"; return }
func (*UpdateTaskCommand) isCommand()           { _ = "STUB: not implemented"; return }
func (*HeartbeatTasksCommand) isCommand()       { _ = "STUB: not implemented"; return }
func (*CreatePromiseAndTaskCommand) isCommand() { _ = "STUB: not implemented"; return }

type StoreCompletion struct {
	Valid   bool
	Results []Result
}

func (c *StoreCompletion) String() string { _ = "STUB: not implemented"; return "" }

type Result interface {
	String() string
	isResult()
}

type QueryPromisesResult struct {
	RowsReturned int64
	LastSortId   int64
	Records      []*promise.PromiseRecord
}

func (r *QueryPromisesResult) String() string { _ = "STUB: not implemented"; return "" }

type AlterPromisesResult struct {
	RowsAffected int64
}

func (r *AlterPromisesResult) String() string { _ = "STUB: not implemented"; return "" }

type AlterCallbacksResult struct {
	RowsAffected int64
}

func (r *AlterCallbacksResult) String() string { _ = "STUB: not implemented"; return "" }

type QuerySchedulesResult struct {
	RowsReturned int64
	LastSortId   int64
	Records      []*schedule.ScheduleRecord
}

func (r *QuerySchedulesResult) String() string { _ = "STUB: not implemented"; return "" }

type AlterSchedulesResult struct {
	RowsAffected int64
}

func (r *AlterSchedulesResult) String() string { _ = "STUB: not implemented"; return "" }

type QueryTasksResult struct {
	RowsReturned int64
	Records      []*task.TaskRecord
}

func (r *QueryTasksResult) String() string { _ = "STUB: not implemented"; return "" }

type AlterTasksResult struct {
	RowsAffected int64
}

func (r *AlterTasksResult) String() string { _ = "STUB: not implemented"; return "" }

func (r *QueryPromisesResult) isResult()  { _ = "STUB: not implemented"; return }
func (r *AlterPromisesResult) isResult()  { _ = "STUB: not implemented"; return }
func (r *AlterCallbacksResult) isResult() { _ = "STUB: not implemented"; return }
func (r *QuerySchedulesResult) isResult() { _ = "STUB: not implemented"; return }
func (r *AlterSchedulesResult) isResult() { _ = "STUB: not implemented"; return }
func (r *QueryTasksResult) isResult()     { _ = "STUB: not implemented"; return }
func (r *AlterTasksResult) isResult()     { _ = "STUB: not implemented"; return }

func AsQueryPromises(r Result) *QueryPromisesResult { _ = "STUB: not implemented"; return nil }

func AsAlterPromises(r Result) *AlterPromisesResult { _ = "STUB: not implemented"; return nil }

func AsAlterCallbacks(r Result) *AlterCallbacksResult { _ = "STUB: not implemented"; return nil }

func AsQuerySchedules(r Result) *QuerySchedulesResult { _ = "STUB: not implemented"; return nil }

func AsAlterSchedules(r Result) *AlterSchedulesResult { _ = "STUB: not implemented"; return nil }

func AsQueryTasks(r Result) *QueryTasksResult { _ = "STUB: not implemented"; return nil }

func AsAlterTasks(r Result) *AlterTasksResult { _ = "STUB: not implemented"; return nil }
