package config

import (
	"math/rand" // nosemgrep

	"github.com/go-viper/mapstructure/v2"
	"github.com/resonatehq/resonate/internal/aio"
	"github.com/resonatehq/resonate/internal/api"
	"github.com/resonatehq/resonate/internal/kernel/system"
	"github.com/resonatehq/resonate/internal/metrics"
	"github.com/resonatehq/resonate/internal/plugins"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Config

type Config struct {
	System      system.Config `flag:"system"`
	API         API           `flag:"api"`
	AIO         AIO           `flag:"aio"`
	MetricsAddr string        `flag:"metrics-addr" desc:"prometheus metrics server address" default:":9090"`
	LogLevel    string        `flag:"log-level" desc:"can be one of: debug, info, warn, error" default:"info"`
}

func (c *Config) Plugins() []plugin { _ = "STUB: not implemented"; return nil }

type API struct {
	Size       int           `flag:"size" desc:"submission buffered channel size" default:"1000" dst:"1:1000"`
	Auth       Auth          `flag:"auth"`
	Subsystems APISubsystems `flag:"-"`
}

type AIO struct {
	Size       int           `flag:"size" desc:"completion buffered channel size" default:"1000" dst:"1:1000"`
	Subsystems AIOSubsystems `flag:"-"`
	Plugins    AIOPlugins    `flag:"-"`
}

type Auth struct {
	PublicKey string `flag:"public-key" desc:"public key path used for jwt based authentication"`
}

func (a *API) Middleware() ([]api.Middleware, error) { _ = "STUB: not implemented"; return nil, nil }

// Plugins

type plugin interface {
	Bind(*cobra.Command, *pflag.FlagSet, *viper.Viper, string)
	Decode(*viper.Viper, mapstructure.DecodeHookFunc) error
}

type enabledFlag struct {
	key string
	val bool
}

type APISubsystems struct {
	subsystems []*apiSubsystem `mapstructure:"-"`
}

func (a *APISubsystems) Add(name string, enabled bool, subsystem APISubsystem) {
	_ = "STUB: not implemented"
	return
}

func (a *APISubsystems) All() []*apiSubsystem { _ = "STUB: not implemented"; return nil }

func (a *APISubsystems) AsPlugins() []plugin { _ = "STUB: not implemented"; return nil }

type apiSubsystem struct {
	subsystem APISubsystem
	prefix    string
	key       string
	name      string
	enabled   *enabledFlag
}

type APISubsystem interface {
	Bind(*cobra.Command, *pflag.FlagSet, *viper.Viper, string, string, string)
	Decode(any, mapstructure.DecodeHookFunc) error
	New(api.API, *metrics.Metrics) (api.Subsystem, error)
}

func (a *apiSubsystem) Name() string { _ = "STUB: not implemented"; return "" }

func (a *apiSubsystem) Enabled() bool { _ = "STUB: not implemented"; return false }

func (a *apiSubsystem) Bind(cmd *cobra.Command, flg *pflag.FlagSet, vip *viper.Viper, name string) {
	_ = "STUB: not implemented"
	return
}

func (a *apiSubsystem) Decode(vip *viper.Viper, hooks mapstructure.DecodeHookFunc) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *apiSubsystem) New(api api.API, metrics *metrics.Metrics) (api.Subsystem, error) {
	_ = "STUB: not implemented"
	return *new(api.Subsystem), nil
}

type AIOSubsystems struct {
	subsystems []*aioSubsystem `mapstructure:"-"`
}

func (a *AIOSubsystems) Add(name string, enabled bool, subsystem AIOSubsystem) {
	_ = "STUB: not implemented"
	return
}

func (a *AIOSubsystems) All() []*aioSubsystem { _ = "STUB: not implemented"; return nil }

func (a *AIOSubsystems) AsPlugins() []plugin { _ = "STUB: not implemented"; return nil }

type aioSubsystem struct {
	subsystem AIOSubsystem
	prefix    string
	key       string
	name      string
	enabled   *enabledFlag
}

type AIOSubsystem interface {
	Bind(*cobra.Command, *pflag.FlagSet, *viper.Viper, string, string, string)
	Decode(any, mapstructure.DecodeHookFunc) error
	New(aio.AIO, *metrics.Metrics) (aio.Subsystem, error)
	NewDST(aio.AIO, *metrics.Metrics, *rand.Rand, chan any) (aio.SubsystemDST, error)
}

func (a *aioSubsystem) Name() string { _ = "STUB: not implemented"; return "" }

func (a *aioSubsystem) Enabled() bool { _ = "STUB: not implemented"; return false }

func (a *aioSubsystem) Bind(cmd *cobra.Command, flg *pflag.FlagSet, vip *viper.Viper, name string) {
	_ = "STUB: not implemented"
	return
}

func (a *aioSubsystem) Decode(vip *viper.Viper, hooks mapstructure.DecodeHookFunc) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *aioSubsystem) New(aio aio.AIO, metrics *metrics.Metrics) (aio.Subsystem, error) {
	_ = "STUB: not implemented"
	return *new(aio.Subsystem), nil
}

func (a *aioSubsystem) NewDST(aio aio.AIO, metrics *metrics.Metrics, r *rand.Rand, c chan any) (aio.SubsystemDST, error) {
	_ = "STUB: not implemented"
	return *new(aio.SubsystemDST), nil
}

type AIOPlugins struct {
	plugins []*aioPlugin `mapstructure:"-"`
}

func (a *AIOPlugins) Add(name string, enabled bool, plugin AIOPlugin) {
	_ = "STUB: not implemented"
	return
}

func (a *AIOPlugins) All() []*aioPlugin { _ = "STUB: not implemented"; return nil }

func (a *AIOPlugins) AsPlugins() []plugin { _ = "STUB: not implemented"; return nil }

type aioPlugin struct {
	plugin  AIOPlugin
	prefix  string
	key     string
	name    string
	enabled *enabledFlag
}

type AIOPlugin interface {
	Bind(*cobra.Command, *pflag.FlagSet, *viper.Viper, string, string, string)
	Decode(any, mapstructure.DecodeHookFunc) error
	New(*metrics.Metrics) (plugins.Plugin, error)
}

func (a *aioPlugin) Name() string { _ = "STUB: not implemented"; return "" }

func (a *aioPlugin) Enabled() bool { _ = "STUB: not implemented"; return false }

func (a *aioPlugin) Bind(cmd *cobra.Command, flg *pflag.FlagSet, vip *viper.Viper, name string) {
	_ = "STUB: not implemented"
	return
}

func (a *aioPlugin) Decode(vip *viper.Viper, hooks mapstructure.DecodeHookFunc) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *aioPlugin) New(metrics *metrics.Metrics) (plugins.Plugin, error) {
	_ = "STUB: not implemented"
	return *new(plugins.Plugin), nil
}
