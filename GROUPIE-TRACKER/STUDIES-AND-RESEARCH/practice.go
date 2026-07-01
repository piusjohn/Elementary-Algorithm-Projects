package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
)

var tpl = template.Must(template.ParseGlob("templates/*.html"))

type Artist struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Image         string    `json:"image"`
	Members       []string  `json:"members"`
	CreationDate  int       `json:"creationDate"`
	FirstAlbum    string    `json:"firstAlbum"`
	Locations     string    `json:"locations"`
	ConcertDates  string    `json:"concertDates"`
	Relations     string    `json:"relations"`

}

type LocationIndex struct{ 
	ID          int         `json:"id"`	
	Locations   []string    `json:"locations"`
	Dates       string      `json:"dates"`
}

type Locations struct {
	Index []LocationIndex  `json:"index"`
}

type DateIndex struct{
	ID     int     `json:"id"`
	Dates  []string `json:"dates"`
}

type Dates struct {
	Index []DateIndex `json:"index"`
}

type RelationIndex struct {
	ID            int                  `json:"id"`
	Dateslocation map[string][]string  `json:"datesLocations"`
}

type Relation struct {
	Index  []RelationIndex `json:"index"`
}

type ArtistPageData struct {
	Artist Artist
	Location LocationIndex
	Date  DateIndex
	Relation RelationIndex
}

func fetch[T any](url string) (T, error) {
	var zero T
    resp, err := http.Get(url)
	if err != nil {
		return zero, fmt.Errorf("error parsing the Artist json")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return zero, fmt.Errorf("error streaming resp.body")
	}

	var data T
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}
	return data, nil
}
var artists []Artist
var locations Locations
var dates  Dates
var relation Relation
func main() {
	var err error
	artists, err = fetch[[]Artist]("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil{
		log.Fatal(err)
	}

	locations, err = fetch[Locations]("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		log.Fatal(err)
	}

	dates, err = fetch[Dates]("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		log.Fatal(err)
	}

	relation, err = fetch[Relation]("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/home", handleHome)
	http.HandleFunc("/artist/{id}", handleArtist)
	log.Fatal(http.ListenAndServe(":8080", nil))
	
}

func handleHome(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w, "index.html", artists)
}

func handleArtist(w http.ResponseWriter, r *http.Request){
	v := r.PathValue("id")
	id, err := strconv.Atoi(v)
	if err != nil{
		fmt.Println("id conversion failed")
		return
	}
	// for i := 0; i < len(artists); i++
	for _, a := range artists{
		if a.ID == id {
			tpl.ExecuteTemplate(w, "artist.html", a)
			return
		}
	}
	http.Error(w, "message: Artist not found", http.StatusNotFound)
}