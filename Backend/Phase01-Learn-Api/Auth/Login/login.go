package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)

	if user.Username == "admin" && user.Password == "123" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Login success",
		})
	} else {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
	}
}

func main() {
	fmt.Println("Simpel Login")
	http.HandleFunc("/login", loginHandler)
	http.ListenAndServe(":8080", nil)
}
