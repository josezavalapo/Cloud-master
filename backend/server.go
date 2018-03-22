package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "bytes"
    "encoding/base64"
    "image"
    "image/png"
  //mux stands for HTTP request multiplexer which matches an incoming
  //request to against a list of routes (registered)
  //mux sirve para solicitar router y dispatcher
)

type Image struct {
  ID  string  `json:"id,omitempty"`
  ImageName string  `json:"imagename,omitempty"`
  Resolution  string  `json:"resolution,omitempty"`
}

type User struct {
  ID string `json:"id,omitempty"`
  Name string `json:"name,omitempty"`
  Password string `json:"password,omitempty"`
  Email string `json:"email,omitempty"`
}

type Users []User

var images []Image

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

func getUsers(w http.ResponseWriter, r *http.Request){

  users := Users{
    User{ID : "1",
      Name: "Alan",
       Password: "vaselina",
       Email: "alan@ganso.com",
     },
    User{ID: "2",
        Name: "Monica",
        Password: "starwars",
        Email: "monica@udlap.mx",
    },
}

json.NewEncoder(w).Encode(users)
}

func signUser(w http.ResponseWriter, r *http.Request) {

}

func login(w http.ResponseWriter, r *http.Request) {

}

func encodeImage(w http.ResponseWriter, r *http.Request) {
  // Create a blank image 10x20 pixels
myImage := image.NewRGBA(image.Rect(0, 0, 10, 20))

// In-memory buffer to store PNG image
// before we base 64 encode it
var buff bytes.Buffer

// The Buffer satisfies the Writer interface so we can use it with Encode
// In previous example we encoded to a file, this time to a temp buffer
png.Encode(&buff, myImage)

// Encode the bytes in the buffer to a base64 string
encodedString := base64.StdEncoding.EncodeToString(buff.Bytes())

// You can embed it in an html doc with this string
htmlImage := "<img style=background-color:yellow; src=\"data:image/png;base64," + encodedString + "\" />"
fmt.Println(htmlImage)
json.NewEncoder(w).Encode(htmlImage)
}

func main() {

    //http.HandleFunc("/", helloWorld)
    //http.HandleFunc("/func", someFunc)
    router := mux.NewRouter()
    images = append(images, Image{ID: "1", ImageName:"Algo", Resolution:"300x600px"})
    images = append(images, Image{ID: "2", ImageName: "Otra", Resolution: "500x200px"})
    router.HandleFunc("/info", getUsers).Methods("GET")
    router.HandleFunc("/signup", getUsers).Methods("POST")
    router.HandleFunc("/login", getUsers).Methods("POST")
    router.HandleFunc("/encode", encodeImage).Methods("GET")
    router.HandleFunc("/images", GetImages).Methods("GET") //se abriría el servidor en localhost:8000/images
    router.HandleFunc("/images/{id}", GetImage).Methods("GET")
    router.HandleFunc("/images/{id}", UploadImage).Methods("POST")
    // Create a blank image 10x20 pixels
  /*  myImage := image.NewRGBA(image.Rect(0, 0, 10, 20))

 // In-memory buffer to store PNG image
 // before we base 64 encode it
    var buff bytes.Buffer

 // The Buffer satisfies the Writer interface so we can use it with Encode
 // In previous example we encoded to a file, this time to a temp buffer
    png.Encode(&buff, myImage)

 // Encode the bytes in the buffer to a base64 string
    encodedString := base64.StdEncoding.EncodeToString(buff.Bytes())

 // You can embed it in an html doc with this string
    htmlImage := "<img src=\"data:image/png;base64," + encodedString + "\" />"
    fmt.Println(htmlImage)*/
    http.ListenAndServe(":8000", router)

}
