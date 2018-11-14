package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// in standard http pack we should always use uri like "/.../"
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/dog/d/toby.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Println("reach here")
	// this is /dog/ handle func
	// src link is relative
	// so the combined path to img is /dog/d/toby.jpg
	io.WriteString(w, `<img src="d/toby.jpg">	`)
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	fmt.Println("reach")
	http.ServeFile(w, req, "toby.jpg")
}
