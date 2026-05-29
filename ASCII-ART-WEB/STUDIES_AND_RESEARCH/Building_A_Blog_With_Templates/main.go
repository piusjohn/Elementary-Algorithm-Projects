package main

import (
	"database/sql"
	"html/template"
	"github.com/go-chi/chi/v5"
	//"github.com/go-chi/chi/v5/middleware"
	_ "github.com/mattn/go-sqlite3"
)

var router *chi.Mux
var db *sql.DB

type Article struct{
	ID int  `json:"id"`
	Title string `json:"title"`
	Content template.HTML `json:"content"`
}