package server

import (
	"net/http"
	"log"
)

const port = ":8080"

func Start() {
	log.Println("Starting Server on port " + port + "...")
	http.Handle("/webhook", &articleHandler{})
	http.ListenAndServe(port, nil)
}
