package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

    
)

func (app *application) home(w http.ResponseWriter, r *http.Request){
    if r.URL.Path != "/"{
        app.notFound(w)
        return
    }


    // initialize a slice containing the paths to the two files. TIs imporant
    //    to note that the file containing our base tempalte must be first
    // file in the slcie
    files := []string{
        "./ui/html/base.html",
        "./ui/html/partials/nav.html",
        "./ui/html/pages/home.html",
    }

    // use the template.ParseFiles() function to read the files and store the
    //    templates in  a temaple set. Notice that we can pass the silce of file
    //    paths as a variadic paramer
    ts, err := template.ParseFiles(files...)

    if err != nil{
        app.serverError(w, err)
        return
    }


    // use the execute tempalte() method to write the contet of the "base"
    err = ts.ExecuteTemplate(w, "base", nil)
    if err != nil{
        app.serverError(w, err)
    }

    w.Write([]byte("Hello form Snippetbox"))
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request){
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil || id < 1{
        app.notFound(w)
        return
    }

    fmt.Fprintf(w, "Displaying snippet with ID %d...", id)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request){
    if r.Method != http.MethodPost{
        w.Header().Set("Allow", http.MethodPost)
        app.clientError(w, http.StatusMethodNotAllowed)
        return
    } 

    w.Write([]byte("Crate a new snippet..."))
}



