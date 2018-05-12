package server

import (
	"google.golang.org/appengine"
	"net/http"
	"encoding/json"
	"github.com/BloodyRainer/articlePrice/search"
	engLog "google.golang.org/appengine/log"
	"github.com/BloodyRainer/articlePrice/dialog"
	"io/ioutil"
	"context"
	"errors"
	"strconv"
)

type articleHandler struct{}

func (rcv *articleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var dfReq *dialog.DfRequest
	var dfRes *dialog.DfResponse
	var err error

	ctx := appengine.NewContext(r)

	if r.Method == http.MethodPost {

		body, err := readPostBody(r)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		logPostBody(ctx, body)

		dfReq, err = dialog.MakeDfRequest(body)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}

	intent := dfReq.QueryResult.Intent.DisplayName
	engLog.Infof(ctx, "intent-name is: "+intent)

	if intent == "start_guess" {
		dfRes, err = askRandomArticle(ctx, dfReq)
		if err != nil {
			engLog.Errorf(ctx, err.Error())
			http.Error(w, err.Error(), 500)
			return
		}
	} else if intent == "say_price" {
		dfRes, err = respondToPriceGuess(ctx, *dfReq)
		if err != nil {
			engLog.Warningf(ctx, "failed to evaluate input" + err.Error())

			dfRes = askForNewInput()
		}
	} else {
		// should not happen...
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(dfRes)

}

func askRandomArticle(ctx context.Context, dfReq *dialog.DfRequest) (*dialog.DfResponse, error) {
	a, err := search.GetRandomArticle(ctx)

	engLog.Infof(ctx, "random article: " + a.String())

	if err != nil {
		return nil, err
	}

	resp := dialog.MakeArticleNameResponse(ctx, *a, *dfReq)

	return resp, nil
}

func askForNewInput() *dialog.DfResponse {
	return dialog.MakeNewInputResponse()
}

func respondToPriceGuess(ctx context.Context, dfReq dialog.DfRequest) (*dialog.DfResponse, error) {

	g, err := dialog.MakeGuessFromDfRequest(dfReq)
	if err != nil {
		return nil, err
	}

	gp := strconv.FormatFloat(g.PriceGuess, 'f', 2, 64)
	ap := strconv.FormatFloat(g.ActualPrice, 'f', 2, 64)

	engLog.Infof(ctx, "articleNumber is: "+g.ArticleNr)
	engLog.Infof(ctx, "articleName is: "+g.ArticleName)
	engLog.Infof(ctx, "guessed price is: "+gp)
	engLog.Infof(ctx, "actual price is: "+ap)

	resp := dialog.MakeEvaluatedResponse(g)

	return resp, nil
}

func readPostBody(r *http.Request) ([]byte, error) {
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
		engLog.Debugf(ctx, "req-body: "+bodyStr)
	} else {
		engLog.Debugf(ctx, "body string is empty")
	}

}
