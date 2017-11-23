package main

import (
	"log"
	"net/http"
	"fmt"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {

	userInput := r.URL.Query().Get("userInput")
	fmt.Fprintf(w, userInput)

	w.Header().Set("Content-Type", "text/html")

	// Serve the file to the browser
	//http.ServeFile(w, r, "Eliza.html")
}

func main() {

	directory := http.Dir("./webFiles")
	fileServer := http.FileServer(directory)
	
	// Handle on our fileServer
	http.Handle("/", fileServer)

	// Call our request handler function
	//http.HandleFunc("/", requestHandler)

	// Serve webpage to port 8080 on localhost
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Create a structure to hold our variables for ElizasResponse
type messageStruct struct {
	ElizaResponse string
	UserInput     string
}
