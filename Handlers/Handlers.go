package handlers

import (
	models "EndPoint/Models"
	"encoding/json"
	"strconv"

	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Se obtienen todos los usuarios")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application-json")
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])

	response := models.GetDefaultResponse()
	user, err := models.GetUser(userId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		response.NotFound()
	} else {
		response.Data = user
	}

	data, _ := json.Marshal(&response)
	fmt.Fprintf(w, string(data))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Se crea un usuario")
}

func UdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Se actualiza un usuario")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Se elimina un usuario")
}
