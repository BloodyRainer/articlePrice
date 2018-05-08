package dialogflow

import (
	"github.com/BloodyRainer/articlePrice/model"
	"strconv"
)

func MakeArticleNameResponse(a model.Article, dfReq DfRequest) *DfResponse {
	params := []byte(`{"articleNumber":"` + a.ArticleNr + `", "articleName":"` + a.Name + `", "actualPrice":` + a.Price + `, "imgUrl": "` + a.ImgUrl + `"}`)
	return &DfResponse{
		Source: "Der Preis ist heiss",
		Payload: MakePayload(true,
			"Wie ist der Preis von "+a.Name+" auf otto D E?",
			"Wie ist der Preis von "+a.Name+" auf otto.de?"),
		OutputContexts: []Context{
			{
				Name:          dfReq.Session + "/contexts/question_was_asked",
				LifespanCount: 1,
				Parameters:    params,
			},
		},
	}
}

func MakeEvaluatedResponse(g Guess) *DfResponse {
	cr := calculateResponse(g)
	return &DfResponse{
		Source:  "Der Preis ist heiss",
		Payload: MakePayload(false, cr, cr),
	}
}

func MakePayload(expectUserResponse bool, textToSpeech string, displayText string) *Payload {
	return &Payload{
		Google: &Google{
			ExpectUserResponse: expectUserResponse,
			RichResponse: &RichResponse{
				Items: []Item{
					{
						SimpleResponse: &SimpleResponse{
							TextToSpeech: textToSpeech,
							DisplayText:  displayText,
						},
					},
				},
			},
		},
	}
}

func calculateResponse(g Guess) string {
	diff := differenceGuessActualInPercent(g.PriceGuess, g.ActualPrice)
	ap := strconv.FormatFloat(g.ActualPrice, 'f', 2, 64)
	gp := strconv.FormatFloat(g.PriceGuess, 'f', 2, 64)

	if diff > 5.00 {
		return "Zu hoch, der Artikel kostet in Wirklichkeit nur " + ap + " Euro."
	} else if diff < -50.00 {
		return "Ja ganz genau, du kannst den Artikel fuer " + gp +
			" Euro sofort in dem Paketshop direkt auf dem Zucker Berg im Schokoladenviertel in der Wuensch Dir Was Allee abholen! " +
			"Nein, der echte Preis ist " + ap + " Euro."
	} else if diff < -5.00 {
		return "Zu tief, der Artikel kostet in Wirklichkeit " + ap + " Euro."
	} else {
		return "Gut geraten, der Artikel kostet tatsaechlich " + ap + " Euro."
	}

}

func differenceGuessActualInPercent(guess float64, actual float64) float64 {
	return (guess - actual) / actual * 100.00
}
