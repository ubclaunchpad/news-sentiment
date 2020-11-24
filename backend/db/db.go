package db

import (
	"context"
	"log"
	"os"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type dbtidy func()

// Database representing mongo database connection

type Database interface {
	CreateNewUser(string, string) (string, error)
	CreateNewArticle(string, string, string) (string, error)
	FindAllArticles() ([]Article, error)
}
type MongoDatabase struct {
	mongo *mongo.Database
}

// Init func connects to mongo db
func InitMongo() (Database, dbtidy, error) {
	dbURI := os.Getenv("MONGO_URI")
	clientOptions := options.Client().ApplyURI(dbURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, nil, errors.Wrap(err, "mongo client connect")
	}
	if err := client.Ping(context.TODO(), nil); err != nil {
		return nil, nil, errors.Wrap(err, "Unable to ping database")
	}
	database := client.Database(os.Getenv("DATABASE_NAME"))
	return &MongoDatabase{mongo: database}, func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatalf("Unable to close connection: %v", err)
		}
	}, nil
}
