package web

import (
	"html/template"
	"log"
	"net/http"
	"thinkPrinter/tools"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		log.Println(err)
	}
	err = tmpl.Execute(w, "Hello World!")
	tools.OutputLog(r)
}
