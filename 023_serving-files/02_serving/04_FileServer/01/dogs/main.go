package main

import (
	"io"
	"net/http"
)

func main() {
	http.Handle("/dog", http.FileServer(http.Dir(".")))
	// Case 1: not work
	// http.HandleFunc("/dog/", dog)
	// Case 2: not work
	// http.HandleFunc("/dog/else", dog)
	// Case 3: work
	// but should notice that the url should not be localhost:8080/dog/ but .../dog (no last slash)
	// when first testing it, it can work because the url is localhost:8080/dog
	// when testing it after testing, the browser will automatically add / after dog
	// so it cannot show the img anymore
	http.HandleFunc("/dog/toby", dog)
	http.ListenAndServe(":8080", nil)
	// from above three cases, we can conclude that the line9's mechanism
	// it will serve all the reqs visiting resources under the root ("/") with current folder's files
	// like localhost:8080 localhost:8080/dog

}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="toby.jpg">`)
}
