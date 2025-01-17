package main

import (
    "fmt"
    "net/http"
    "runtime/debug"
)

// serverError helper
// writes error message and stack trace to errorLog
// sends generic 500 Internal Server Error response to user
func (app *application) serverError(w http.ResponseWriter, err error) {
    trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
    app.errorLog.Output(2, trace)

    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// clientError helper
// sends a specific status code and corresponding description to user
func (app *application) clientError(w http.ResponseWriter, status int) {
    http.Error(w, http.StatusText(status), status)
}

// notFound helper
// sends 404 Not Found status code to user
func (app *application) notFound(w http.ResponseWriter) {
    app.clientError(w, http.StatusNotFound)
}

