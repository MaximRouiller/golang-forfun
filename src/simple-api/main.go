package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello from golang on path %q", html.EscapeString(req.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}
