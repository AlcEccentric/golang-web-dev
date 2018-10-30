package main

import (
	"log"
	"os"
	"text/template"
)

// set in global scope
// so templates only be parsed once
var tpl *template.Template

// 1. init func is used for initialization, it is a func embedded in Go
// 2. Every package could have multiple init()
// 3. The order of multiple init() is not deterministic
// 4. init() will be called automatically before main()

func init() {
	// Use Must to make sure that
	// all the files parsed correctly
	// else, it would panics
	// so you don't need to check correctness by hard code
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
