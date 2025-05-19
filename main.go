package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Filesharer - Sharing files within your local network"))
}

func share(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Share a file within the domain"))
}

func download(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Downloading file..."))
}

func downloadWithID(w http.ResponseWriter, r *http.Request) {

	fileID := r.PathValue("fileID")
	w.Write([]byte("Downloading file..."))
}

func main() {

		mux := http.NewServeMux()

		mux.HandleFunc("/{$}", home)
		mux.HandleFunc("/share", share)
		mux.HandleFunc("/download", download)
		mux.HandleFunc("/download/{fileID}", downloadWithID)

		log.Print("Starting server on port 5050")

		err := http.ListenAndServe(":5050", mux)

		log.Fatal(err)
}