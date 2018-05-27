package intentHandlers

import "github.com/BloodyRainer/articlePrice/df"

func RespondToNamePlayerTwo(dfReq df.Request) (*df.Response, error) {

	nameP2 := dfReq.QueryResult.QueryText

	gs, err := df.MakeGameSessionFromDfRequest(dfReq)
	if err != nil {
		return nil, err
	}

	gs.NamePlayerTwo = nameP2

	payload := df.MakeSimpleRespPayload(true,
		"<speak>Alles klar! der Name von Spieler 2 ist "+nameP2+"! Seid ihr bereit für die erste Runde?</speak>",
		"Alles klar, der Name von Spieler 2 ist "+nameP2+"! Seid ihr bereit für die erste Runde?")

	resp := &df.Response{
		Source:  source,
		Payload: payload,
		OutputContexts: []df.Context{
			df.MakeOutputContext("ask_first_player", 3, gs.ToParameters(), dfReq),
			df.MakeOutputContext("game_session", 5, gs.ToParameters(), dfReq),
		},
		//FollowupEventInput: &df.EventInput{
		//	Name: "quiz_turn_1_question",
		//	Parameters: df.MakeParameters("event-input", "value_of_event_input"),
		//},
	}

	return resp, nil

}
