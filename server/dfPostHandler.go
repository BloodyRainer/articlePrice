package server

import (
	"google.golang.org/appengine"
	"net/http"
	"encoding/json"
	engLog "google.golang.org/appengine/log"
	"github.com/BloodyRainer/articlePrice/dialog"
	"io/ioutil"
	"context"
	"errors"
)

type articleHandler struct{}

func (rcv *articleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var dfReq *dialog.DfRequest
	var dfRes *dialog.DfResponse
	var err error

	ctx := appengine.NewContext(r)

	// check and log dialogflow-post-request
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

	// dispatch intents in order to assemble dialogflow-response
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
		dfRes, err = respondToPriceGuess(*dfReq)
		if err != nil {
			engLog.Warningf(ctx, "failed to evaluate input"+err.Error())

			dfRes = askForNewInput()
		}
	} else {
		engLog.Errorf(ctx, "unknown intend")
	}

	// respond to dialogflow
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(dfRes)
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
