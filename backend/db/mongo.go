package db

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// insert User into mongo
func (c *Database) insertUser(user User) (string, error) {
	collection := c.database.Collection("users") // TODO: Un-hardcode this
	insertResult, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return "", errors.Wrap(err, "Error inserting into mongo database")
	}
	if oid, ok := insertResult.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex(), nil
	}
	return "", errors.New("Invalid id created")
}

// insert NewsPiece into mongo
// func (c *Database) insertArticle(article Article) error {
// collection := c.database.Collection("articles")
// to set the id ourselves we need to
// insertResult, err := collection.InsertOne(context.TODO(), &Article{})
// 	return errors.New("Invalid id created")
// }

// fetch User from mongo
// func fetchUser(findUser User) (User, error) {
// 	return User{}, nil
// }

// fetch NewsPiece from mongo
// func fetchArticle(findNews Article) (Article, error) {
// 	return Article{}, nil
// }

// Insert a Vote
// func insertVote(vote Vote) error {
// 	return nil
// }

// Get votes for article
// func fetchVotes(article Article) ([]Vote, error) {
// 	return nil, nil
// }
