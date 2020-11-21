package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	db "github.com/ubclaunchpad/news-sentiment/db"
)

func mockServer() *server {
	return &server{
		db: &db.MockDatabase{},
	}
}

func makeJSONRequest(t *testing.T, method string, path string, body interface{}) *http.Request {
	requestBody, err := json.Marshal(body)

	if err != nil {
		t.Fatalf("Unable to marshal test body: %v\n", err)
	}
	request, err := http.NewRequest(method, path, bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Error creating request: %v\n", err)
	}

	return request
}

func TestAddArticle(t *testing.T) {
	s := mockServer()

	request := makeJSONRequest(t, "POST", "/articles", ArticleJSON{
		Source: "I am a source",
		Title:  "bobo",
		URL:    "foo.com",
	})

	rr := httptest.NewRecorder()

	handler := s.handleAddArticle()
	handler.ServeHTTP(rr, request)

	result := rr.Result()
	defer result.Body.Close()
	if result.StatusCode != http.StatusCreated {
		t.Errorf("Status code should be %v, found %v\n", http.StatusCreated, result.StatusCode)
	}
	// TODO: Add tests for the result, right now, this only returns an id
	// We could use the `GET` endpoint. This way we test both together
}
