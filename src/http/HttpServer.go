package main

import (
    "fmt"
    "net/http"
)

func csvFileDownloadHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/vnd.ms-excel")
	w.Header().Set("content-disposition", "attachment; filename='" + "test.xlsx" +"'")
	fmt.Fprintf(w, "TEST, CSV, FIEL, DOWNLOAD")
}
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/csv", csvFileDownloadHandler)
    fmt.Println("send your request to http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}