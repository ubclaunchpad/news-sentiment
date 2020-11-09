package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents the data that a user would have
type User struct {
	ID         string
	Name       string
	Email      string
	Preference string
	Votes      []Vote
}

// Article represents a News Article
type Article struct {
	ID    primitive.ObjectID
	Title string
	URL   string
	Votes []Vote
}

// Vote represents a vote
type Vote struct {
	UserID    string
	ArticleID string
	VoteValue int32
}
