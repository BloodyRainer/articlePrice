package server

import (
	"net/http"
	"log"
	"github.com/BloodyRainer/articlePrice/search"
)

const port = ":8080"

func Start() {
	log.Println("Starting Server on port " + port + "...")

	http.Handle("/webhook", &fullfilmentHandler{})
	http.Handle("/webhook2Player", &fullfilmentHandler2Player{})

	// test the full of articles one by one
	http.HandleFunc("/testArticles", search.IntegrationTestArticleListHandler)

	http.ListenAndServe(port, nil)
}
