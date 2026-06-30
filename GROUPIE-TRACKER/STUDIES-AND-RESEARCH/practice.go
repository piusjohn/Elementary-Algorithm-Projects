package main

import (
)

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

func main() {

}