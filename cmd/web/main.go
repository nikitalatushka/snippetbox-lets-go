package main 

import (
    "log"
    "net/http"
)

func main() {
    // Initialize router
    mux := http.NewServeMux()
    
    // Register handlers
    mux.HandleFunc("/", home)
    mux.HandleFunc("/snippet/view", snippetView)
    mux.HandleFunc("/snippet/create", snippetCreate)
    
    // Start web server
    log.Println("starting server on :4000")
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}
