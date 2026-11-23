package migrations

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type migrateLogger struct {
	*log.Logger
	verbose bool
}

func (m migrateLogger) Verbose() bool {
	return m.verbose
}

func NewMigrate(s *sql.DB, verbose bool) (*migrate.Migrate, error) {
	driver, err := postgres.WithInstance(s, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	d, err := iofs.New(FS, ".")
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithInstance("iofs", d, "postgres", driver)
	if err != nil {
		return nil, err
	}
	m.Log = &migrateLogger{log.New(os.Stderr, "[migrate]", 0), verbose}

	return m, nil
}
