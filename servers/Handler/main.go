package main

import (
	"fmt"
	"net/http"
)

// any value of type that implements ServeHTTP function from the Handler interface
// in the net/http package becomes a Handler
type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Any code you want in this func")
}

func main() {

}
