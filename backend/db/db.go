package db

import (
	"log"

	"github.com/joho/godotenv"
)

// init mongo
func init() {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: %s", err.Error())
	}

	// get credentials
	// example: mIP := os.Getenv("MONGO_IP")
	initSession()
}

// create new ...
func createNewUser() (string, error) {
	return "test", nil
}

// need more specifications on what the user will be doing
