package server

import (
	"net/http"
	"github.com/BloodyRainer/articlePrice/search"
	"encoding/json"
	"io/ioutil"
	engLog "google.golang.org/appengine/log"
	"google.golang.org/appengine"
	"context"
	"log"
)

type articleHandler struct{}

func (rcv *articleHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	ctx := appengine.NewContext(req)

	a, err := search.GetRandomArticle(req)
	if err != nil {
		engLog.Errorf(ctx, err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(err.Error()))
		return
	}

	if req.Method == http.MethodPost {
		logPostRequest(ctx, req)
	}

	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	enc := json.NewEncoder(res)
	enc.Encode(a)

	engLog.Infof(ctx, "responded Article:", *a)

}

func Start() {
	log.Println("Starting Server...")
	http.Handle("/getArticle", &articleHandler{})
	http.ListenAndServe(":8080", nil)
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
