package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

var tmpl *template.Template

var m = 9
var fm = template.FuncMap{
	"formMdy": formMdy,
}

func init() {
	tmpl = template.Must(template.New("").Funcs(fm).ParseGlob("./*"))
}

func formMdy(t time.Time) string {
	return t.Format("01-02-2006")
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is root.")
}

func handleDog(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is dog.")

}

func handleMe(w http.ResponseWriter, r *http.Request) {

	// NEW FINDING:
	// if you want to render content which is in a struct in template
	// if the rendered content having struct in it
	// the variable name in that struct should be able to be exported
	// i.e. the first alpha should be uppercase
	myInfo := struct {
		Name   string
		Age    int
		Gender string
		Birth  time.Time
	}{
		"Hongyi",
		22,
		"male",
		time.Date(1997, time.March, 19, 0, 0, 0, 0, time.Local),
	}
	fmt.Println("trying to render")
	err := tmpl.ExecuteTemplate(w, "temp.html", myInfo)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.Handle("/", http.HandlerFunc(handleMain))
	http.Handle("/dog/", http.HandlerFunc(handleDog))
	http.Handle("/me/", http.HandlerFunc(handleMe))

	http.ListenAndServe(":8080", nil)
}
