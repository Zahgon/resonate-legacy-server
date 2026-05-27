package kafka

import (
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/go-viper/mapstructure/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/resonatehq/resonate/internal/metrics"
	"github.com/resonatehq/resonate/internal/plugins"
)

type Config struct {
	Size        int           `flag:"size" desc:"submission buffered channel size" default:"100"`
	Workers     int           `flag:"workers" desc:"number of workers" default:"1"`
	Timeout     time.Duration `flag:"timeout" desc:"kafka request timeout" default:"30s"`
	TimeToRetry time.Duration `flag:"ttr" desc:"time to wait before resending" default:"15s"`
	TimeToClaim time.Duration `flag:"ttc" desc:"time to wait for claim before resending" default:"0"`
	Brokers     []string      `flag:"brokers" desc:"kafka broker addresses" default:"localhost:9092"`
	Compression string        `flag:"compression" desc:"compression type (none, gzip, snappy, lz4, zstd)" default:"none"`
}

func (c *Config) Bind(cmd *cobra.Command, flg *pflag.FlagSet, vip *viper.Viper, name string, prefix string, keyPrefix string) {
	_ = "STUB: not implemented"
	return
}

func (c *Config) Decode(value any, decodeHook mapstructure.DecodeHookFunc) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Config) New(metrics *metrics.Metrics) (plugins.Plugin, error) {
	_ = "STUB: not implemented"
	return *new(plugins.Plugin), nil
}

type Producer interface {
	Produce(msg *kafka.Message, deliveryChan chan kafka.Event) error
	Close()
}

type Worker struct {
	i        int
	sq       <-chan *plugins.Message
	timeout  time.Duration
	metrics  *metrics.Metrics
	config   *Config
	producer Producer
}

type Kafka struct {
	sq      chan *plugins.Message
	workers []*Worker
}

type Addr struct {
	Topic   string            `json:"topic"`
	Key     *string           `json:"key,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
}

func (a *Addr) validate() error { _ = "STUB: not implemented"; return nil }

func New(metrics *metrics.Metrics, config *Config) (*Kafka, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewWithProducer(metrics *metrics.Metrics, config *Config, producer Producer) (*Kafka, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Kafka) String() string { _ = "STUB: not implemented"; return "" }

func (k *Kafka) Type() string { _ = "STUB: not implemented"; return "" }

func (k *Kafka) Addr() string { _ = "STUB: not implemented"; return "" }

func (k *Kafka) Start(chan<- error) error { _ = "STUB: not implemented"; return nil }

func (k *Kafka) Stop() error { _ = "STUB: not implemented"; return nil }

func (k *Kafka) Enqueue(msg *plugins.Message) bool { _ = "STUB: not implemented"; return false }

func (w *Worker) String() string { _ = "STUB: not implemented"; return "" }

func (w *Worker) Start() { _ = "STUB: not implemented"; return }

func (w *Worker) Process(data []byte, body []byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// nosemgrep: range-over-map
