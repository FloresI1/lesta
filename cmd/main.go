package main

import (
	"net/http"
	"github.com/FloresI1/lesta/handlers"
)
func main() {
	http.HandleFunc("/users", handlers.AddPlayer)
	http.HandleFunc("/get", handlers.GetPlayers)
	http.ListenAndServe(":8080" , nil)
}