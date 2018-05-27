package intentHandlers

import (
	"github.com/BloodyRainer/articlePrice/df"
	"math"
	"strconv"
	"bytes"
	"context"
)

func SavePriceSecondPlayerAndResultsOfTurn(ctx context.Context, dfReq df.Request) (*df.Response, error) {

	gs, err := df.MakeGameSessionFromDfRequest(dfReq)
	if err != nil {
		return nil, err
	}

	priceSP, err := df.GetPriceAnswer(dfReq, "second_player_answer")
	if err != nil {
		return nil, err
	}

	gs.SaveSecondPriceAnswer(priceSP)

	resp := makeResultsCurrentTurn(ctx, dfReq, gs)

	return resp, nil

}

func makeResultsCurrentTurn(ctx context.Context, dfReq df.Request, gs *df.GameSession) *df.Response {

	w, l, wp, lp := calculateResults(ctx, gs)

	var tts bytes.Buffer
	tts.WriteString("<speak>")
	tts.WriteString(w)
	tts.WriteString(" lag näher am tatsächlichen Preis von ")
	tts.WriteString(df.PriceInEuroTTS(gs.CurrentArticlePrice))
	tts.WriteString(" und bekommt <say-as interpret-as='cardinal'>")
	tts.WriteString(wp)
	tts.WriteString("</say-as> Punkte! ")
	tts.WriteString(l)
	tts.WriteString(" erhält <say-as interpret-as='cardinal'>")
	tts.WriteString(lp)
	tts.WriteString("</say-as> Punkte!")
	tts.WriteString(" Bereit für die nächste Runde?</speak>")

	var text bytes.Buffer
	text.WriteString("")
	text.WriteString(w)
	text.WriteString(" lag näher am tatsächlichen Preis von ")
	text.WriteString(df.PriceInEuroText(gs.CurrentArticlePrice))
	text.WriteString(" und bekommt ")
	text.WriteString(wp)
	text.WriteString(" Punkte! ")
	text.WriteString(l)
	text.WriteString(" erhält ")
	text.WriteString(lp)
	text.WriteString(" Punkte!")
	text.WriteString(" Bereit für die nächste Runde?")

	// reset article for next turn
	gs.CurrentArticlePrice = -1
	gs.CurrentArticleName = ""
	gs.CurrentArticleImgUrl = ""
	gs.CurrentArticleNumber = ""

	payload := df.MakeSimpleRespPayload(true, tts.String(), text.String())

	resp := &df.Response{
		Source:  source,
		Payload: payload,
		OutputContexts: []df.Context{
			df.MakeOutputContext("ask_first_player", 3, nil, dfReq),
			df.MakeOutputContext("game_session", 5, gs.ToParameters(), dfReq),
		},
	}

	return resp

}

func calculateResults(ctx context.Context, gs *df.GameSession) (winner string, loser string, pWinner string, pLoser string) {

	pointsWinner := gs.GetPointsOfCurrentTurn()
	pointsLoser := pointsWinner / 2

	diffPlayerOne := math.Abs(gs.CurrentArticlePrice - gs.GetPriceGuessOfCurrentTurnPlayerOne())
	diffPlayerTwo := math.Abs(gs.CurrentArticlePrice - gs.GetPriceGuessOfCurrentTurnPlayerTwo())

	if diffPlayerOne > diffPlayerTwo {
		gs.PointsPlayerTwo = + pointsWinner
		gs.PointsPlayerOne = + pointsLoser
		return gs.NamePlayerTwo, gs.NamePlayerTwo, strconv.Itoa(int(pointsWinner)), strconv.Itoa(int(pointsLoser))
	} else if diffPlayerOne == diffPlayerTwo {
		// TODO:
	}

	gs.PointsPlayerOne = + pointsWinner
	gs.PointsPlayerTwo = + pointsLoser

	return gs.NamePlayerOne, gs.NamePlayerTwo, strconv.Itoa(int(pointsWinner)), strconv.Itoa(int(pointsLoser))

}
