package main

import (
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
		Title  string
		Result string
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.Error(w, "Not Found", http.StatusNotFound)
        return
    }
    data := PageData{Title: "ASCII Art Generator"}

    if r.Method == "POST" {
		banner := r.FormValue("banner")
		if banner ==""{
			http.Error(w, "Wrong Request: Invalid banner selected", http.StatusBadRequest)
			return
		}
		text := r.FormValue("text")
		if text == ""{
			http.Error(w, "Bad Request: Empty Input", http.StatusBadRequest)
			return 
		}
        data.Result = "you typed: " + text + " " + "with banner: " + banner
    }

    tmpl, err := template.ParseFiles("templates/index.html")
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", homeHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
