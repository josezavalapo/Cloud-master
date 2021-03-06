package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
  //mux stands for HTTP request multiplexer which matches an incoming
  //request to against a list of routes (registered)
  //mux sirve para solicitar router y dispatcher
)

type Image struct {
  ID  string  `json:"id,omitempty"`
  ImageName string  `json:"imagename,omitempty"`
  Resolution  string  `json:"resolution,omitempty"`
}

var images []image

func GetImages(w http.ResponseWriter, r *http.Request) {
  json.NewEncoder(w).Encode(images)
}

func GetImage(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  for _, item := range images {
    if item.ID == params["id"] {
      json.NewEncoder(w).Encode(item)
      return
    }
  }
  json.NewEncoder(w).Encode(&Image{})
}

func UploadImage(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  var image Image
  _ = json.NewDecoder(r.Body).Decode(&image)
  image.ID = params ["id"]
  images = append(images, image)
  json.NewEncoder(w).Encode(images)
}

func helloWorld(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hello World")
}

func someFunc(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Another func")
}

func main() {

    //http.HandleFunc("/", helloWorld)
    //http.HandleFunc("/func", someFunc)
    router := mux.NewRouter()
    images = append(images, Image{ID: "1", ImageName:"Algo", Resolution:"300x600px"})
    images = append(images, Image{ID: "2", ImageName: "Otra", Resolution: "500x200px"})
    router.HandleFunc("/images", GetImages).Methods("GET") //se abriría el servidor en localhost:8000/images
    router.HandleFunc("/images/{id}", GetImage).Methods("GET")
    router.HandleFunc("/images/{id}", UploadImage).Methods("POST")
    http.ListenAndServe(":8000", router)

}
