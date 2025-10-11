package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

// User model
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var (
	users = []User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	}
	mu sync.Mutex
)

func Home(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "Welcome to My Go HTTP Server!")
	if err != nil {
		return
	}
}

func Users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(users)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

	case http.MethodPost:
		var newUser User
		if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		mu.Lock()
		newUser.ID = len(users) + 1
		users = append(users, newUser)
		mu.Unlock()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		err := json.NewEncoder(w).Encode(newUser)
		if err != nil {
			http.Error(w, "Encode JSON Error", http.StatusBadRequest)
			return
		}

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
