package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// Pointer to a template is "a container holding all the templates parsed"
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
