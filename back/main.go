package main

import (
    "fmt"
    "log"
    "net/http"
)

func healthcheck(writer http.ResponseWriter, request *http.Request) {
    fmt.Fprintln(writer, "Sharetube")
}

func main() {
    http.HandleFunc("/healthcheck", healthcheck)

    log.Print("Listening on port 8090 ...")
    err := http.ListenAndServe(":8090", nil)
    if err != nil {
        log.Fatal(err)
    }
}
