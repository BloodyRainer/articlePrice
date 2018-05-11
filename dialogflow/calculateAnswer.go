package dialogflow

import (
	"github.com/BloodyRainer/articlePrice/model"
	"strconv"
	"context"
	englog "google.golang.org/appengine/log"
)

const source = "Der Preis ist heiss"

func MakeArticleNameResponse(ctx context.Context, a model.Article, dfReq DfRequest) *DfResponse {
	params := []byte(`{"articleNumber":"` + a.ArticleNr + `", "articleName":"` + a.Name + `", "actualPrice":` + a.Price + `, "imgUrl":"` + a.ImgUrl + `", "link":"` + a.Link + `"}`)
	payload := MakeSimpleRespPayload(true,
		"Wie ist der Preis von "+a.Name+" auf otto D E?",
		"Wie ist der Preis von "+a.Name+" auf otto.de?")

	bc := Item{
		BasicCard: &BasicCard{
			Title: a.Name,
			Image: &Image{
				Url:               a.ImgUrl,
				AccessibilityText: "zu diesem Artikel konnte keine Vorschau gefunden werden",
			},
		},
	}

	payload.Google.RichResponse.Items = append(payload.Google.RichResponse.Items, bc)

	resp := &DfResponse{
		Source:  source,
		Payload: payload,
		OutputContexts: []Context{
			{
				Name:          dfReq.Session + "/contexts/question_was_asked",
				LifespanCount: 3,
				Parameters:    params,
			},
		},
	}

	englog.Infof(ctx, resp.Payload.Google.RichResponse.Items[0].SimpleResponse.DisplayText)

	return resp
}

// the response asks for new input
func MakeNewInputResponse() *DfResponse {
	text := "Das habe ich nicht verstanden. Sage Preise am besten ohne Cent-Betraege, also zum Beispiel 59 oder 59 Euro."
	payload := MakeSimpleRespPayload(false, text, text)

	return &DfResponse{
		Source:  source,
		Payload: payload,
	}
}

func MakeEvaluatedResponse(g Guess) *DfResponse {
	tts, dt := calculateResponse(g)

	payload := MakeSimpleRespPayload(false, tts, dt)

	suggestions := []Suggestion{
		{
			Title: "noch mal",
		},
	}

	payload.Google.RichResponse.Suggestions = suggestions

	//TODO: must be verified by url-owner...
	//los := &LinkOutSuggestion{
	//	DestinationName: "auf otto.de",
	//	Url: g.Link,
	//}
	//
	//payload.Google.RichResponse.LinkOutSuggestion = los

	return &DfResponse{
		Source:  source,
		Payload: payload,
	}
}

func MakeSimpleRespPayload(expectUserResponse bool, textToSpeech string, displayText string) *Payload {
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

// first response is text-to-speech, second is display-text
func calculateResponse(g Guess) (string, string) {

	diffPercent := differenceGuessActualInPercent(g.PriceGuess, g.ActualPrice)
	ap := strconv.FormatFloat(g.ActualPrice, 'f', 2, 64)
	gp := strconv.FormatFloat(g.PriceGuess, 'f', 2, 64)

	t := calculateDiffPercentTreshold(g)

	if diffPercent > t {
		return "<speak> Zu hoch! <break time='500ms'/> Der Artikel kostet in Wirklichkeit nur " + ap + " Euro. </speak>",
			"Zu hoch, der Artikel kostet in Wirklichkeit nur " + ap + " Euro. "
	} else if diffPercent < -t * 10 {
		return "<speak> Ja ganz genau! <break time='500ms'/> Du kannst den Artikel fuer " + gp +
			" Euro sofort in dem Paketshop direkt auf dem Zucker Berg im Schokoladenviertel in der Wuensch Dir Was Allee abholen! " +
			"<break time='1500ms'/> Nein, der echte Preis ist " + ap + " Euro. </speak>",
			"Ja ganz genau, du kannst den Artikel fuer " + gp +
				" Euro sofort in dem Paketshop direkt auf dem Zuckerberg im Schokoladenviertel in der Wuensch-Dir-Was-Allee abholen!" +
				" Nein, der echte Preis ist " + ap + " Euro. "
	} else if diffPercent < -t {
		return "<speak> Zu tief! <break time='500ms'/> Der Artikel kostet in Wirklichkeit " + ap + " Euro. </speak>",
			"Zu tief, der Artikel kostet in Wirklichkeit " + ap + " Euro. "
	} else {
		return "<speak> Gut geraten! <break time='500ms'/> Der Artikel kostet tatsaechlich " + ap + " Euro. </speak>",
			"Gut geraten, der Artikel kostet tatsaechlich " + ap + " Euro."
	}

}

func calculateDiffPercentTreshold(g Guess) float64 {
	if g.ActualPrice < 10.00 {
		return 20.00
	} else if g.ActualPrice < 30.00{
		return 15.00
	} else {
		return 7.00
	}
}

func differenceGuessActualInPercent(guess float64, actual float64) float64 {
	return (guess - actual) / actual * 100.00
}
