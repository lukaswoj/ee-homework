package main

import (
  "os"
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
  "github.com/gorilla/handlers"
)

func homeLink(w http.ResponseWriter, r *http.Request) {

  // TODO: extract fetching of hostname into single place, executed once per process, not once per request
  hostname, err := os.Hostname()

  if err != nil {
    panic(err)
  }


  fmt.Fprintf(w, "Hello world from " + hostname + " on node " + os.Getenv("NODE_IP") + " <br />\n")
}

func main(){
  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/", homeLink)
  loggedRouter := handlers.LoggingHandler(os.Stdout, router)
  http.ListenAndServe(":8080", loggedRouter)
}
