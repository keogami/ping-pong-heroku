package main

import (
  "fmt"
  "log"
  "os"
  "net/http"
)

func main() {
  port := os.Getenv("PORT") // you need to get the port from the environment, heroku explains that

  // added a handler for the root of our website "/", even 404s fallback to this path
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "requested root, try /ping")
    log.Println("root requested")
  })

  // this is the ping pong path. when some goes to /ping, we return a pong. simple~
  http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "pong :3")
    log.Println("got a ping~!")
  })

  log.Println("starting the server >:3")
  
  // servers should run indefinitely, so if the function returns, it means we crashed
  log.Fatal(http.ListenAndServe(":" + port, nil))
}