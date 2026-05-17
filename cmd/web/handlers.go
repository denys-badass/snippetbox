package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func internalError(w http.ResponseWriter, err error) {
	log.Print(err.Error())
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

func getHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/home.tmpl",
		"./ui/html/partials/nav.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		internalError(w, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		internalError(w, err)
		return
	}
}

func getSnippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display information a specify snippet with ID #%d...", id)
}

func getSnippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Show form for create a snippet"))
}

func postSnippetCreate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Create new snippet"))
}
