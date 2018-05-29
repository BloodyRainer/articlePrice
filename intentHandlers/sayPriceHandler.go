package intentHandlers

import (
	"github.com/BloodyRainer/articlePrice/df"
	"context"
)

// in case guessed price could not be parsed
func AskForNewInput() *df.Response {
	return makeNewInputResponse()
}

func RespondToPriceGuess(ctx context.Context, dfReq df.Request) (*df.Response, error) {

	g, err := df.MakeGuessFromDfRequest(dfReq)
	if err != nil {
		return nil, err
	}

	resp := makeEvaluatedResponse(g)

	return resp, nil
}

// the response asks for new input
func makeNewInputResponse() *df.Response {
	text := "Das habe ich nicht verstanden. Sage Preise am besten ohne Cent-Beträge, also zum Beispiel 59 Euro."
	payload := df.MakeSimpleRespPayload(false, text, text)

	return &df.Response{
		Source:  source,
		Payload: payload,
	}
}

func makeEvaluatedResponse(g df.Guess) *df.Response {
	tts, dt := calculateResponse(g)

	payload := df.MakeSimpleRespPayload(false, tts, dt)

	suggestions := []df.Suggestion{
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

	return &df.Response{
		Source:  source,
		Payload: payload,
	}
}

// first response is text-to-speech, second is display-text
func calculateResponse(g df.Guess) (string, string) {

	diffPercent := differenceGuessActualInPercent(g.Price.Amount, g.ActualPrice)
	tr := calculateDiffPercentTreshold(g)

	if diffPercent > 2.5*tr {
		return "<speak> Zu hoch! <break time='500ms'/> Der Artikel kostet in Wirklichkeit nur " + df.PriceInEuroTTS(g.ActualPrice) + ". </speak>",
			"Zu hoch, der Artikel kostet in Wirklichkeit nur " + df.PriceInEuroText(g.ActualPrice) + ". "
	} else if diffPercent > tr {
		return "<speak> <audio src='https://firebasestorage.googleapis.com/v0/b/whatisit-72c26.appspot.com/o/success.mp3?alt=media'></audio> Gar nicht schlecht! <break time='500ms'/> Der Artikel kostet in Wirklichkeit aber nur " + df.PriceInEuroTTS(g.ActualPrice) + ". </speak>",
			"Gar nicht schlecht, der Artikel kostet in Wirklichkeit aber nur " + df.PriceInEuroText(g.ActualPrice) + ". "
	} else if diffPercent < -tr*9 {
		return "<speak> Ja ganz genau! <break time='500ms'/> Du kannst den Artikel fuer " + df.PriceInEuroTTS(g.Price.Amount) +
			" sofort in dem Paketshop direkt im Schokoladenviertel auf dem Zucker Berg in der Wuensch Dir Was Allee abholen! " +
			"<audio src='https://actions.google.com/sounds/v1/cartoon/drum_roll.ogg' clipEnd='4s'></audio> <break time='700ms'/> Nein, " +
			"der echte Preis ist natuerlich " + df.PriceInEuroTTS(g.ActualPrice) + ". </speak>",
			"Ja ganz genau, du kannst den Artikel fuer " + df.PriceInEuroText(g.Price.Amount) +
				" sofort in dem Paketshop direkt im Schokoladenviertel auf dem Zuckerberg in der Wuensch-Dir-Was-Allee abholen!" +
				" Nein, der echte Preis ist natuerlich " + df.PriceInEuroText(g.ActualPrice) + "."
	} else if diffPercent < - 2.5*tr {
		return "<speak> Zu tief! <break time='500ms'/> Der Artikel kostet in Wirklichkeit " + df.PriceInEuroTTS(g.ActualPrice) + ".</speak>",
			"Zu tief, der Artikel kostet in Wirklichkeit " + df.PriceInEuroText(g.ActualPrice) + "."
	} else if diffPercent < -tr {
		return "<speak> <audio src='https://firebasestorage.googleapis.com/v0/b/whatisit-72c26.appspot.com/o/success.mp3?alt=media'></audio> Gar nicht schlecht! <break time='500ms'/> Der Artikel kostet in Wirklichkeit allerdings " + df.PriceInEuroTTS(g.ActualPrice) + ".</speak>",
			"Gar nicht schlecht, der Artikel kostet in Wirklichkeit allerdings " + df.PriceInEuroText(g.ActualPrice) + "."
		//} else if diffPercent < tr / tr {
		//	return "<speak> <audio src='https://firebasestorage.googleapis.com/v0/b/whatisit-72c26.appspot.com/o/success.mp3?alt=media'></audio> Das wusstest du wohl! <break time='500ms'/> Der Artikel kostet genau " + ap + " Euro. </speak>",
		//		"Das wusstest du wohl! Der Artikel kostet genau " + ap + "."
	} else {
		return "<speak> <audio src='https://firebasestorage.googleapis.com/v0/b/whatisit-72c26.appspot.com/o/success.mp3?alt=media'></audio> Gut geraten! <break time='500ms'/> Der Artikel kostet tatsächlich " + df.PriceInEuroTTS(g.ActualPrice) + ".</speak>",
			"Gut geraten, der Artikel kostet tatsächlich " + df.PriceInEuroText(g.ActualPrice) + "."
	}

}

func calculateDiffPercentTreshold(g df.Guess) float64 {
	if g.ActualPrice < 10.00 {
		return 20.00
	} else if g.ActualPrice < 30.00 {
		return 15.00
	} else {
		return 10.00
	}
}

func differenceGuessActualInPercent(guess float64, actual float64) float64 {
	return (guess - actual) / actual * 100.00
}
