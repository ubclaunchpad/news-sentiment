package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"

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
func (c *Database) FindAllArticles() (int64, []Article, error) {
	collection := c.database.Collection("articles")
	cursor, err := collection.Find(context.TODO(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.TODO())

	count, err := collection.CountDocuments(context.TODO(), bson.D{})
	articles := make([]Article, count)
	//if err := cursor.All(context.TODO(), &articles); err != nil {
	//	log.Fatal(err)
	//}
	for cursor.Next(context.TODO()) {
		var article Article
		if err := cursor.Decode(&article); err != nil {
			log.Fatal(err)
		}
		articles = append(articles, article)
		fmt.Println(article)
	}
	return count, articles, errors.New("Unable to find all articles")

}

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
