package main

import (
  "fmt"
  "log"
  "os"
  "net/http"
)

func main() {
  port := os.Getenv("PORT")

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "requested root, try /ping")
    log.Println("root requested")
  })

  http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "pong :3")
    log.Println("got a ping~!")
  })

  log.Println("starting the server >:3")

  log.Fatal(http.ListenAndServe(":" + port, nil))
}