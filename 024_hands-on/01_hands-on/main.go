package main

import (
	"html/template"
	"io"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./*"))
}

func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "foo ran\n")
}

func dog(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "temp.html", "This is a dog.")
}

func dogPic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "toby.jpg")
}
func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/dog/toby.jpg", dogPic)

	http.ListenAndServe(":8080", nil)
}
