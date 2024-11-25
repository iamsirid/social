package main

import (
	"fmt"
	"log"
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {

	log.Printf("internal server error: %s path: %s error: %s", r.Method, r.URL.Path, err)

	writeJSONError(w, http.StatusInternalServerError, fmt.Sprintf("the server encountered a problem: %s", err))

}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {

	log.Printf("bad request error: %s path: %s error: %s", r.Method, r.URL.Path, err)

	writeJSONError(w, http.StatusBadRequest, err.Error())

}

func (app *application) conflictResponse(w http.ResponseWriter, r *http.Request, err error) {

	log.Printf("conflict error: %s path: %s error: %s", r.Method, r.URL.Path, err)

	writeJSONError(w, http.StatusConflict, err.Error())

}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {

	log.Printf("not found error: %s path: %s error: %s", r.Method, r.URL.Path, err)

	writeJSONError(w, http.StatusNotFound, "resource not found")
}
