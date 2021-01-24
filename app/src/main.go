package main

import (
  "os"
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
  "github.com/gorilla/handlers"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello world")
}

func main(){
  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/", homeLink)
  loggedRouter := handlers.LoggingHandler(os.Stdout, router)
  http.ListenAndServe(":8080", loggedRouter)
}
