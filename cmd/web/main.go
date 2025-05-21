package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /share", share)
	mux.HandleFunc("POST /share", createShare)
	mux.HandleFunc("GET /download", download)
	mux.HandleFunc("POST /download/{fileID}", downloadWithID)

	log.Print("Starting server on port 5050")

	err := http.ListenAndServe(":5050", mux)

	log.Fatal(err)
}
