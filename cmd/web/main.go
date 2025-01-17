package main 

import (
    "flag"
    "log"
    "net/http"
    "os"
)

func main() {
    // Set command-line flags for runtime
    addr := flag.String("addr", ":4000", "HTTP network address")
    flag.Parse()

    // Create decoupled loggers
    // redirect streams with >> for Stdout and 2>> for Stderr
    infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
    errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

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
    infoLog.Printf("starting server on %s", *addr)
    err := http.ListenAndServe(*addr, mux)
    errorLog.Fatal(err)
}
