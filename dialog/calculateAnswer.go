package dialog

import (
	"github.com/BloodyRainer/articlePrice/model"
	"context"
	englog "google.golang.org/appengine/log"
)

const source = "Der Preis ist heiss"

func MakeArticleNameResponse(ctx context.Context, a model.Article, dfReq DfRequest) *DfResponse {
	params := []byte(`{"articleNumber":"` + a.ArticleNr + `", "articleName":"` + a.Name + `", "actualPrice":` + a.Price + `, "imgUrl":"` + a.ImgUrl + `", "link":"` + a.Link + `"}`)
	payload := MakeSimpleRespPayload(true,
		"<speak>Wie ist der Preis von "+ModifyForTTS(a.Name)+" auf otto D E?</speak>",
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
	tr := calculateDiffPercentTreshold(g)

	if diffPercent > tr {
		return "<speak> Zu hoch! <break time='500ms'/> Der Artikel kostet in Wirklichkeit nur " + PriceInEuroTTS(g.ActualPrice) + ". </speak>",
			"Zu hoch, der Artikel kostet in Wirklichkeit nur " + PriceInEuroText(g.ActualPrice) + ". "
	} else if diffPercent < -tr*10 {
		return "<speak> Ja ganz genau! <break time='500ms'/> Du kannst den Artikel fuer " + PriceInEuroTTS(g.PriceGuess) +
			" sofort in dem Paketshop direkt im Schokoladenviertel auf dem Zucker Berg in der Wuensch Dir Was Allee abholen! " +
			"<audio src='https://actions.google.com/sounds/v1/cartoon/drum_roll.ogg' clipEnd='4s'></audio> <break time='700ms'/> Nein, " +
			"der echte Preis ist natuerlich " + PriceInEuroTTS(g.ActualPrice) + ". </speak>",
			"Ja ganz genau, du kannst den Artikel fuer " + PriceInEuroText(g.PriceGuess) +
				" sofort in dem Paketshop direkt im Schokoladenviertel auf dem Zuckerberg in der Wuensch-Dir-Was-Allee abholen!" +
				" Nein, der echte Preis ist natuerlich " + PriceInEuroText(g.ActualPrice) + "."
	} else if diffPercent < -tr {
		return "<speak> Zu tief! <break time='500ms'/> Der Artikel kostet in Wirklichkeit " + PriceInEuroTTS(g.ActualPrice) + ".</speak>",
			"Zu tief, der Artikel kostet in Wirklichkeit " + PriceInEuroText(g.ActualPrice) + "."
		//} else if diffPercent < tr / tr {
		//	return "<speak> <audio src='https://firebasestorage.googleapis.com/v0/b/whatisit-72c26.appspot.com/o/success.mp3?alt=media'></audio> Das wusstest du wohl! <break time='500ms'/> Der Artikel kostet genau " + ap + " Euro. </speak>",
		//		"Das wusstest du wohl! Der Artikel kostet genau " + ap + "."
	} else {
		return "<speak> <audio src='https://firebasestorage.googleapis.com/v0/b/whatisit-72c26.appspot.com/o/success.mp3?alt=media'></audio> Gut geraten! <break time='500ms'/> Der Artikel kostet tatsaechlich " + PriceInEuroTTS(g.ActualPrice) + ".</speak>",
			"Gut geraten, der Artikel kostet tatsaechlich " + PriceInEuroText(g.ActualPrice) + "."
	}

}

func calculateDiffPercentTreshold(g Guess) float64 {
	if g.ActualPrice < 10.00 {
		return 20.00
	} else if g.ActualPrice < 30.00 {
		return 15.00
	} else {
		return 7.00
	}
}

func differenceGuessActualInPercent(guess float64, actual float64) float64 {
	return (guess - actual) / actual * 100.00
}
