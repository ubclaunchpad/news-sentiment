package db

import (
	"os"
	"testing"
)

func TestMongoConnection(t *testing.T) {
	os.Setenv("MONGO_URI", "mongodb://admin:admin@localhost:27017")
	os.Setenv("DATABASE_NAME", "news-sentiment-test")
	_, _, err := InitMongo()
	if err != nil {
		t.Errorf("Unable to init the db\n")
	}
}
