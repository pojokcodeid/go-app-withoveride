package main

import (
	"fmt"
	"log"
	"myapp-main/core"
	"net/http"
)

func main() {
	fmt.Println("=== Starting MyApp Main ===")

	handler := core.NewUserHandler()

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler.CreateUser(w, r)
		case http.MethodGet:
			handler.GetUsers(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
