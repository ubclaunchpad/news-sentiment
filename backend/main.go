package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	db "github.com/ubclaunchpad/news-sentiment/db"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "Hello world!\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	db.Init()

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	fmt.Println("Running server on port 8090")
	http.ListenAndServe(":8090", nil)
}
