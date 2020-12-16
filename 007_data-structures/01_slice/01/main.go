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

	sages := []string{"Gandhi", "MLK", "Buddha", "Jesus", "Muhammad"}
	f, err := os.Create("sysFile.html")
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	err1 := tpl.Execute(os.Stdout, sages)
	err2 := tpl.Execute(f, sages)
	if err1 != nil || err2 != nil {
		log.Fatalln(err1, err2)
	}
}
