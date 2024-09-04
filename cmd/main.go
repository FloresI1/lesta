package main

import (
	"log"
	"net/http"
	"github.com/FloresI1/lesta/handlers"
)

func main() {
	http.HandleFunc("/users", handlers.AddPlayer)
	http.HandleFunc("/get", handlers.GetPlayers)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
