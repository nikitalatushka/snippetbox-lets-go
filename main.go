package main 

import (
    "log"       // for simple logging
    "net/http"  // for http client and server implementations
)

// create handler functions
// write byte slice as the response body
func home(w http.ResponseWriter, r *http.Request) {
    // this handler is used for the "/" url path
    // "/" is a subtree path which acts as a wildcard
    // we want the page to be displayed only for the fixed "/" path
    // users should receive a 404 response otherwise using http.NotFound()
    // restrict the root URL pattern with this simple check
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    w.Write([]byte("Hello from Snippetbox"))
}
func snippetView (w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Display a specific snippet"))
}
func snippetCreate (w http.ResponseWriter, r *http.Request) {
    // this handler will eventually create a new snippet in the database
    // which is a non-idempotent action that changes server state
    // so subsequent calls with the same input can yield different results
    // HTTP good practice is to restrict methods on routes
    // in this case we want to restrict route to act on POST request only
    // it will send a 405 Method Not Allowed unless the method is POST
    // return from create snippet so subsequent code is not executed
    // test with `$ curl -i -X POST http://localhost:4000/snippet/create`
    if r.Method != "POST" {
        
        // add "Allow: POST" to the response header map
        w.Header().Set("Allow", "POST")

        // send status code and plain-text response body
        // use helper function http.Error()
        // to replace w.WriteHeader() and w.Write([]byte())
        http.Error(w, "Method Not Allowed", 405)
        return
    }

    w.Write([]byte("Create a new snippet"))
}

func main() {
    // create servmux router
    // 'mux' recieves HTTP requests and checks URL path 
    // longer URL patterns always take precedence over shorter ones
    // you can register patterns in any order
    // request URLs are automatically sanitized
    // automatic 301 Permanent Redirects for subtree paths and sanitized paths
    // servmux doesn't support regexp-based patterns
    // does not support URL variables for RESTFUL api
    mux := http.NewServeMux()
    // register handler functions 
    // 'mux' will dispatch to the handler function of a URL path
    // fixed URL paths don't end with trailing slash "/view"
    // require exact match to call corresponding handler
    // subtree URL paths do end with trailing slach "/view/"
    // call handlers when start of request path matches
    // like a wild card "/view/" == "/view/**"
    mux.HandleFunc("/", home)
    mux.HandleFunc("/snippet/view", snippetView)
    mux.HandleFunc("/snippet/create", snippetCreate)

    // start web server listening on port 4000 and routing with 'mux'
    // if server returns an error, log.Fatal will exit server
    log.Println("starting server on :4000")
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}
