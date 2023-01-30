package main

import (
	"encoding/json"
	"fmt"
	"github.com/CiucurDaniel/my-thesis/backend/data"
	"github.com/go-chi/chi/v5"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
)

func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Backend is healthy."))
}

func login(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	// anonymous struct
	usr := struct {
		Email    string
		Password string
	}{}

	err := decoder.Decode(&usr)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	fmt.Printf("Received from Angular: %s %s", usr.Email, usr.Password)

	// todo: check in db

	// todo: return jwt

	// then Angular will use the token and store it with localStorage("jwt", data)

	w.Header().Add("Content-Type", "text")
	w.Write([]byte("secret-token"))

}

func register(w http.ResponseWriter, r *http.Request) {
	// take data

	// append

	// return response
}

func getProjectById(w http.ResponseWriter, r *http.Request) {
	projectId := chi.URLParam(r, "projectId")
	fmt.Printf("Got project id: %s\n", projectId)

	formattedProjectId, err := uuid.FromString(projectId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Printf("Got this strange %v", formattedProjectId)

	project, err := data.GetProjectById(formattedProjectId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Printf("%+v\n", project)

	err = json.NewEncoder(w).Encode(project)
	if err != nil {
		log.Fatal(err)
	}
}

func getUserById(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	fmt.Printf("Got project id: %s\n", userId)

	formattedUserId, err := uuid.FromString(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Printf("Got this strange %v", formattedUserId)

	user, err := data.GetProjectById(formattedUserId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Printf("%+v\n", user)

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Fatal(err)
	}
}