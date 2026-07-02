package handlers

import (
	"html/template"
	"net/http"
)
var tpl = template.Must(template.ParseGlob("template/*.html"))
func Home(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/"{
		renderErrorTemplate(w, http.StatusNotFound, "template/errors/400.html")
		return
	}
	if r.Method != http.MethodGet{
		renderErrorTemplate(w, http.StatusMethodNotAllowed, "templates/errors/400.html")
		return
	}
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func renderErrorTemplate(w http.ResponseWriter, statusCode int, templatePath string){
	tpl := template.Must(template.ParseFiles(templatePath))
	w.WriteHeader(statusCode)
	tpl.Execute(w, nil)
}