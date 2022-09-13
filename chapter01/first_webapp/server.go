package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, Request *http.Request) {
	fmt.Fprintf(writer, "Hello world %s!", Request.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
