package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", reqURL)
	http.ListenAndServe(":8080", nil)
}
func reqURL(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, ">>>>>>"+req.URL.Path)
}
