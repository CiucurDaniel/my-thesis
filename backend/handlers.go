package main

import (
	"encoding/json"
	"fmt"
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
