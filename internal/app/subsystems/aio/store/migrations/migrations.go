package migrations

import (
	"errors"
)

var ErrPendingMigrations = errors.New("pending migrations exist")

// Migration represents a single database migration
type Migration struct {
	Version int
	Name    string
	SQL     string
}

func ParseMigrationFilename(filename string) (version int, name string, err error) {
	_ = "STUB: not implemented"
	return 0, "", nil
}

func LoadMigrations(store MigrationStore) ([]Migration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Sort by version

// GetPendingMigrations returns migrations that need to be applied
func GetPendingMigrations(currentVersion int, store MigrationStore) ([]Migration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ValidateMigrationSequence ensures migrations are sequential with no gaps
func ValidateMigrationSequence(migrations []Migration, startVersion int) error {
	_ = "STUB: not implemented"
	return nil
}

// ApplyMigrations executes migrations in a transaction
func ApplyMigrations(migrations []Migration, store MigrationStore) error {
	_ = "STUB: not implemented"
	return nil
}

// MigrationError represents an error during migration execution
type MigrationError struct {
	Version int
	Name    string
	Err     error
}

func (e *MigrationError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *MigrationError) Unwrap() error { _ = "STUB: not implemented"; return nil }
