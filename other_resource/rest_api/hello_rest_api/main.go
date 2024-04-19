package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello REST API")
	})

	mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Getting hello from someone")
	})

	mux.HandleFunc("GET /hello/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "Getting hello from someone with id %v", id)
	})

	mux.HandleFunc("POST /hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Post to say hello for someone")
	})

	fmt.Println("Serve http on localhost:8089")
	err := http.ListenAndServe("localhost:8089", mux)
	if err != nil {
		log.Fatalf("Error %v", err.Error())
	}
}
