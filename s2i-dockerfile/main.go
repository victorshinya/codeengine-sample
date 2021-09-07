package main

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello team! Welcome to Code Engine from Source to Image feature using a Dockerfile")
}

func main() {
	http.HandleFunc("/", Handler)
	fmt.Printf("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
