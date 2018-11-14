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
	// ServeMux is a Handler
	// it also has function serveHTTP
	// but it is much more often to use Handle and HandleFunc
	mux := http.NewServeMux()
	// Notice the slash after "dog"
	// it means all the req visit /dog and its child path will be handled by d
	// so visit /dog/nothing /dog/some/place/else wont get 404 response
	mux.Handle("/dog/", d)
	// without slash after the path
	// c will only handle req visiting /cat
	// visit /cat/else will get 404
	mux.Handle("/cat", c)

	http.ListenAndServe(":8080", mux)
}
