package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// Understand *Template(tpl) as a container,
	// which contains all the files parsed
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

// Do not use the above code in production
// We will learn about efficiency improvements soon!
