package main

import (
	"html/template"
	"net/http"
)
type grocerylist []string
var tpl *template.Template
var g grocerylist
func main() {
	g = grocerylist{"okra", "Fish", "Garri", "maggi"}
	tpl = template.Must(template.ParseFiles("index.html"))
	http.HandleFunc("/grocery", listHandler)
	http.ListenAndServe(":4001", nil)
}

func listHandler(w http.ResponseWriter, r *http.Request){
	err := tpl.ExecuteTemplate(w, "index.html", g)
	if err != nil {
		http.Error(w, "Template Error: "+err.Error(), http.StatusInternalServerError)
	}
}