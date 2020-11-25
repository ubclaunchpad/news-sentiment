package db

import (
	"context"
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



// fetch ALL Article from mongo if numArticles == -1 (FLAG)
//else only return FIRST "numArticles" entries IN DB
func (c *Database) FindAllArticles(numArticles int) ([]Article, error) {
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

	numToGet := 0
	for cursor.Next(context.TODO()) {
		numToGet++
		var article Article
		if err := cursor.Decode(&article); err != nil {
			return nil, errors.Wrap(err, "Unable to iterate on cursor")
		}
		articles = append(articles, article)

		//Return early if we have found the necessary number of articles
		if numArticles != -1 && numToGet >= numArticles {
			return articles, nil
		}
	}
	return articles, nil
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
