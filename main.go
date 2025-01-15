package main 

import (
    "log"       // for simple logging
    "net/http"  // for http client and server implementations
)

// create handler function
// writes a byte slice as the response body
func home(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello from Snippetbox"))
}

func main() {
    // create router (initialize new servmux)
    // register handler function for a route "/"
    // "/" is a catch-all URL pattern in servemux 
    // servmux will receive HTTP request and check the URL path
    // then dispatch to the handler function
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)

    // start web server listening on port 4000
    // uses mux for routing
    // if server returns an error, log.Fatal will exit server
    log.Println("starting server on :4000")
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}
