package db

import (
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
	"log"
)

func (s *Store) Migrate() error {
	driver, err := postgres.WithInstance(s.db.DB, &postgres.Config{})
	if err != nil {
		return nil
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres",
		driver,
	)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil {
		if err.Error() == "no change" {
			log.Println("no change made by migrations")
		}

		return err
	}

	return nil
}
