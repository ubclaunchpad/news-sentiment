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
func createNewUser(id string) (string, error) {
	userError := insertUser(User{
		ID: id,
	})

	if userError != nil {
		log.Fatal(userError.Error())
	}
	return "User created!", userError
}

func createNewNewsPiece(id string, author string, title string, source string, user User) (string, error) {
	newsError := insertNewsPiece(NewsPiece{
		ID:     id,
		Title:  title,
		Author: author,
		Source: source,
	}, user)

	if newsError != nil {
		log.Fatal(newsError.Error())
	}
	return "NewsPiece created!", newsError

}

// need more specifications on what the user will be doing
