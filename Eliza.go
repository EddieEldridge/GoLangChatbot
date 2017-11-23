package main

import (
	"fmt"
	"log"
	"net/http"

	"./elizaResponse"
)

func chatHandler(w http.ResponseWriter, r *http.Request) {

	userInput := r.URL.Query().Get("userInput")
	response := askEliza.ElizaResponseFunc(userInput)
	fmt.Fprintf(w, response)

	w.Header().Set("Content-Type", "text/html")

}

func main() {

	// Serve the directory of our HTML and JS files
	directory := http.Dir("./webFiles")

	// Serve our directory
	fileServer := http.FileServer(directory)

	// Get a handle on our fileServer
	http.Handle("/", fileServer)

	// Call our request handler function
	http.HandleFunc("/chat", chatHandler)

	// Serve webpage to port 8080 on localhost
	log.Fatal(http.ListenAndServe(":8080", nil))
}
