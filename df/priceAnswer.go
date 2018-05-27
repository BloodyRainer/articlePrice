package df

import (
	"strings"
	"encoding/json"
)

type PriceAnswer struct {
	PriceGuess float64 `json:"price"`
	OrgNumber  string  `json:"price.original"`
}

func GetPriceAnswer(dfReq Request, contextName string) (float64, error) {
	contexts := dfReq.QueryResult.OutputContexts

	var idx int

	for i, c := range contexts {
		if strings.Contains(c.Name, contextName) {
			idx = i
		}
	}

	return makePriceAnwerFromParameters(contexts[idx].Parameters)

}

func makePriceAnwerFromParameters(parameters []byte) (float64, error) {
	pa := PriceAnswer{}

	err := json.Unmarshal(parameters, &pa)

	return pa.PriceGuess, err
}
