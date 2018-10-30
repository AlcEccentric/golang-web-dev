package main

// When using func in template.
// The funcs passed should only be used to modify the text
// They should not get access any data outside that text
// Since we should not violate the MVC model
import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

// create a FuncMap to register functions.
// "uc" is what the func will be called in the template
// "uc" is the ToUpper func from package strings
// "ft" is a func I declared
// "ft" slices a string, returning the first three characters
// then you can call these func in template by their corresponding key
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {

	// Use .New("").Funcs(FuncMapVariable). ... to assign FuncMap to all the templates
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))

	// NOTICE:
	// Following will not work, because:
	// Go statically parses files
	// When fm is attached to the tpl using method Funcs,
	// the files have already been parsed
	// So, fm will not be avaiable to parsed files
	// The FuncMap must be attached to files before parsing
	// That is the reason to use template.New("").Funcs(fm).ParseGlob/Files(...)
	// Which generates a empty temp, add FuncMap and then parse all files to and attach them to empty temp
	// tpl = template.Must(template.ParseFiles("xxx.html"))
	// tpl = tpl.Funcs(fm)
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
}

func main() {

	b := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	g := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	m := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed.",
	}

	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	c := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	sages := []sage{b, g, m}
	cars := []car{f, c}

	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		sages,
		cars,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
