package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"github.com/joho/godotenv"
	db "github.com/ubclaunchpad/news-sentiment/db"
	"github.com/gorilla/mux"
)



func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "HOME: News Analysis App")
    fmt.Println("Endpoint Hit: homePage")
}





//GET: endpoint for a single user
func getUser(w http.ResponseWriter, req *http.Request){

	vars := mux.Vars(req)
	id := vars["id"]
	fmt.Fprintf(w, "Requesting details for user: " + id)

}


//POST: endpoint to add a single user
func addUser(w http.ResponseWriter, req *http.Request){

	// get the body of our POST request
    // return the string response containing the request body    
    reqBody, _ := ioutil.ReadAll(req.Body)
    fmt.Fprintf(w, "%+v", string(reqBody))

}


//GET: endpoint for a single article
func getArticle(w http.ResponseWriter, req *http.Request){

	vars := mux.Vars(req)
	id := vars["id"]
	fmt.Fprintf(w, "Requesting details for article: " + id)

}


//POST: endpoint to add a single article
func addArticle(w http.ResponseWriter, req *http.Request){

	// get the body of our POST request
    // return the string response containing the request body    
    reqBody, _ := ioutil.ReadAll(req.Body)
    fmt.Fprintf(w, "%+v", string(reqBody))

}




//Driver for all routes
func handleRequests(){

	router := mux.NewRouter().StrictSlash(true)
	
	router.HandleFunc("/", homePage)

	router.HandleFunc("/users/", addUser).Methods("POST")
	router.HandleFunc("/users/{id}", getUser)

	router.HandleFunc("/articles/", addArticle).Methods("POST")
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
