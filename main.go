package main 

import (
    "log"       // for simple logging
    "net/http"  // for http client and server implementations
)

// create handler functions
// write byte slice as the response body
func home(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello from Snippetbox"))
}
func snippetView (w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Display a specific snippet"))
}
func snippetCreate (w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Create a new snippet"))
}

func main() {
    // create servmux router
    // 'mux' recieves HTTP requests and checks URL path 
    mux := http.NewServeMux()
    // register handler functions 
    // 'mux' will dispatch to the handler function of a URL path
    mux.HandleFunc("/", home)
    mux.HandleFunc("/snippet/view", snippetView)
    mux.HandleFunc("/snippet/crete", snippetCreate)

    // start web server listening on port 4000 and routing with 'mux'
    // if server returns an error, log.Fatal will exit server
    log.Println("starting server on :4000")
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}
