package migrations

import (
	"database/sql"
)

// SqliteMigrationStore implements MigrationStore for SQLite databases
type SqliteMigrationStore struct {
	db *sql.DB
}

// NewSqliteMigrationStore creates a new SQLite migration store
func NewSqliteMigrationStore(db *sql.DB) *SqliteMigrationStore {
	_ = "STUB: not implemented"
	return nil
}

func (s *SqliteMigrationStore) GetMigrationFiles() ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SqliteMigrationStore) GetMigrationContent(path string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *SqliteMigrationStore) GetInsertMigrationSQL() string { _ = "STUB: not implemented"; return "" }

func (s *SqliteMigrationStore) GetCurrentVersion() (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// We create the migrations table on store start, we can assume that ErrNoRows means that no migrations have been applied
// and the db is fresh

func (s *SqliteMigrationStore) GetDB() *sql.DB { _ = "STUB: not implemented"; return nil }

func (s *SqliteMigrationStore) CheckMigrations() error { _ = "STUB: not implemented"; return nil }

func (s *SqliteMigrationStore) String() string { _ = "STUB: not implemented"; return "" }

func (s *SqliteMigrationStore) Close() error { _ = "STUB: not implemented"; return nil }
