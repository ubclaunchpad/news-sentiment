package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	db "github.com/ubclaunchpad/news-sentiment/db"
	"github.com/jdkato/prose/v2"  //NLP Library
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "Hello wdsdsorld!\n")

	// Create a new document with the default configuration:
    doc, err := prose.NewDocument("Go is an open-source programming language created at Google.")
    if err != nil {
        log.Fatal(err)
    }

    // Iterate over the doc's tokens:
    for _, tok := range doc.Tokens() {
        fmt.Println(tok.Text, tok.Tag, tok.Label)
        // Go NNP B-GPE
        // is VBZ O
        // an DT O
        // ...
    }



}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	if err := godotenv.Load("example.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	db.Init()

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	fmt.Println("Running server on port 8090")
	http.ListenAndServe(":8090", nil)
}
