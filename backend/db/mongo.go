package db

import (
	"context"
	"log"
	"net/url"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateNewUser adds new user to mongo db
func (md *MongoDatabase) CreateNewUser(email string, name string) (string, error) {
	user := User{
		Name:  name,
		Email: email,
	}

	collection := md.mongo.Collection("users") // TODO: Un-hardcode this
	insertResult, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return "", errors.Wrap(err, "Error inserting into mongo database")
	}
	if oid, ok := insertResult.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex(), nil
	}
	return "", errors.New("Invalid id created")
}

func (md *MongoDatabase) CreateNewArticle(url string, title string, source string) (string, error) {
	// need to add votes
	article := Article{
		Title:  title,
		URL:    url,
		Source: source,
	}

	collection := md.mongo.Collection("articles")
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
func (md *MongoDatabase) FindAllArticles() ([]Article, error) {
	collection := md.mongo.Collection("articles")
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

func (md *MongoDatabase) GetAllSources() ([]Source, error) {
	collection := md.mongo.Collection("articles")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	var sources []Source

	for cursor.Next(context.TODO()) {
		var article Article
		if err := cursor.Decode(&article); err != nil {
			return nil, err
		}
		u, err := url.Parse(article.URL)
		if err != nil {
			log.Println("Invalid URL Found in database!")
			continue
		}
		sources = appendIfMissing(sources, Source{URL: u.Hostname()})
	}
	return sources, nil
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
