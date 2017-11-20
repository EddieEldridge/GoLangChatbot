package main

import (
	"log"
	"net/http"
)

func main() {
	// Handles HTTP requests
	http.Handle("/", http.FileServer(http.Dir("./basic")))

	// Displays Errors to the console
	log.Fatal(http.ListenAndServe(":8080", nil))
}
