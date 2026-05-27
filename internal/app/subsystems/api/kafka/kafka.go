package kafka

import (
	"context"
	"encoding/json"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/go-viper/mapstructure/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	i_api "github.com/resonatehq/resonate/internal/api"
	"github.com/resonatehq/resonate/internal/app/subsystems/api"
	"github.com/resonatehq/resonate/internal/kernel/t_api"
	"github.com/resonatehq/resonate/internal/metrics"
	"github.com/resonatehq/resonate/pkg/promise"
)

// ----------------- PROMISES -----------------

type CreatePromisePayload struct {
	ID      string            `json:"id"`
	Timeout int64             `json:"timeout"`
	Param   promise.Value     `json:"param,omitempty"`
	Tags    map[string]string `json:"tags,omitempty"`
}

type CreatePromiseAndTaskPayload struct {
	Promise CreatePromisePayload `json:"promise"`
	Task    CreateTaskPayload    `json:"task"`
}

type CreateTaskPayload struct {
	ProcessID string `json:"processId"`
	TTL       int64  `json:"ttl"`
}

type ReadPromisePayload struct {
	ID string `json:"id"`
}

type CompletePromisePayload struct {
	ID    string        `json:"id"`
	State promise.State `json:"state"`
	Value promise.Value `json:"value,omitempty"`
}

type CreateCallbackPayload struct {
	PromiseID     string          `json:"promiseId"`
	RootPromiseID string          `json:"rootPromiseId"`
	Timeout       int64           `json:"timeout"`
	Recv          json.RawMessage `json:"recv"`
}

type CreateSubscriptionPayload struct {
	ID        string          `json:"id"`
	PromiseID string          `json:"promiseId"`
	Timeout   int64           `json:"timeout"`
	Recv      json.RawMessage `json:"recv"`
}

type SearchPromisesPayload struct {
	ID     string  `json:"id"`
	State  *string `json:"state,omitempty"`
	Limit  *int    `json:"limit,omitempty"`
	Cursor *string `json:"cursor,omitempty"`
}

// ----------------- SCHEDULES -----------------

type CreateSchedulePayload struct {
	ID             string            `json:"id,omitempty"`
	Description    string            `json:"description,omitempty"`
	Cron           string            `json:"cron,omitempty"`
	Tags           map[string]string `json:"tags,omitempty"`
	PromiseID      string            `json:"promiseId,omitempty"`
	PromiseTimeout int64             `json:"promiseTimeout,omitempty"`
	PromiseParam   promise.Value     `json:"promiseParam,omitempty"`
	PromiseTags    map[string]string `json:"promiseTags,omitempty"`
}

type ReadSchedulePayload struct {
	ID string `json:"id"`
}

type DeleteSchedulePayload struct {
	ID string `json:"id"`
}

type SearchSchedulesPayload struct {
	ID     *string `json:"id"`
	Limit  *int    `json:"limit,omitempty"`
	Cursor *string `json:"cursor,omitempty"`
}

// ----------------- TASKS -----------------

type ClaimTaskPayload struct {
	ID        string `json:"id"`
	Counter   int    `json:"counter"`
	ProcessID string `json:"processId"`
	TTL       int64  `json:"ttl"`
}

type CompleteTaskPayload struct {
	ID      string `json:"id"`
	Counter int    `json:"counter"`
}

type DropTaskPayload struct {
	ID      string `json:"id"`
	Counter int    `json:"counter"`
}

type HeartbeatTasksPayload struct {
	ProcessID string `json:"processId"`
}

type Config struct {
	Brokers       []string      `flag:"brokers" desc:"kafka broker addresses" default:"localhost:9092"`
	Topic         string        `flag:"topic" desc:"kafka request topic" default:"resonate"`
	Target        string        `flag:"target" desc:"target identifier for this server" default:"resonate.server"`
	ConsumerGroup string        `flag:"consumer-group" desc:"kafka consumer group" default:"resonate-servers"`
	Timeout       time.Duration `flag:"timeout" desc:"kafka server graceful shutdown timeout" default:"10s"`
}

