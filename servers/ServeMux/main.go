package main

import (
	"io"
	"net/http"
)

func firstMethod(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "first first first")
}

func secondMethod(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "second second second")
}

func main() {
	//routing
	http.HandleFunc("/first", firstMethod)
	http.HandleFunc("/second", secondMethod)

	// when handler is nil DefaultServeMux is used
	http.ListenAndServe(":8080", nil)
}
