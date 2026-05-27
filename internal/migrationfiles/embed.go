package migrationfiles

import (
	"embed"
)

//go:embed migrations/sqlite migrations/postgres
var MigrationsFS embed.FS

func GetMigrationFiles(dir string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func GetMigrationContent(path string) (string, error) { _ = "STUB: not implemented"; return "", nil }
