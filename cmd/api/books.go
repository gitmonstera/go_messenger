package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (app *application) createBookHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Create a new Books")
}

func (app *application) showBookHandler(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(param, 10, 64)

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show the details of book %d\n", id)
}
