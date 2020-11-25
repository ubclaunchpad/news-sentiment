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
type Database struct {
	database *mongo.Database
}

// Init func connects to mongo db
func Init() (*Database, dbtidy, error) {
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
	return &Database{database: database}, func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatalf("Unable to close connection: %v", err)
		}
	}, nil
}

// CreateNewUser adds new user to mongo db
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

func (c *Database) CreateNewArticle(url string, title string, source string) (string, error) {
	// TODO: need to add votes
	result, newsError := c.insertArticle(Article{
		Title:  title,
		URL:    url,
		Source: source,
	})
	if newsError != nil {
		return "", errors.Wrap(newsError, "Unable to create article")
	}
	return result, nil
}

// GetAllArticles gets all articles from db, formatted to Article type
func (c *Database) GetAllArticles() ([]Article, error) {
	return c.FindAllArticles()
}

func (c *Database) CreateNewVote(articleUrl string, userId string, voteValue int32) (string, error) {
	result, newsError := c.insertVote(Vote {
		ArticleURL: articleUrl,
		UserID: userId,
		VoteValue: voteValue,
	})
	if newsError != nil {
		return "", errors.Wrap(newsError, "Unable add vote")
	}
	return result, nil
}

func (c *Database) AddVoteToArticle(articleUrl string, userId string, voteValue int32) (string, error) {
	// link vote to associated article
	return "sth", nil
} // stub

func (c *Database) AddVoteToUser(userId string, articleUrl string, voteValue int32) (string, error) {
	// link vote to associated user
	return "sth", nil
} // stub