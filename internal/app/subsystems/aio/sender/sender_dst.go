package sender

import (
	"math/rand" // nosemgrep

	"github.com/go-viper/mapstructure/v2"
	"github.com/resonatehq/resonate/internal/aio"
	"github.com/resonatehq/resonate/internal/kernel/bus"
	"github.com/resonatehq/resonate/internal/kernel/t_aio"
	"github.com/resonatehq/resonate/internal/metrics"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Config

type ConfigDST struct {
	P float64 `flag:"p" desc:"probability of simulated unsuccessful request" default:"0.5" dst:"0:1"`
}

func (c *ConfigDST) Bind(cmd *cobra.Command, flg *pflag.FlagSet, vip *viper.Viper, name string, prefix string, keyPrefix string) {
	_ = "STUB: not implemented"
	return
}

func (c *ConfigDST) Decode(value any, decodeHook mapstructure.DecodeHookFunc) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ConfigDST) New(aio.AIO, *metrics.Metrics) (aio.Subsystem, error) {
	_ = "STUB: not implemented"
	return *new(aio.Subsystem), nil
}

func (c *ConfigDST) NewDST(aio aio.AIO, metrics *metrics.Metrics, r *rand.Rand, backchannel chan interface{}) (aio.SubsystemDST, error) {
	_ = "STUB: not implemented"
	return *new(aio.SubsystemDST), nil
}

// Subsystem

type SenderDST struct {
	config      *ConfigDST
	r           *rand.Rand
	backchannel chan interface{}
}

func NewDST(r *rand.Rand, backchannel chan interface{}, config *ConfigDST) (*SenderDST, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SenderDST) String() string { _ = "STUB: not implemented"; return "" }

func (s *SenderDST) Kind() t_aio.Kind { _ = "STUB: not implemented"; return *new(t_aio.Kind) }

func (s *SenderDST) Start(chan<- error) error { _ = "STUB: not implemented"; return nil }

func (s *SenderDST) Stop() error { _ = "STUB: not implemented"; return nil }

func (s *SenderDST) Process(sqes []*bus.SQE[t_aio.Submission, t_aio.Completion]) []*bus.CQE[t_aio.Submission, t_aio.Completion] {
	_ = "STUB: not implemented"
	return nil
}

// propagate the tags
