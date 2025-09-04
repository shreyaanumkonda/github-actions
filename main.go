package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", sayhello)
	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
func sayhello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request: %s %s\n", r.Method, r.URL.Path)

	if r.Method == "GET" {
		fmt.Println("Sending response: Hello, World!")
		fmt.Fprintf(w, "Hello, World!")
	} else {
		fmt.Printf("Method not allowed: %s\n", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
