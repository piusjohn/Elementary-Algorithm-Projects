package handlers

import (
	asciiart "ascii-art-web/ascii-art"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func renderTemplateWithData(w http.ResponseWriter, tmplName string, data interface{}){
	lp := filepath.Join("templates", "layout.html")	
	fp := filepath.Join("templates", tmplName)

	info, err := os.Stat(fp)
		if err != nil || info.IsDir(){
			renderErrorTemplate(w, http.StatusInternalServerError, "errors/500.html")
			return
	}
	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		renderErrorTemplate(w, http.StatusInternalServerError, "errors/500.html")
		log.Print(err.Error())
		return
	}
	err = tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		renderErrorTemplate(w, http.StatusInternalServerError, "errors/500.html")
		log.Print(err.Error())
	}

}

func ServeTemplate(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/index.html" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	lp := filepath.Join("templates", "layout.html")

	

	var fp string
	if r.URL.Path == "/" {
		fp = filepath.Join("templates", "index.html")
	} else {
		fp = filepath.Join("templates", filepath.Clean(r.URL.Path))
	}

	info, err := os.Stat(fp)
	if err != nil || info.IsDir(){
		renderErrorTemplate(w, http.StatusNotFound, "errors/404.html")
		return
	}
	
	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		renderErrorTemplate(w, http.StatusInternalServerError, "errors/500.html")
		log.Print(err.Error())
		return
	}
	
	if err = tmpl.ExecuteTemplate(w, "layout", nil); err != nil{
		renderErrorTemplate(w, http.StatusInternalServerError, "errors.500.html")
		log.Print(err.Error())
	}

}

func HandleAsciiArt(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	artStyle := r.FormValue("banner")
	userText := r.FormValue("text")
	result, err := asciiart.Asciiart(userText, artStyle)
	if err != nil{
		renderErrorTemplate(w, http.StatusInternalServerError, "errors/500.html")
		log.Println(err.Error())
		return
	}
	data := struct {
		ASCIIArtResult string
	}{

		ASCIIArtResult: result,
	}
	tmplName := "index.html"
	renderTemplateWithData(w, tmplName, data)
}

func renderErrorTemplate(w http.ResponseWriter, statusCode int, templatePath string) {
	w.WriteHeader(statusCode)
	renderTemplateWithData(w, templatePath, nil)
}