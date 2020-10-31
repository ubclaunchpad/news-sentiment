package db

//import (
//	"context"
//	"fmt"
//	"os"
//	"time"
//
//	"go.mongodb.org/mongo-driver/bson"
//	"go.mongodb.org/mongo-driver/bson/primitive"
//	"go.mongodb.org/mongo-driver/mongo"
//	"go.mongodb.org/mongo-driver/mongo/options"
//)

// User represents the data that a user would have
type User struct {
	ID         string
	Name       string
	Email      string
	Preference string
	Votes      []Vote
}

// NewsPiece represents a News Article
type Article struct {
	Source string
	Title  string
	URL    string
	Votes  []Vote
}

// Vote represents a vote
type Vote struct {
	UserId    string
	ArticleId string
	VoteValue int32
}