func (c *Config) Bind(cmd *cobra.Command, flg *pflag.FlagSet, vip *viper.Viper, name string, prefix string, keyPrefix string) {
	_ = "STUB: not implemented"
	return
}

func (c *Config) Decode(value any, decodeHook mapstructure.DecodeHookFunc) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Config) New(a i_api.API, _ *metrics.Metrics) (i_api.Subsystem, error) {
	_ = "STUB: not implemented"
	return *new(i_api.Subsystem), nil
}

type Kafka struct {
	config         *Config
	api            *api.API
	consumer       *kafka.Consumer
	producer       *kafka.Producer
	shutdownCtx    context.Context
	shutdownCancel context.CancelFunc
}

func New(a i_api.API, config *Config) (i_api.Subsystem, error) {
	_ = "STUB: not implemented"
	return *new(i_api.Subsystem), nil
}

// Create consumer

// Create producer

func (k *Kafka) String() string { _ = "STUB: not implemented"; return "" }

func (k *Kafka) Kind() string { _ = "STUB: not implemented"; return "" }

func (k *Kafka) Addr() string { _ = "STUB: not implemented"; return "" }

func (k *Kafka) Start(errors chan<- error) { _ = "STUB: not implemented"; return }

// if topic does not exist, continue

// it appears that the second call to ReadMessage will block
// until the topic exists, so don't sleep (!!)

func (k *Kafka) Stop() error {
	_ = "STUB: not implemented"
	// Cancel shutdown context
	return nil
}

// Ensure all outstanding messages are flushed before closing producer
// (optional but closer to SyncProducer semantics)

type server struct {
	api    *api.API
	config *Config
	kafka  *Kafka
}

// KafkaRequest wraps a t_api request with Kafka-specific metadata
type KafkaRequest struct {
	Target        string            `json:"target"`
	ReplyTo       ReplyTo           `json:"replyTo"`
	CorrelationId string            `json:"correlationId"`
	Operation     string            `json:"operation"`
	RequestId     string            `json:"requestId,omitempty"`
	Metadata      map[string]string `json:"metadata,omitempty"`
	Payload       json.RawMessage   `json:"payload"`
}

type ReplyTo struct {
	Topic     string  `json:"topic"`
	Target    string  `json:"target"`
	Partition *int32  `json:"partition,omitempty"`
	Key       *string `json:"key,omitempty"`
}

// KafkaResponse wraps a response or error for Kafka
type KafkaResponse struct {
	Target        string          `json:"target"`
	CorrelationId string          `json:"correlationId"`
	Operation     string          `json:"operation"`
	Success       bool            `json:"success"`
	Response      json.RawMessage `json:"response,omitempty"`
	Error         *api.Error      `json:"error,omitempty"`
}

func (s *server) handleRequest(msg *kafka.Message) { _ = "STUB: not implemented"; return }

// Filter by target - discard if not for this server

// Route based on operation

// Promises

// Schedules

// Tasks

func (s *server) log(operation string, err error) { _ = "STUB: not implemented"; return }

// Helper function to process requests
func (s *server) processRequest(kafkaReq *KafkaRequest, payload t_api.RequestPayload) (t_api.ResponsePayload, *api.Error) {
	_ = "STUB: not implemented"
	return *new(t_api.ResponsePayload), nil
}

// Process the request

func (s *server) respondError(kafkaReq *KafkaRequest, error *api.Error) {
	_ = "STUB: not implemented"
	return
}

func (s *server) sendReply(kafkaReq *KafkaRequest, data []byte) { _ = "STUB: not implemented"; return }

func (s *server) send(kafkaReq *KafkaRequest, value []byte) { _ = "STUB: not implemented"; return }
