package main

import (
	"html/template"
)

var tpl *template.Template

type User struct{
	Name string
	language string
	member bool
}
var U User
func main() {
	u = User{Name: "peter", language: "English", member: false}
	tpl, _ = tpl.ParseGlob("templates/*.html")
	
}