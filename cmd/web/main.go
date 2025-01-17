package main 

import (
    "flag"
    "log"
    "net/http"
    "os"
)

// Define application-wide dependencies in a struct
type application struct {
    errorLog    *log.Logger
    infoLog     *log.Logger
}

func main() {
    // Set command-line flags for runtime
    addr := flag.String("addr", ":4000", "HTTP network address")
    flag.Parse()

    // Create decoupled loggers
    // redirect streams with >> for Stdout and 2>> for Stderr
    infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
    errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

    // Initialize application struct
    app := &application{
        errorLog: errorLog,
        infoLog: infoLog,
    }

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

    // Configure server
    srv := &http.Server{
        Addr:       *addr,
        ErrorLog:   errorLog,
        Handler:    mux,
    }


    // Start web server
    infoLog.Printf("starting server on %s", *addr)
    err := srv.ListenAndServe()
    errorLog.Fatal(err)
}
