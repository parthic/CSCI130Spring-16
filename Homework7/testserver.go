/*create and configure a linux machine on digital ocean; serve a web page at the IP address of your
linux machine that shows your favorite quote and your name; you will need a credit card to create a digital ocean account*/

package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":9000", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	io.WriteString(res, "Going In One More Round When You Don't Think You Can That's What Makes All The Difference In Your Lifeù- Rocky")
}
