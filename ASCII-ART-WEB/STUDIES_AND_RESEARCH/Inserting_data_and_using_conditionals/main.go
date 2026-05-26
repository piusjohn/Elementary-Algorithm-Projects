package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

type User struct{
	Name string
	language string
	member bool
}
var U User
func main() {
	U = User{Name: "peter", language: "English", member: false}
	tpl, _ = tpl.ParseGlob("templates/*.html") 
	http.HandleFunc("/welcome", welcome)
	http.ListenAndServe(":4001", nil)

}

func welcome(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w, "welcome", U)
}