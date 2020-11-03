package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"strings"

	"github.com/joho/godotenv"
	db "github.com/ubclaunchpad/news-sentiment/db"
	"github.com/gorilla/mux"
)



func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}





//GET: endpoint for a single user
//Just print stuff
// MAKE TYPE SAFE
func getUser(w http.ResponseWriter, req *http.Request){
	id := strings.TrimPrefix(req.URL.Path, "/users/")
	fmt.Fprintf(w, id)
	
	user := db.User {
		ID:    id,
		Name:  "sdsfsfsd",
		Email: "email",
	}

	// db.fetchUser(user)
	fmt.Fprintf(w, user.Name)
}


//GET: endpoint for a single article
func getArticle(w http.ResponseWriter, req *http.Request){
	reqBody, err := ioutil.ReadAll(req.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s", reqBody)
}




func handleRequests(){

	// creates a new instance of a mux router
	router := mux.NewRouter().StrictSlash(true)
	
    // replace http.HandleFunc with router.HandleFunc
    router.HandleFunc("/", homePage)
	// router.HandleFunc("/all", returnAllArticles)
	


    // finally, instead of passing in nil, we want
    // to pass in our newly created router as the second
    // argument
    log.Fatal(http.ListenAndServe(":8090", router))
}




func main() {
	if err := godotenv.Load("example.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	db.Init()

	//Controls all endpoints
	handleRequests()
	// // 
	// http.HandleFunc("/hello", hello)
	// http.HandleFunc("/headers", headers)

	// //API Route stubs--set em up...
	// http.HandleFunc("/users/:id", getUser)
	// http.HandleFunc("/articles/:id", getArticle)
	//PUT POST GET DEL... (User Article)




	// fmt.Println("Running server on port 8090")
	// http.ListenAndServe(":8090", nil)
}
