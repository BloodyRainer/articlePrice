package intentHandlers

import (
	"github.com/BloodyRainer/articlePrice/df"
)

func RespondToNamePlayerOne(dfReq df.Request) (*df.Response, error) {

	nameP1 := dfReq.QueryResult.QueryText

	gs := df.GameSession{
		NamePlayerOne: nameP1,
		Turn: 0,
		PricesPlayerOne: []float64{0.00},
		PricesPlayerTwo: []float64{0.00},
	}

	payload := df.MakeSimpleRespPayload(true,
		"<speak>Ok, der Name von Spieler <say-as interpret-as='cardinal'>1</say-as> ist "+nameP1+"! Wie ist der Name von Spieler 2?</speak>",
		"Ok, der Name von Spieler 1 ist "+nameP1+"! Wie ist der Name von Spieler 2?")

	resp := &df.Response{
		Source:  source,
		Payload: payload,
		OutputContexts: []df.Context{
			df.MakeOutputContext("game_session", 5, gs.ToParameters(), dfReq),
			df.MakeOutputContext("name_player_two", 3, nil, dfReq),
		},
	}

	return resp, nil

}
