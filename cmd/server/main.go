package main

import (
	"github.com/Ekod/go-grpc/internal/db"
	"github.com/Ekod/go-grpc/internal/rocket"
	"log"
)

func Run() error {
	//TODO gRPC server initializer
	rocketStore, err := db.New()
	if err !=nil {
		return err
	}

	err = rocketStore.Migrate()
	if err != nil {
		log.Println("Failed to run migrations")
		return err
	}

	_ = rocket.New(rocketStore)

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
