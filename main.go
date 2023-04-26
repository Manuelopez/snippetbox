package main

import (
    "log"
    "net/http"
)

func home(w http.ResponseWriter, r *http.Request){
    if r.URL.Path != "/"{
        http.NotFound(w, r)
        return
    }
    w.Write([]byte("Hello from Snippetbox"))
}

// add a snippetVIew handler function.
func snippetView(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("Display a specific snippet..."))
}

// add a snippetCreate handle function

func snippetCreate(w http.ResponseWriter, r *http.Request){
    // use r.Method to check whter request is using post or not

    if r.Method != "POST"{
        // use the Header().Set() method to add an 'Allow: POST header' to the 
        //    resposne header map. The first parameter is the header name, and
        //    the second parameter is the hader value
        w.Header().Set("Allow", http.MethodPost)

        // if it's not, use the w.WriteHeader() method to send a 405 datus
        // code and the w.Write() method to wirte a "method not allowed"
        // response body. We then return from the function so that the 
        //    subsequent code is not executed
       // *w.WriteHeader(405)
        //*w.Write([]byte("Mehotd Not Allowed"))

        // Us the http.Error() function to send a 405 status code and "Method Not allowed"
        //    string as the presonse body

        http.Error(w, "Mehotd Not Allowed", http.StatusMethodNotAllowed)
        return
    }
    w.Write([]byte("Create a new snippet"))
}

func main(){
    // define a home handler funciton which writes a byte silce containg
    // register the home function as the handler "" url pattern
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)
    mux.HandleFunc("/snippet/view", snippetView)
    mux.HandleFunc("/snippet/create", snippetCreate)

    // use the http.ListenAndServe() function to start a new webserver. We pass in:
    //      two paratmetrs the TCP network address to listen on (in this case ": 4000")
    //      and the servemux we just created. if http.ListenAndServe() return an error
    //      we us the lof.Fatal() function to lof the error message and exit. Note
    //          that nay error returned by http.ListenAnsdServe() is always non null
    log.Print("Starting server on :4000")
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)


}
