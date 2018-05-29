package df

import (
	"strings"
	"encoding/json"
)

type Guess struct {
	ArticleNr   string       `json:"articleNumber"`
	ArticleName string       `json:"articleName"`
	ActualPrice float64      `json:"actualPrice"`
	Price       UnitCurrency `json:"unit-currency"`
	Link        string       `json:"link"`
}

type UnitCurrency struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

func MakeGuessFromDfRequest(dfReq Request) (Guess, error) {

	contexts := dfReq.QueryResult.OutputContexts

	var indexOfQwa int

	for i, c := range contexts {
		if strings.Contains(c.Name, "question_was_asked") {
			indexOfQwa = i
		}
	}

	return makeGuessFromContextParameters(contexts[indexOfQwa].Parameters)

}

func makeGuessFromContextParameters(parameters []byte) (Guess, error) {
	g := Guess{}

	err := json.Unmarshal(parameters, &g)

	return g, err
}
