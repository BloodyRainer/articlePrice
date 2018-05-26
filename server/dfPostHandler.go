package server

import (
	"google.golang.org/appengine"
	"net/http"
	engLog "google.golang.org/appengine/log"
	"github.com/BloodyRainer/articlePrice/df"
	"io/ioutil"
	"context"
	"errors"
	"github.com/BloodyRainer/articlePrice/intentHandlers"
	"encoding/json"
)

type articleHandler struct{}

func (rcv *articleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

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
		dfRes, err = intentHandlers.RespondToPriceGuess(*dfReq)
		if err != nil {
			engLog.Warningf(ctx, "failed to evaluate input: "+err.Error())
			dfRes = intentHandlers.AskForNewInput()
		}

	// two-player mode intents
	case "say_name_player_one":
		dfRes, err = intentHandlers.RespondToNamePlayerOne(*dfReq)
		if err != nil {
			engLog.Errorf(ctx, "failed to answer to name of player one: "+err.Error())
			http.Error(w, err.Error(), 500)
			return
		}
	case "say_name_player_two":
		dfRes, err = intentHandlers.RespondToNamePlayerTwo(*dfReq)
		if err != nil {
			engLog.Errorf(ctx, "failed to answer to name of player one: "+err.Error())
			http.Error(w, err.Error(), 500)
			return
		}
	case "quiz_ask_question_firstPlayer":
		dfRes, err = intentHandlers.AskArticleQuestionFirstPlayer(ctx, *dfReq)
		if err != nil {
			engLog.Errorf(ctx, "failed to ask random article: "+err.Error())
			http.Error(w, err.Error(), 500)
			return
		}
	case "quiz_answer_firstPlayer":
		//TODO:
	default:
		engLog.Errorf(ctx, "unknown intent: "+intent)
		return
	}

	// respond to dialogflow
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
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
