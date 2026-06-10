package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/ascii-art", Ascii)
	http.ListenAndServe(":8080", nil)
}

var tpl = template.Must(template.ParseFiles("templates/index.html"))

type PageData struct {
	Title  string
	Result string
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Path Not Found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Available", http.StatusMethodNotAllowed)
		return
	}
	data := PageData{Title: "ASCII-ART-GENERATOR"}
	tpl.Execute(w, data)
}

func Ascii(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		http.Error(w, "Path Not Found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Available", http.StatusMethodNotAllowed)
		return
	}
	text := r.FormValue("text")
	banner := r.FormValue("banner")
	temp, err := Asciiart(text, banner)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := PageData{Title: "Text to ASCII Art Generator: Create ASCII Art from Text", Result: temp}
	if r.Header.Get("HX-Request") == "true"{
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(data.Result))
		return
	}
	if err := tpl.Execute(w, data); err != nil {
		http.Error(w, "template execution failed", http.StatusInternalServerError)
		return
	}

}
