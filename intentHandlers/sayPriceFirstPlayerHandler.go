package intentHandlers

import (
	"github.com/BloodyRainer/articlePrice/df"
	"strconv"
	"context"
	engLog "google.golang.org/appengine/log"
)

func SavePriceFristPlayerAskSecondPlayer(ctx context.Context, dfReq df.Request) (*df.Response, error) {

	gs, err := df.MakeGameSessionFromDfRequest(dfReq)
	if err != nil {
		return nil, err
	}

	priceFP, err := df.GetPriceAnswer(dfReq, "first_player_answer")
	if err != nil {
		return nil, err
	}

	engLog.Infof(ctx, "price first player: %v", priceFP)

	gs.SaveFirstPriceAnswer(priceFP)

	resp := makeQuestionSecondPlayer(dfReq, gs)

	return resp, nil
}

func makeQuestionSecondPlayer(dfReq df.Request, gs *df.GameSession) *df.Response {

	points := strconv.Itoa(int(gs.GetPointsOfCurrentTurn()))
	pName := gs.GetSecondPlayerName()

	payload := df.MakeSimpleRespPayload(true,
		"<speak> " + pName+", für "+ points +" Punkte, wie ist dein Tipp?</speak>",
		pName+" - für "+ points +" Punkte!")

	resp := &df.Response{
		Source:  source,
		Payload: payload,
		OutputContexts: []df.Context{
			df.MakeOutputContext("second_player_answer", 3, nil, dfReq),
			df.MakeOutputContext("game_session", 5, gs.ToParameters(), dfReq),
			df.MakeOutputContext("first_player_answer", 0, nil, dfReq),
		},
	}

	return resp

}

