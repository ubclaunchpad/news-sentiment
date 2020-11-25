package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"

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

// insert NewsPiece associated with User into mongo
func (c *Database) insertArticle(article Article) (string, error) {
	collection := c.database.Collection("articles")
	insertResult, err := collection.InsertOne(context.TODO(), article)
	if err != nil {
		return "failed", errors.Wrap(err, "Error inserting article into mongo database.")
	}
	id, ok := insertResult.InsertedID.(primitive.ObjectID)
	if ok {
		return id.Hex(), nil
	}
	return "failed", errors.New("Invalid id created: " + id.Hex())
}

// fetch User from mongo
// func fetchUser(findUser User) (User, error) {
// 	return User{}, nil
// }

// fetch all Article from mongo
func (c *Database) FindAllArticles() ([]Article, error) {
	collection := c.database.Collection("articles")
	cursor, err := collection.Find(context.TODO(), bson.D{})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	var articles []Article

	for cursor.Next(context.TODO()) {
		var article Article
		if err := cursor.Decode(&article); err != nil {
			return nil, errors.Wrap(err, "Unable to iterate on cursor")
		}
		articles = append(articles, article)
	}
	return articles, nil

}

// fetch NewsPiece from mongo
// func fetchArticle(findNews Article) (Article, error) {
// 	return Article{}, nil
// }

// Insert a Vote
func (c *Database) insertVote(vote Vote) (string, error) {
	// make vote on database
	// link vote to article, search article by url, update vote field
	articles := c.database.Collection("articles")
	articleFilter := bson.D{{"url", vote.ArticleURL}}
	articleUpdate := bson.D{
		{"$added", bson.D{
			{"votefield", "newValue"}, // stub
		}},
	}
	articleResult, err := articles.UpdateOne(context.TODO(), articleFilter, articleUpdate)
	if err != nil {
		return "", errors.Wrap(err, "Error inserting vote to articles in mongo database")
	}

	// link vote to user, search user by id, update vote field
	users := c.database.Collection("users")
	userFilter := bson.D{{"id", vote.UserID}}
	userUpdate := bson.D{
		{"$added", bson.D{
			{"votefield", "newValue"}, // stub
		}},
	}
	userResult, err := users.UpdateOne(context.TODO(), userFilter, userUpdate)
	if err != nil {
		return "", errors.Wrap(err, "Error inserting vote to user in mongo database")
	}

	fmt.Print(userResult)
	fmt.Print(articleResult)
	return "stub", nil // stub
}

// Get votes for article
// func fetchVotes(article Article) ([]Vote, error) {
// 	return nil, nil
// }
