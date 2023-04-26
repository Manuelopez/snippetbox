package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request){
    if r.URL.Path != "/"{
        http.NotFound(w, r)
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
        log.Print(err.Error())
        http.Error(w, "Internal server error", 500)
        return
    }


    // use the execute tempalte() method to write the contet of the "base"
    err = ts.ExecuteTemplate(w, "base", nil)
    if err != nil{
        log.Print(err.Error())
        http.Error(w, "Internal Server error", 500)
    }

    w.Write([]byte("Hello form Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request){
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil || id < 1{
        http.NotFound(w, r)
        return
    }

    fmt.Fprintf(w, "Displaying snippet with ID %d...", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request){
    if r.Method != http.MethodPost{
        w.Header().Set("Allow", http.MethodPost)
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    } 

    w.Write([]byte("Crate a new snippet..."))
}



