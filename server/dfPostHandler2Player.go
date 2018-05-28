package server

import (
	"net/http"
	"github.com/BloodyRainer/articlePrice/df"
	"google.golang.org/appengine"
	engLog "google.golang.org/appengine/log"
	"github.com/BloodyRainer/articlePrice/intentHandlers"
	"encoding/json"
)

type fullfilmentHandler2Player struct{}

func (rcv *fullfilmentHandler2Player) ServeHTTP(w http.ResponseWriter, r *http.Request) {

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
		dfRes, err = intentHandlers.SavePriceFristPlayerAskSecondPlayer(ctx, *dfReq)
		if err != nil {
			engLog.Errorf(ctx, "failed to process price answer of first player: "+err.Error())
			http.Error(w, err.Error(), 500)
			return
		}
	case "quiz_answer_secondPlayer":
		dfRes, err = intentHandlers.SavePriceSecondPlayerAndResultsOfTurn(ctx, *dfReq)
		if err != nil {
			engLog.Errorf(ctx, "failed to process price answer of second player: "+err.Error())
			http.Error(w, err.Error(), 500)
			return
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