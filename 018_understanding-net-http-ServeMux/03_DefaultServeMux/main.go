package main

import (
	"io"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

type hotcat int

func (c hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {
	var d hotdog
	var c hotcat

	http.Handle("/dog", d)
	http.Handle("/cat", c)
	// Use DefaultServeMux which is built in the http package
	// Just pass nil to listenAndServe's second parameter
	// So we can use the functions of the DefaultServeMux
	// by http.Handle or http.HandleFunc
	// which is more elegant
	http.ListenAndServe(":8080", nil)
}
