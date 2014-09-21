package main

import "github.com/gorilla/handlers"
import "net/http"
import "fmt"
import "log"
import "os"

func main() {
  http.HandleFunc("/", handler)
  log.Println("listening...")

  err := http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
  if err != nil {
    panic(err)
  }
}

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, world!")
}