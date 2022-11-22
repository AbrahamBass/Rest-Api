package main

import (
	handlers "EndPoint/Handlers"
	models "EndPoint/Models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	mux := mux.NewRouter()
	models.SetDefaultUser()

	mux.HandleFunc("/api/v1/users", handlers.GetUsers).Methods("GET")
	mux.HandleFunc("/api/v1/user/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	mux.HandleFunc("/api/v1/user", handlers.CreateUser).Methods("POST")
	mux.HandleFunc("/api/v1/user/{id:[0-9]+}", handlers.UdateUser).Methods("PUT")
	mux.HandleFunc("/api/v1/user/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", mux))

}
