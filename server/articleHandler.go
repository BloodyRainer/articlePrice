package server

import (
	"google.golang.org/appengine"
	"net/http"
	"encoding/json"
	"github.com/BloodyRainer/articlePrice/search"
	engLog "google.golang.org/appengine/log"
	"github.com/BloodyRainer/articlePrice/dialogflow"
	"io/ioutil"
	"context"
	"errors"
)

type articleHandler struct{}

func (rcv *articleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var dfReq *dialogflow.DfRequest
	var err error

	ctx := appengine.NewContext(r)

	if r.Method == http.MethodPost {

		body, err := readPostBody(ctx, r)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		logPostBody(ctx, body)

		dfReq, err = dialogflow.MakeDfRequest(body)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}

	//TODO: GetArctileByNumber if request contains ArticleNr
	a, err := search.GetRandomArticle(r)
	if err != nil {
		engLog.Errorf(ctx, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	dr := dialogflow.MakeArticleNameResponse(*a)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(dr)

	engLog.Infof(ctx, "dfReq responseId: " + dfReq.ResponseId)
	//engLog.Infof(ctx, "dfReq languageCode: " +dfReq.QueryResult.LanguageCode)

}

func readPostBody(ctx context.Context, r *http.Request) ([]byte, error) {
	if r.Body == nil {
		return nil, errors.New("body of post request is nil")
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.New("unable to read body: " + err.Error())
	}

	return body, nil
}

func logPostBody(ctx context.Context, body []byte) {

	bodyStr := string(body)

	if bodyStr != "" {
		engLog.Debugf(ctx, "req-body: " + bodyStr)
	} else {
		engLog.Debugf(ctx, "body string is empty")
	}

}
