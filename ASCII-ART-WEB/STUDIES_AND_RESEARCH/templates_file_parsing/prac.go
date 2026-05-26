package main

import (
	//"fmt"
	"net/http"
	"html/template"
)
var tpl *template.Template


func home(w http.ResponseWriter, r *http.Request){
	tpl.Execute(w, nil)
}
func main() {
	//var err error
	tpl, _ = template.ParseFiles("index.html")
	// if err != nil{
	// 	panic(err)
	// }
	http.HandleFunc("/home", home)
	http.ListenAndServe(":4001", nil)
}