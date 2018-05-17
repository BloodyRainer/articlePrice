package server

import "github.com/BloodyRainer/articlePrice/dialog"

// in case guessed price could not be parsed
func askForNewInput() *dialog.DfResponse {
	return dialog.MakeNewInputResponse()
}

func respondToPriceGuess(dfReq dialog.DfRequest) (*dialog.DfResponse, error) {

	g, err := dialog.MakeGuessFromDfRequest(dfReq)
	if err != nil {
		return nil, err
	}

	resp := dialog.MakeEvaluatedResponse(g)

	return resp, nil
}
