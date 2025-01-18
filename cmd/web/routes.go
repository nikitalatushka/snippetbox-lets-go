package main

import "net/http"

// The routes() method returns a servemux containing our application routes
func (app *application) routes() *http.ServeMux {

    // Initialize router
    mux := http.NewServeMux()

    // Initialize file server
    fileServer := http.FileServer(http.Dir("./ui/static"))

    // Register file server handler
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

    // Register route handlers
    mux.HandleFunc("/", app.home)
    mux.HandleFunc("/snippet/view", app.snippetView)
    mux.HandleFunc("/snippet/create", app.snippetCreate)

    return mux
}
