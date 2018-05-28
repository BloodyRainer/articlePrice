package server

import (
	"net/http"
	"log"
)

const port = ":8080"

func Start() {
	log.Println("Starting Server on port " + port + "...")
	http.Handle("/webhook", &fullfilmentHandler{})
	http.Handle("/webhook2Player", &fullfilmentHandler2Player{})
	http.ListenAndServe(port, nil)
}
