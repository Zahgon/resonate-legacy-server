package migrations

import (
	"database/sql"
)

// PostgresMigrationStore implements MigrationStore for PostgreSQL databases
type PostgresMigrationStore struct {
	db *sql.DB
}

// NewPostgresMigrationStore creates a new PostgreSQL migration store
func NewPostgresMigrationStore(db *sql.DB) *PostgresMigrationStore {
	_ = "STUB: not implemented"
	return nil
}

func (s *PostgresMigrationStore) GetMigrationFiles() ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *PostgresMigrationStore) GetMigrationContent(path string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *PostgresMigrationStore) GetInsertMigrationSQL() string {
	_ = "STUB: not implemented"
	return ""
}

func (s *PostgresMigrationStore) GetCurrentVersion() (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// We create the migrations table on store start, we can assume that ErrNoRows means that no migrations have been applied
// and the db is fresh

func (s *PostgresMigrationStore) GetDB() *sql.DB { _ = "STUB: not implemented"; return nil }

func (s *PostgresMigrationStore) CheckMigrations() error { _ = "STUB: not implemented"; return nil }

func (s *PostgresMigrationStore) String() string { _ = "STUB: not implemented"; return "" }

func (s *PostgresMigrationStore) Close() error { _ = "STUB: not implemented"; return nil }
