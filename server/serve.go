package server

import (
	"net/http"
	"io/ioutil"
	engLog "google.golang.org/appengine/log"
	"context"
	"log"
)

const port = ":8080"

func Start() {
	log.Println("Starting Server on port " + port + "...")
	http.Handle("/webhook", &articleHandler{})
	http.ListenAndServe(port, nil)
}

func logPostRequest(ctx context.Context, req *http.Request) {

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		engLog.Errorf(ctx, err.Error())
	}

	engLog.Debugf(ctx, "logging post-body...")

	bodyStr := string(body)

	if bodyStr != "" {
		engLog.Debugf(ctx, bodyStr)
	} else {
		engLog.Debugf(ctx, "body string is empty")
	}

}
