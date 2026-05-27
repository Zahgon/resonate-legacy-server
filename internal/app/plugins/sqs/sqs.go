package sqs

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/go-viper/mapstructure/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/resonatehq/resonate/internal/app/plugins/base"
	"github.com/resonatehq/resonate/internal/metrics"
	"github.com/resonatehq/resonate/internal/plugins"
)

type Config struct {
	Size        int           `flag:"size" desc:"submission buffered channel size" default:"100"`
	Workers     int           `flag:"workers" desc:"number of workers" default:"1"`
	Timeout     time.Duration `flag:"timeout" desc:"request timeout" default:"30s"`
	TimeToRetry time.Duration `flag:"ttr" desc:"time to wait before resending" default:"15s"`
	TimeToClaim time.Duration `flag:"ttc" desc:"time to wait for claim before resending" default:"0"`
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

type Client interface {
	SendMessage(ctx context.Context, params *sqs.SendMessageInput, opt ...func(*sqs.Options)) (*sqs.SendMessageOutput, error)
}

type SQS struct {
	*base.Plugin
}

type Addr struct {
	Url    string  `json:"url"`
	Region *string `json:"region,omitempty"`
}

type processor struct {
	client  Client
	timeout time.Duration
}

func New(metrics *metrics.Metrics, config *Config) (*SQS, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewWithClient(metrics *metrics.Metrics, config *Config, client Client) (*SQS, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *processor) Process(data []byte, head map[string]string, body []byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// nosemgrep: range-over-map

func parseSQSRegion(sqsUrl string) (string, bool) { _ = "STUB: not implemented"; return "", false }
