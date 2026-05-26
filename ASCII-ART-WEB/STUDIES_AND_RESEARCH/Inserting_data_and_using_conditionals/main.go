package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

type User struct{
	Name string
	Language string
	Member bool
}
var U User
func main() {
	U = User{Name: "peter", Language: "Madarin", Member: true}
	tpl, _ = template.ParseGlob("templates/*.html") 
	http.HandleFunc("/welcome", welcome)
	http.ListenAndServe(":4001", nil)

}

func welcome(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w, "index2.html", U)
}