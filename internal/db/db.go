package db

import (
	"fmt"
	"github.com/Ekod/go-grpc/internal/rocket"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
)

type Store struct {
	db *sqlx.DB
}

func New() (Store, error) {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSL_MODE")

	connectString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", dbHost, dbPort, dbUsername, dbTable, dbPassword, dbSSLMode)

	db, err := sqlx.Connect("postgres", connectString)
	if err != nil {
		return Store{}, nil
	}

	return Store{db: db}, nil
}

func (s Store) GetRocketByID(id string) (rocket.Rocket, error) {
	return rocket.Rocket{}, nil
}

// InsertRocket - Добавляет рокету в store
func (s Store) InsertRocket(rkt rocket.Rocket) (rocket.Rocket, error) {
	return rocket.Rocket{}, nil
}

// DeleteRocket - удаляет рокету - очень быстрый демонтаж
func (s Store) DeleteRocket(id string) error {
	return nil
}