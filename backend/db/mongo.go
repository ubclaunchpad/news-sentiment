package db

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// global variable to mongo connection

// make connection with mongo
func initSession() error {
	return nil
}

// insert User into mongo
func (c *Database) insertUser(user User) (string, error) {
	collection := c.database.Collection("users") // TODO: Un-hardcode this
	insertResult, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return "", errors.Wrap(err, "Error inserting into mongo database")
	}
	if oid, ok := insertResult.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex(), nil
	} else {
		return "", errors.New("Invalid id created")
	}
}

// insert NewsPiece associated with User into mongo
func insertArticle(article Article) error {
	return nil
}

// fetch User from mongo
func fetchUser(findUser User) (User, error) {
	return User{}, nil
}

// fetch NewsPiece associated with User from mongo
func fetchArticle(findNews Article) (Article, error) {
	return Article{}, nil
}

// Insert a Vote
func insertVote(vote Vote) error {
	return nil
}

// Get votes for article
func fetchVotes(article Article) ([]Vote, error) {
	return nil, nil
}
