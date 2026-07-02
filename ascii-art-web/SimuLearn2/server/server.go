package server

import (
	"ascii-art-web/handlers"
	"log"
	"net/http"
)

func RegisterHandlers() {
	http.HandleFunc("/", handlers.Home)
}

func StartServer() {
	RegisterHandlers()
	log.Print("Listening on :http://localhost:3000/")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}