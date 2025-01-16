package main 

import (
    "flag"
    "log"
    "net/http"
)

func main() {
    // Set command-line flags for runtime
    addr := flag.String("addr", ":4000", "HTTP network address")
    flag.Parse()

    // Initialize router
    mux := http.NewServeMux()

    // Initialize file server
    fileServer := http.FileServer(http.Dir("./ui/static"))
    
    // Register file server handler
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

    // Register route handlers
    mux.HandleFunc("/", home)
    mux.HandleFunc("/snippet/view", snippetView)
    mux.HandleFunc("/snippet/create", snippetCreate)
    
    // Start web server
    log.Printf("starting server on %s", *addr)
    err := http.ListenAndServe(*addr, mux)
    log.Fatal(err)
}
