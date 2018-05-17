package intentHandlers

import (
	"github.com/BloodyRainer/articlePrice/df"
	"context"
	engLog "google.golang.org/appengine/log"
)

func RespondToNamePlayerOne(ctx context.Context, dfReq df.Request) (*df.Response, error) {

	engLog.Infof(ctx, "test0")

	nameP1 := dfReq.QueryResult.QueryText

	gs := df.GameSession{
		NamePlayerOne: nameP1,
	}

	payload := df.MakeSimpleRespPayload(true,
		"<speak>Ok, danke "+nameP1+"! Wie ist der Name von Spieler Nummer 2?</speak>",
		"Ok, danke "+nameP1+"! Wie ist der Name von Spieler 2?")

	engLog.Infof(ctx, "test1")

	resp := &df.Response{
		Source:  source,
		Payload: payload,
		OutputContexts: []df.Context{
			df.MakeOutputContext("game_session", 5, gs.ToParameters(), dfReq),
			df.MakeOutputContext("name_player_two", 3, nil, dfReq),
		},
	}

	engLog.Infof(ctx, "test2")

	return resp, nil

}
