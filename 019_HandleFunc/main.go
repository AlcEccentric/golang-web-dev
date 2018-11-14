package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {

	// The difference between Handle and HandleFunc:
	// We need to pass a Handler instance to Handle
	// but pass a func having same func signature as serveHTTP to HandleFunc
	// in this way, we dont need to declare any Handler instance at all
	// just define the func used to handle different routes

	http.HandleFunc("/dog", d)
	http.HandleFunc("/cat", c)

	http.ListenAndServe(":8080", nil)
}
