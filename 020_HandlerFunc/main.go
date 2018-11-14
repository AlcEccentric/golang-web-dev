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
	// Although HandlerFunc is defined in go like:
	// type HandlerFunc func(ResponseWriter, * Request)
	// but HandlerFunc is not only a func, it's also a type
	// which receive ServeHTTP func
	// So HandlerFunc is also a Handler
	// We can not write like these:
	// http.Handle("/dog", d)
	// http.Handle("/cat", c)
	// since d, c are just func not a type which receives ServeHTTO (i.e., they are not Hanlders)
	// but when we force to conver their type into HandlerFunc
	// We can write following code:
	http.Handle("/dog", http.HandlerFunc(d))
	http.Handle("/cat", http.HandlerFunc(c))

	http.ListenAndServe(":8080", nil)
}

// this is similar to this:
// https://play.golang.org/p/X2dlgVSIrd
// ---and this---
// https://play.golang.org/p/YaUYR63b7L
