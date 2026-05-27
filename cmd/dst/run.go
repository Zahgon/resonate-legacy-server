package dst

import (

	// nosemgrep

	"github.com/resonatehq/resonate/cmd/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func RunDSTCmd(cfg *config.Config, vip *viper.Viper) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}

// decode config

// decode plugins

// logger

// suppress time attr

// instantiate metrics

// nolint: errcheck

// set up scenarios

// instantiate backchannel

// api/aio

// aio subsystems

// ensure only one store is enabled at a time,
// this is a hack

// start api/aio

// instantiate system

// request coroutines

// background coroutines

// ms

// stop api/aio

// bind config file flag

// dst related values

// bind config

// bind plugins
