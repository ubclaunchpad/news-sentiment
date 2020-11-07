package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	db "github.com/ubclaunchpad/news-sentiment/db"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	if err := runServer(); err != nil {
		log.Fatalf("Error running the server: %v", err)
	}
}

func runServer() error {
	db, dbtidy, err := db.Init()
	if err != nil {
		return errors.Wrap(err, "setup database")
	}
	defer dbtidy()
	srv := &server{
		db: db,
	}
	return srv.handleRoutes()
}

type server struct {
	db *db.Database
}
