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
	ID string
}

type NewsPiece struct {
	ID 		string
	Title 	string
	Author 	string
	Source	string
}
