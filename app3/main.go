
package main

import (
    "fmt"
    "log"
    "net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Health is OK")
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Data endpoint from app3")
}

func main() {
    http.HandleFunc("/health", healthHandler)
    http.HandleFunc("/data", dataHandler)
    log.Println("Starting server on port 8003...")
    log.Fatal(http.ListenAndServe(":8003", nil))
}
