package main 

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    // define path to the html template files within a slice
    // base template must come first
    files := []string{
        "./ui/html/base.tmpl",
        "./ui/html/partials/nav.tmpl",
        "./ui/html/pages/home.tmpl",
    }

    // read template files into 'ts' template set; catch errors
    ts, err := template.ParseFiles(files...)
    if err != nil {
        app.errorLog.Println(err.Error())
        http.Error(w, "Internal Server Error", 500)
        return
    }

    // write template set content as the response body 'w'
    err = ts.ExecuteTemplate(w, "base", nil)
    if err != nil {
        app.errorLog.Println(err.Error())
        http.Error(w, "Internal Server Error", 500)
    }
}

func (app *application) snippetView (w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil || id < 1 {
        http.NotFound(w, r)
        return
    }

    fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func (app *application) snippetCreate (w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.Header().Set("Allow", http.MethodPost)
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }

    w.Write([]byte("Create a new snippet"))
}
