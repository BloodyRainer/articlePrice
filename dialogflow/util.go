package dialogflow

import (
	"github.com/BloodyRainer/articlePrice/model"
	"strconv"
)

func MakeArticleNameResponse(a model.Article, dfReq DfRequest) *DfResponse {
	return &DfResponse{
		Source:  "Der Preis ist heiss",
		Payload: MakePayload(true, "Wie ist der Preis von: "+a.Name+"?"),
		OutputContexts: []Context{
			{
				Name:          dfReq.Session + "/contexts/question_was_asked",
				LifespanCount: 1,
				Parameters:    []byte(`{"articleNumber":"` + a.ArticleNr + `", "articleName":"` + a.Name + `", "actualPrice":` + a.Price + `}`),
			},
		},
	}
}

func MakeEvaluatedResponse(g Guess) *DfResponse {
	return &DfResponse{
		Source:  "Der Preis ist heiss",
		Payload: MakePayload(false, assembleTextResponse(g)),
	}
}

func MakePayload(expectUserResponse bool, textToSpeech string) *Payload {
	return &Payload{
		Google: &Google{
			ExpectUserResponse: expectUserResponse,
			RichResponse: &RichResponse{
				Items: []Item{
					{
						SimpleResponse: &SimpleResponse{
							TextToSpeech: textToSpeech,
						},
					},
				},
			},
		},
	}
}

func assembleTextResponse(g Guess) string {
	diff := differenceGuessActualInPercent(g.PriceGuess, g.ActualPrice)
	ap := strconv.FormatFloat(g.ActualPrice, 'f', 2, 64)
	gp := strconv.FormatFloat(g.PriceGuess, 'f', 2, 64)

	if diff > 5.00 {
		return "Zu hoch, der Artikel kostet in Wirklichkeit: " + ap + " Euro."
	} else if diff < -50.00 {
		return "Ja ganz genau, du kannst den Artikel fuer " + gp +
			" Euro sofort in unserem Paketshop direkt auf dem Zucker Berg im Schokoladenviertel in der Wuensch Dir Was Allee abholen! " +
			"Nein, echter Preis: " + ap + " Euro."
	} else if diff < -5.00 {
		return "Zu tief, der Artikel kostet in Wirklichkeit: " + ap + " Euro."
	} else {
		return "Gut geraten, der Artikel kostet in Wirklichkeit: " + ap + " Euro."
	}

}

func differenceGuessActualInPercent(guess float64, actual float64) float64 {
	return (guess - actual) / actual * 100.00
}
