package intentHandlers

import (
	"github.com/BloodyRainer/articlePrice/df"
	"strings"
)

func RespondToNamePlayerOne(dfReq df.Request) (*df.Response, error) {

	nameP1 := strings.Title(dfReq.QueryResult.QueryText)

	gs := df.GameSession{
		NamePlayerOne: nameP1,
		Turn: 0,
		PricesPlayerOne: []float64{0.00},
		PricesPlayerTwo: []float64{0.00},
	}

	payload := df.MakeSimpleRespPayload(true,
		"<speak>Ok, wie ist der Name von Spieler 2?</speak>",
		"Wie ist der Name von Spieler 2?")

	resp := &df.Response{
		Source:  source,
		Payload: payload,
		OutputContexts: []df.Context{
			df.MakeOutputContext("game_session", 5, gs.ToParameters(), dfReq),
			df.MakeOutputContext("name_player_two", 3, nil, dfReq),
			df.MakeOutputContext("name_player_one", 0, nil, dfReq),
		},
	}

	return resp, nil

}
