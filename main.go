package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"rest/user-service/controllers"
	"rest/user-service/models"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

func main() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/users", getUsers)
	myRouter.HandleFunc("/users/add", saveUser).Methods("POST")

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", myRouter))
}

func saveUser(w http.ResponseWriter, r *http.Request) {

	var request models.UserRequest
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

func getUsers(w http.ResponseWriter, r *http.Request) {
}
