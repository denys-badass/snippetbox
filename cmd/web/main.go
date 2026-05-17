package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fileserver := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileserver))

	mux.HandleFunc("GET /{$}", getHome)
	mux.HandleFunc("GET /snippet/view/{id}", getSnippetView)
	mux.HandleFunc("GET /snippet/create", getSnippetCreate)
	mux.HandleFunc("POST /snippet/create", postSnippetCreate)

	log.Print("Server is starting on :4000")

	log.Fatal(http.ListenAndServe(":4000", mux))
}
