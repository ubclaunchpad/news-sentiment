package db

import (
	"context"
	"log"
	"os"

	// "go.mongodb.org/mongo-driver/bson"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type dbtidy func()

type Database struct {
	database *mongo.Database
}

func Init() (*Database, dbtidy, error) {
	dbUri := os.Getenv("MONGO_URI")
	clientOptions := options.Client().ApplyURI(dbUri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, nil, errors.Wrap(err, "mongo client connect")
	}
	if err := client.Ping(context.TODO(), nil); err != nil {
		return nil, nil, errors.Wrap(err, "Unable to ping database")
	}
	database := client.Database(os.Getenv("DATABASE_NAME"))
	return &Database{database: database}, func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatalf("Unable to close connection: %v", err)
		}
	}, nil
}

func (c *Database) CreateNewUser(email string, name string) (string, error) {
	id, userError := c.insertUser(User{
		Name:  name,
		Email: email,
	})
	if userError != nil {
		return "", errors.Wrap(userError, "Unable to create user")
	}

	return id, nil
}

func createNewArticle(url string, title string, source string) (string, error) {
	if newsError := insertArticle(Article{
		Title:  title,
		URL:    url,
		Source: source,
	}); newsError != nil {
		log.Fatal(newsError.Error())
	}

	return "Article created!", nil

}
