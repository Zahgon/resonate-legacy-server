package migrate

import (
	"github.com/resonatehq/resonate/cmd/config"
	"github.com/resonatehq/resonate/internal/app/subsystems/aio/store/migrations"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func NewCmd(cfg *config.Config, vip *viper.Viper) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}

// decode subsystems

// bind config file flag

// bind plugins

// add subcommands

func newStatusCmd(cfg *config.Config) *cobra.Command { _ = "STUB: not implemented"; return nil }

func newDryRunCmd(cfg *config.Config) *cobra.Command { _ = "STUB: not implemented"; return nil }

func newUpCmd(cfg *config.Config) *cobra.Command { _ = "STUB: not implemented"; return nil }

// getMigrationStore infers the store type, opens the database connection,
// and returns the appropriate migration store. The store owns the database
// connection and the caller is responsible for calling Close() on the store.
func getMigrationStore(cfg *config.Config) (migrations.MigrationStore, error) {
	_ = "STUB: not implemented"
	return *new(migrations.MigrationStore), nil
}
