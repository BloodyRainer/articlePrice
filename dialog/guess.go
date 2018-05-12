package dialog

import (
	"strings"
	"encoding/json"
)

type Guess struct {
	ArticleNr   string  `json:"articleNumber"`
	ArticleName string  `json:"articleName"`
	ActualPrice float64 `json:"actualPrice"`
	PriceGuess  float64 `json:"number"`
	OrgNumber   string  `json:"number.original"`
	Link        string  `json:"link"`
}

func MakeGuessFromDfRequest(dfReq DfRequest) (Guess, error) {

	contexts := dfReq.QueryResult.OutputContexts

	var indexOfQwa int

	for i, c := range contexts {
		if strings.Contains(c.Name, "question_was_asked") {
			indexOfQwa = i
		}
	}

	return makeGuessFromContextParameters(contexts[indexOfQwa].Parameters)

}

func makeGuessFromContextParameters(parameters json.RawMessage) (Guess, error) {
	g := Guess{}

	err := json.Unmarshal(parameters, &g)

	return g, err
}

//func (rcv Guess) getActualPriceAsFloat64() (float64, error) {
//	ap, err := strconv.ParseFloat(rcv.ActualPrice, 64)
//
//}
