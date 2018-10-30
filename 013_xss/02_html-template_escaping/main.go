package main

import (
	"html/template"
	"log"
	"os"
)

type Page struct {
	Title   string
	Heading string
	Input   string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	// 使用html/template能够防止
	// 传给template的字符串被浏览器识别为html代码而
	// 影响网页本身
	// 比如字符串中有>、<等特殊符号，会被自动转为&gt, &lt等html转义符
	// 之所以这样做是为了安全，防止传入字符串影响网页的正常显示

	home := Page{
		Title:   "Escaped",
		Heading: "Danger is escaped with html/template",
		Input:   `<script>alert("Yow!");</script>`,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", home)
	if err != nil {
		log.Fatalln(err)
	}
}
