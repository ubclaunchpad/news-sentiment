package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/joho/godotenv"
	db "github.com/ubclaunchpad/news-sentiment/db"
	"github.com/gorilla/mux"
)



func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}





//GET: endpoint for a single user
func getUser(w http.ResponseWriter, req *http.Request){

	vars := mux.Vars(req)
	id := vars["id"]
	fmt.Fprintf(w, "Requesting details for user: " + id)

}


//GET: endpoint for a single article
func getArticle(w http.ResponseWriter, req *http.Request){

	vars := mux.Vars(req)
	id := vars["id"]
	fmt.Fprintf(w, "Requesting details for article: " + id)

}




func handleRequests(){

	// creates a new instance of a mux router
	router := mux.NewRouter().StrictSlash(true)
	
	router.HandleFunc("/", homePage)
	router.HandleFunc("/users/{id}", getUser)
	router.HandleFunc("/articles/{id}", getArticle)

    log.Fatal(http.ListenAndServe(":8090", router))
}




func main() {
	if err := godotenv.Load("example.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	db.Init()

	fmt.Println("Starting Server...")
	//Controls all endpoints
	handleRequests()
}
