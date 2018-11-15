package main

import (
	"html/template"
	"net/http"

	"github.com/satori/go.uuid"
)

type user struct {
	UserName string
	First    string
	Last     string
}

var tpl *template.Template
var dbUsers = map[string]user{}      // user ID, user
var dbSessions = map[string]string{} // session ID, user ID

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

// this func is responsible for giving visitor cookie,
// allowing them to start a session and show their info
func index(w http.ResponseWriter, req *http.Request) {

	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}

	// if the user exists already, get user
	var u user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}

	// process form submission
	// PS: I think there is a fault here.
	// in the index.gohtml we should not show the form if the visitor has started a session
	// if we show that form again, and the visitor does post,
	// the original session will be erased by the visitor
	// and be replaced a new one.
	// I dont think it is reasonable
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		u = user{un, f, l}
		dbSessions[c.Value] = un
		dbUsers[un] = u
	}

	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

// using this function, we make /bar unvisitable by visitors
// who have not started a session with our domain
func bar(w http.ResponseWriter, req *http.Request) {

	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		// if the visitor first time visit the site
		// redirect him to index page and dont show bar
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	un, ok := dbSessions[c.Value]
	if !ok {
		// if the visitor have visited the site
		// but has not started a session
		// redirect him to index page and dont show bar
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	u := dbUsers[un]
	// only if the visitor has visited the site and has started a sesion,
	// we will show him this page
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

// map examples with the comma, ok idiom
// https://play.golang.org/p/OKGL6phY_x
// https://play.golang.org/p/yORyGUZviV
