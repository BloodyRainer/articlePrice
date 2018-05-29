package server

import (
	"google.golang.org/appengine"
	"net/http"
	engLog "google.golang.org/appengine/log"
	"github.com/BloodyRainer/articlePrice/df"
	"github.com/BloodyRainer/articlePrice/intentHandlers"
	"encoding/json"
)

type fullfilmentHandler struct{}

func (rcv *fullfilmentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var dfReq *df.Request
	var dfRes *df.Response
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

		dfReq, err = df.MakeDfRequest(body)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}

	// dispatch intents in order to assemble dialogflow-response
	intent := dfReq.QueryResult.Intent.DisplayName
	engLog.Infof(ctx, "intent-name is: "+intent)

	switch intent {
	case "start_guess":
		dfRes, err = intentHandlers.AskRandomArticle(ctx, dfReq)
		if err != nil {
			engLog.Errorf(ctx, "failed to ask random article: "+err.Error())
			http.Error(w, err.Error(), 500)
			return
		}
	case "say_price":
		dfRes, err = intentHandlers.RespondToPriceGuess(ctx, *dfReq)
		if err != nil {
			engLog.Warningf(ctx, "failed to evaluate input: "+err.Error())
			dfRes = intentHandlers.AskForNewInput()
		}

	default:
		engLog.Errorf(ctx, "unknown intent: "+intent)
		return
	}

	// respond to dialogflow
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(dfRes)
}
