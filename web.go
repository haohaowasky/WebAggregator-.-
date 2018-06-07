package main

import ("fmt"
        "net/http"
        "time"
        )
func index_handler(w http.ResponseWriter, r *http.Request){
    day := time.Now()
    fmt.Fprintf(w, "<h1>Hello Alan</h1>")
    fmt.Fprintf(w, "<p>I am an software engineer and entrepreneur!</p>")
    fmt.Fprintf(w, "<p> My Email: alan91wang@gmail.com</p>")
    fmt.Fprintf(w,"<p>Today is %s </p>" , day )

}
func main(){
    http.HandleFunc("/", index_handler)
    http.HandleFunc("/about", index_about)
    http.ListenAndServe(":8000", nil)
}

func index_about(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "About Alan")
}
