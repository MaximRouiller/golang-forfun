package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
	// 	fmt.Fprintf(w, "Hello from golang on path %q", html.EscapeString(req.URL.Path))
	// })
	//
	// log.Fatal(http.ListenAndServe(":5000", nil))

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":5000", router))
}

func Index(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "Hello again from golang!")
}
