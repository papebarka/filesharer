package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	_ "strconv"
)

func home(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Server", "Go")

	files := []string{
		"./ui/html/pages/layout.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// err = ts.Execute(w, nil)
	err = ts.ExecuteTemplate(w, "layout", nil)

	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
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
