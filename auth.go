package main

import (
	"encoding/json"
	"net/http"
)

var users = []User{
	{ID: 1, Username: "user1", Password: "pass1"},
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)

	for _, u := range users {
		if u.Username == user.Username && u.Password == user.Password {
			json.NewEncoder(w).Encode(u)
			return
		}
	}
	http.Error(w, "Invalid credentials", http.StatusUnauthorized)
}

func Register(w http.ResponseWriter, r *http.Request) {
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = len(users) + 1
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}
