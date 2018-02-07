package main

import (
    "fmt"
    "net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hello World")
}

func sum(w http.ResponseWriter, r *http.Request) {


}

func someFunc(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Another func")
}

func main() {

    http.HandleFunc("/", helloWorld)
    http.HandleFunc("/func", someFunc)
    http.ListenAndServe(":8080", nil)

}
