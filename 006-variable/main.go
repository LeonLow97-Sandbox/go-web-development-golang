package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	// only can pass one piece of data
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", "Release self-focus; embrace other-focus")
	if err != nil {
		log.Fatalln(err)
	}
}
