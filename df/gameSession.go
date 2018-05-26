package df

import (
	"encoding/json"
	"strconv"
	"strings"
	"math"
)


type GameSession struct {
	NamePlayerOne        string    `json:"name-player-one"`
	NamePlayerTwo        string    `json:"name-player-two"`
	PricesPlayerOne      []float64 `json:"prices-player-one"`
	PricesPlayerTwo      []float64 `json:"prices-player-two"`
	PointsPlayerOne      float64   `json:"points-player-one"`
	PointsPlayerTwo      float64   `json:"points-player-two"`
	Turn                 float64   `json:"turn"`
	CurrentArticleName   string    `json:"current-article-name"`
	CurrentArticlePrice  string    `json:"current-article-price"`
	CurrentArticleNumber string    `json:"current-article-number"`
	CurrentArticleImgUrl string    `json:"current-article-img-url"`
}

func (rcv GameSession) ToParameters() []byte {

	params := MakeParameters("name-player-one", rcv.NamePlayerOne)
	params = AppendParameter(params, "name-player-two", rcv.NamePlayerTwo)
	params = AppendParameter(params, "prices-player-one", floatArray2String(rcv.PricesPlayerOne))
	params = AppendParameter(params, "prices-player-two", floatArray2String(rcv.PricesPlayerTwo))
	params = AppendParameter(params, "points-player-one", strconv.FormatFloat(rcv.PointsPlayerOne, 'f', 0, 64))
	params = AppendParameter(params, "points-player-two", strconv.FormatFloat(rcv.PointsPlayerTwo, 'f', 0, 64))
	params = AppendParameter(params, "turn", strconv.FormatFloat(rcv.Turn, 'f', 0, 64))
	params = AppendParameter(params, "current-article-name", rcv.CurrentArticleName)
	params = AppendParameter(params, "current-article-price", rcv.CurrentArticlePrice)
	params = AppendParameter(params, "current-article-number", rcv.CurrentArticleNumber)
	params = AppendParameter(params, "current-article-img-url", rcv.CurrentArticleImgUrl)

	return params
}

func MakeGameSessionFromDfRequest(dfReq Request) (GameSession, error) {

	contexts := dfReq.QueryResult.OutputContexts

	var idx int

	for i, c := range contexts {
		if strings.Contains(c.Name, "game_session") {
			idx = i
		}
	}

	return makeGameSessionFromContextParameters(contexts[idx].Parameters)

}

func makeGameSessionFromContextParameters(parameters []byte) (GameSession, error) {
	gs := GameSession{}

	err := json.Unmarshal(parameters, &gs)

	return gs, err
}

func floatArray2String(arr []float64) string {

	var sa []string

	for _, f := range arr {
		sa = append(sa, strconv.FormatFloat(f, 'f', 2, 64))
	}

	return "[" + strings.Join(sa, ",") + "]"
}

func (rcv GameSession) GetPointsOfCurrentTurn() float64 {
	switch int(rcv.Turn) {
	case 1:
		return 10.00
	case 2:
		return 30.00
	case 3:
		return 50.00
	case 4:
		return 80.00
	default:
		return 100.00
	}
}

func (rcv GameSession) GetCurrentTurnFirstPlayerName() string {
	remainder := math.Mod(rcv.Turn, 2)
	if remainder == 0 {
		return rcv.NamePlayerTwo
	}
	return rcv.NamePlayerOne

}
