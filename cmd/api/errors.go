package main

import (
	"fmt"
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {

	app.logger.Errorw("internal error", "method", r.Method, "path", r.URL.Path, "err", err)

	writeJSONError(w, http.StatusInternalServerError, fmt.Sprintf("the server encountered a problem: %s", err))

}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {

	app.logger.Warnw("bad request", "method", r.Method, "path", r.URL.Path, "err", err)

	writeJSONError(w, http.StatusBadRequest, err.Error())

}

func (app *application) conflictResponse(w http.ResponseWriter, r *http.Request, err error) {

	app.logger.Errorw("conflict response", "method", r.Method, "path", r.URL.Path, "err", err)

	writeJSONError(w, http.StatusConflict, err.Error())

}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {

	app.logger.Warnw("not found error", "method", r.Method, "path", r.URL.Path, "err", err)

	writeJSONError(w, http.StatusNotFound, "resource not found")
}
