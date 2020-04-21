package main

import (
	"io"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "foo ran")
}

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog bark's")
}

func me(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "I only think about me, me, me")
}
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)

}
