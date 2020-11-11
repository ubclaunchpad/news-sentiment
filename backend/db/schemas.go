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
	Source string
	Title  string
	URL    string // is url
	Votes  []Vote
}

// Vote represents a vote
type Vote struct {
	UserID    string
	ArticleID string // url of article
	VoteValue int32
}
