package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.New("index.html")

	tpl.Funcs(template.FuncMap{
		"lastIndex": func(s []string) string {
			lastIndex := len(s) - 1
			return s[lastIndex]
		},
	})

	_, err := tpl.ParseFiles("index.html")

	if err != nil {
		log.Fatal(err)
	}

}

var g []string

func main() {
	g = []string{"milk", "egg", "custard", "banana"}
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", g)
}
