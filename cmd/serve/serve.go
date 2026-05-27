package serve

import (
	"github.com/resonatehq/resonate/cmd/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewCmd(cfg *config.Config, vip *viper.Viper) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}

// decode config

// decode plugins

// bind config file flag

// bind config

// bind plugins

// bind other flags

func Serve(cfg *config.Config) error {
	_ = "STUB: not implemented"
	// logger
	return nil
}

// metrics

// api/aio

// api middleware

// plugins

// api subsystems

// aio subsystems

// ensure only one store is enabled at a time,
// this is a hack

// start api/aio

// set default url

// instantiate system

// request coroutines

// background coroutines

// metrics server

// listen for shutdown signal

// halt until we get a shutdown signal or an error
// occurs, whichever happens first

// shutdown system

// shutdown metrics server

// control loop

// stop api/aio
