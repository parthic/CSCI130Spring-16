package main

import (
	"fmt"
	"net/http"
  "strings"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
    var name string
    fs := strings.Split(req.URL.Path, "/")
    name = fs[1]
    fmt.Fprintf(res, name)
	})

	http.ListenAndServe(":8080", nil)
}
