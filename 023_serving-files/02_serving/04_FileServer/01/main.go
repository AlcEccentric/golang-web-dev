package main

import (
	"io"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	// Case 1: not work
	// http.HandleFunc("/dog/", dog)
	// Case 2: not work
	// http.HandleFunc("/dog/else", dog)
	// Case 3: work
	// but should notice that the url should not be localhost:8080/dog/ but .../dog (no last slash)
	// when first testing it, it can work because the url is localhost:8080/dog
	// when testing it after testing, the browser will automatically add / after dog
	// so it cannot show the img anymore
	http.HandleFunc("/dogs", dog)
	http.ListenAndServe(":8080", nil)
	// from above three cases, we can conclude that the FileServer's mechanism
	// the second answer in this link explains the mechanism clearly:
	// https://stackoverflow.com/questions/27945310/why-do-i-need-to-use-http-stripprefix-to-access-my-static-files
	//
	// The Handler returned by http.FileServer() will look for and serve the content of a file relative to the folder specified as its parameter
	// http.Handle("/", http.FileServer(http.Dir("."))) which will show files under "./" when visiting localhost:8080/
	// http.Handle("/dogs/", http.FileServer(http.Dir("."))) which will show files under "./dogs/" when visiting localhost:8080/dogs/

}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="toby.jpg">`)
}
