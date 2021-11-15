package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"io/ioutil"
	"log"
	"net/http"
	"rest/user-service/controllers"
	"rest/user-service/models"
	"time"
)

func main() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/users", getUsers)
	myRouter.HandleFunc("/users/add", saveUser).Methods("POST")

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", myRouter))
}

func saveUser(w http.ResponseWriter, r *http.Request) {

	var request SaveUserRequest
	reqBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(reqBody, &request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	id, _ := uuid.New()

	user := &models.User{
		ID:          id,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Description: request.Description,
	}
	errC := controllers.SaveUser(user)

	if errC != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

type user struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var users = []user{
	{ID: "1", Name: "User 01"},
}

type SaveUserRequest struct {
	Description string `json:"description"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
}
