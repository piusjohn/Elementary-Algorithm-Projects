package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)
type PageData struct{
	Title string
	Result string
}
var tpl = template.Must(template.ParseFiles("templates/index.html"))
func HomeHandler(w http.ResponseWriter, req *http.Request){
	if req.URL.Path != "/"{
		http.Error(w, "route not found", http.StatusNotFound)
		return
	}
	data := PageData{Title: "ASCII ART GENERATOR"}
	if req.Method == http.MethodGet{
		if err := tpl.ExecuteTemplate(w, "index.html", data); err != nil{
		http.Error(w, "Sever error", http.StatusInternalServerError)
		return
	}
	return
}
	if req.Method == http.MethodPost{
	text := req.FormValue("text")
	banner := req.FormValue("banner")

	bannerPath := fmt.Sprintf("banners/%s.txt", banner)
	bannerMap, err := LoadBanner(bannerPath)
	if err != nil {
		http.Error(w, "Banner does not exist", http.StatusInternalServerError)
		return
	}
	data.Result = GenerateArt(text, bannerMap)
	 w.Write([]byte(data.Result))
	// if err := tpl.ExecuteTemplate(w, "index.html", data); err != nil{
	// 	http.Error(w, "Sever error", http.StatusInternalServerError)
	// 	return
	// }
	return
}
	http.Error(w, "Method not available", http.StatusMethodNotAllowed)

}

func main(){
	http.HandleFunc("/", HomeHandler)
	err := http.ListenAndServe(":5000", nil)
	log.Fatal(err)
}