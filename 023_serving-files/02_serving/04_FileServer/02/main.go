package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	// we know that
	// http.Handle("/resources/",  http.FileServer(http.Dir("./assets")))
	// which will look for and show files under "./assets/resources/"
	// when visiting localhost:8080/resources/
	// in order to just look for and show files under "./assets/" with same url (localhost:8080/resources/)
	// we should use StripPrefix to cut /resources from /resources/ with / left
	// so that the server will look for and show files just under "./assets/"
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/resources/toby.jpg">`)
}

/*

./assets/toby.jpg

*/
