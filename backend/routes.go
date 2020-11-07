package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type ErrorJson struct {
	Err string `json:"error"`
}

func makeErrorResponse(err error) *ErrorJson {
	return &ErrorJson{Err: err.Error()}
}

func (s *server) handleRoutes() error {
	port := os.Getenv("PORT")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", s.handleHomePage())

	router.HandleFunc("/users/", s.handleAddUser()).Methods("POST")
	router.HandleFunc("/users/{id}", s.handleGetUser()).Methods("GET")
	router.HandleFunc("/articles/", s.handleAddArticle()).Methods("POST")
	router.HandleFunc("/articles/{id}", s.handleGetArticle()).Methods("GET")
	fmt.Printf("Running server on port %s\n", port)
	return http.ListenAndServe(":"+port, router)
}

func (s *server) handleHomePage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "HOME: News Analysis App")
		fmt.Println("Endpoint Hit: homePage")
	}
}

//GET: endpoint for a single user
func (s *server) handleGetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id := vars["id"]
		fmt.Fprintf(w, "Requesting details for user: "+id)
	}
}

//POST: endpoint to add a single user
func (s *server) handleAddUser() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		type UserJson struct {
			Name       string `json:"name"`
			Email      string `json:"email"`
			Preference string `json:"preference"`
		}

		var user UserJson
		if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
			s.respond(w, req, makeErrorResponse(err), 400)
			return
		}
		id, err := s.db.CreateNewUser(user.Name, user.Email)
		if err != nil {
			s.respond(w, req, makeErrorResponse(err), 500)
			return
		}
		type UserAddedResponse struct {
			ID string `json:"id"`
		}
		s.respond(w, req, UserAddedResponse{ID: id}, 200)
	}
}

//GET: endpoint for a single article
func (s *server) handleGetArticle() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id := vars["id"]
		fmt.Fprintf(w, "Requesting details for article: "+id)
	}
}

//POST: endpoint to add a single article
func (s *server) handleAddArticle() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		reqBody, _ := ioutil.ReadAll(req.Body)
		fmt.Fprintf(w, "%+v", string(reqBody))
	}
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	w.WriteHeader(status)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			json.NewEncoder(w).Encode(ErrorJson{Err: "Unable to encode response"})
		}
	}
}
