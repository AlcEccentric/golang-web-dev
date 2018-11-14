package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*"))
}
func main() {
	// be aware of the last slash in the route
	// in this case
	// if we use /apply/ instead of /apply
	// we cannot render applyProcess.gohtml
	// because the method is always GET, even if the form in apply.gohtml uses POST to /apply
	// So, although "/apply/" will also response to "/apply", yet it will make some change to the request.
	// So it is a good practice to fit the exactly same routes when handle requests:
	// use:
	// http.HandleFunc("/apply", apply) in main.go
	// or:
	// modify <form method="POST" action="/apply"> to <form method="POST" action="/apply/">
	// and use http.HandleFunc("/apply/", apply) in main.go
	http.HandleFunc("/apply", apply)
	http.HandleFunc("/", index)
	http.HandleFunc("/about/", about)
	http.HandleFunc("/contact/", contact)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {

	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	HandleError(w, err)
}

func about(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "about.gohtml", nil)
	HandleError(w, err)
}

func contact(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "contact.gohtml", nil)
	HandleError(w, err)
}

func apply(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("in apply with method %s url %s\n", req.Method, req.URL)
	if req.Method == http.MethodGet {
		// code here
		fmt.Println("Handle get")
		err := tpl.ExecuteTemplate(w, "apply.gohtml", nil)
		HandleError(w, err)
		return
	} else if req.Method == http.MethodPost {
		// code here
		fmt.Println("Hanlde post")
		err := tpl.ExecuteTemplate(w, "applyProcess.gohtml", nil)
		HandleError(w, err)
		return
	} else {
		return
	}
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
