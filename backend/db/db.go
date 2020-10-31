package db

import (
	"log"
)

// init mongo
func Init() {
	// get credentials
	// example: mIP := os.Getenv("MONGO_IP")
	initSession()
}

// create new ...
func createNewUser(id string, email string, name string) (string, error) {
	if userError := insertUser(User{
		ID:    id,
		Name:  name,
		Email: email,
	}); userError != nil {
		log.Fatal(userError.Error())
	}

	return "User created!", nil
}

func createNewNewsPiece(id string, url string, title string, source string) (string, error) {
	if newsError := insertArticle(Article{
		ID:     id,
		Title:  title,
		URL:    url,
		Source: source,
	}); newsError != nil {
		log.Fatal(newsError.Error())
	}

	return "NewsPiece created!", nil

}

// need more specifications on what the user will be doing
