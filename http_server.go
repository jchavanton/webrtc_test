package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
)

func main() {

    fileServer := http.FileServer(http.Dir("./public"))
    http.Handle("/", fileServer)

    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        io.WriteString(w, "Hello there!\n")
    })

    fmt.Printf("Starting server at port 8080\n")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
