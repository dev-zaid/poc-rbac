package main

import (
	"encoding/json"
	"net/http"
	util "github.com/dev-zaid/poc-rbac/util"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	var newUser User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Process the new user (e.g., save to database)

	response := map[string]string{"message": "Signup successful"}
	util.jsonResponse(w, http.StatusCreated, response)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var loginUser User
	if err := json.NewDecoder(r.Body).Decode(&loginUser); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Authenticate the user (e.g., check credentials against a database)

	response := map[string]string{"message": "Login successful"}
	helper.jsonResponse(w, http.StatusOK, response)
}
