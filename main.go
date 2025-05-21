package main

import (
	"fmt"
	"log"
	"net/http"
	_ "strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Filesharer - Sharing files within your local network"))
}

func share(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Share a file within the domain"))
}

func createShare(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Hello World")
	// http.NotFound(w, r); return
	// w.Header().Add() - Set() - Del()
	// w.Header().Set("Content-Type", "application/json")

}

func download(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Downloading file..."))
}

func downloadWithID(w http.ResponseWriter, r *http.Request) {

	// w.WriteHeader(201) - WrtieHeader used to set http status code
	// w.WriteHeader(http.StatusCreated) - Status constants already available in the http package

	fileID := r.PathValue("fileID")
	msg := fmt.Sprintf("Downloading file %v ....", fileID)
	w.Write([]byte(msg))
}

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
