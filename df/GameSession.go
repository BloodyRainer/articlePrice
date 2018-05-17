package df

import (
	"encoding/json"
	"strconv"
)

type GameSession struct {
	NamePlayerOne   string `json:"name-player-one"`
	NamePlayerTwo   string `json:"name-player-two"`
	PointsPlayerOne int    `json:"points-player-one"`
	PointsPlayerTwo int    `json:"points-player-two"`
	Round           int    `json:"round"`
}

func (rcv GameSession) ToParameters() []byte {
	params := MakeParameters("name-player-one", rcv.NamePlayerOne)
	params = AppendParameter(params, "name-player-two", rcv.NamePlayerTwo)
	params = AppendParameter(params, "points-player-one", strconv.Itoa(rcv.PointsPlayerOne))
	params = AppendParameter(params, "points-player-two", strconv.Itoa(rcv.PointsPlayerTwo))
	params = AppendParameter(params, "round", strconv.Itoa(rcv.Round))

	return params
}

func makeGameSessionFromContextParameters(parameters []byte) (GameSession, error) {
	gs := GameSession{}

	err := json.Unmarshal(parameters, &gs)

	return gs, err
}
