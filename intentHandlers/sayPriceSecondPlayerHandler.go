package intentHandlers

import (
	"github.com/BloodyRainer/articlePrice/df"
	"math"
	"strconv"
	"bytes"
	"context"
)

func SavePriceSecondPlayerAndResultsOfTurn(ctx context.Context, dfReq df.Request) (*df.Response, error) {

	gs, err := df.MakeGameSessionFromDfRequest(dfReq)
	if err != nil {
		return nil, err
	}

	priceSP, err := df.GetPriceAnswer(dfReq, "second_player_answer")
	if err != nil {
		return nil, err
	}

	gs.SaveSecondPriceAnswer(priceSP)

	resp := makeResultsCurrentTurn(ctx, dfReq, gs)

	return resp, nil

}

func makeResultsCurrentTurn(ctx context.Context, dfReq df.Request, gs *df.GameSession) *df.Response {

	lifeSpanGs := 5
	lifeSpanAfp := 3
	turn := gs.Turn

	tts, text := generateTTSAndTextAnswer(gs)

	if turn == endingTurn {
		appendFinalResult(&tts, &text, gs)
		lifeSpanGs = 0
		lifeSpanAfp = 0
	} else if turn == endingTurn -2 {
		tts.WriteString(" <break time='2000ms'/> Es folgt die vorletzte Runde! <break time='1000ms'/>Seid ihr bereit?</speak>")
		text.WriteString("\n\nVorletzte Runde! Bereit?")
	} else if turn == endingTurn -1 {
		tts.WriteString(" <break time='2000ms'/> Es folgt die letzte Runde! <break time='1000ms'/>Seid ihr bereit?</speak>")
		text.WriteString("\n\nLetzte Runde! Bereit?")
	} else {
		tts.WriteString(" <break time='2000ms'/> Bereit für die nächste Runde?</speak>")
		text.WriteString("\n\nNächste Runde?")
	}

	// reset article for next turn
	gs.CurrentArticlePrice = -1
	gs.CurrentArticleName = ""
	gs.CurrentArticleImgUrl = ""
	gs.CurrentArticleNumber = ""

	payload := df.MakeSimpleRespPayload(true, tts.String(), text.String())

	appendSuggestion(payload, turn)

	resp := &df.Response{
		Source:  source,
		Payload: payload,
		OutputContexts: []df.Context{
			df.MakeOutputContext("ask_first_player", lifeSpanAfp, nil, dfReq),
			df.MakeOutputContext("game_session", lifeSpanGs, gs.ToParameters(), dfReq),
			df.MakeOutputContext("second_player_answer", 0, nil, dfReq),
		},
	}

	return resp

}

func appendSuggestion(p *df.Payload, turn float64) {

	var suggestions []df.Suggestion

	if turn > endingTurn-1 {
		suggestions = []df.Suggestion{
			{
				Title: "noch mal spielen",
			},
		}
	} else {
		suggestions = []df.Suggestion{
			{
				Title: "ja",
			},
		}
	}

	p.Google.RichResponse.Suggestions = suggestions
}

func appendFinalResult(tts, text *bytes.Buffer, gs *df.GameSession) {
	w, l, wp, lp := gs.GetFinalResult()

	if wp != lp {
		tts.WriteString("<break time='2000ms'/> Die letzte Runde ist abgeschlossen. <break time='1000ms'/>")
		tts.WriteString(l)
		tts.WriteString(" hat insgesamt <say-as interpret-as='cardinal'>")
		tts.WriteString(lp)
		tts.WriteString("</say-as> Punkte gesammelt. <break time='1000ms'/>")
		tts.WriteString(w)
		tts.WriteString(" kommt hingegen auf <say-as interpret-as='cardinal'>")
		tts.WriteString(wp)
		tts.WriteString("</say-as>  Punkte und hat dieses mal gewonnen! <break time='500ms'/> Gut gespielt! </speak>")

		text.WriteString("\n\n")
		text.WriteString("Endergebnis: \n")
		text.WriteString(l)
		text.WriteString(": \t")
		text.WriteString(lp)
		text.WriteString(" Punkte. \n")
		text.WriteString(w)
		text.WriteString(": \t")
		text.WriteString(wp)
		text.WriteString(" Punkte.")
		text.WriteString("\n\nGewinner: ")
		text.WriteString(w)
		text.WriteString("!")
	} else {
		tts.WriteString("<break time='2000ms'/> Die letzte Runde ist abgeschlossen. ")
		tts.WriteString("Beide Spieler kommen insgesamt auf ")
		tts.WriteString(wp)
		tts.WriteString(" Punkte. Damit endet das Spiel mit einem Unentschieden! Gut gespielt! </speak>")

		text.WriteString("\n\nDie letzte Runde ist abgeschlossen. ")
		text.WriteString("Beide Spieler kommen insgesamt auf ")
		text.WriteString(wp)
		text.WriteString(" Punkte. \nDamit endet das Spiel mit einem Unentschieden! \n\nGut gespielt!")
	}
}

func generateTTSAndTextAnswer(gs *df.GameSession) (bytes.Buffer, bytes.Buffer) {
	var tts, text bytes.Buffer

	w, l, wp, lp := calcAndSaveResults(gs)

	if wp != lp {
		tts.WriteString("<speak>")
		tts.WriteString("<audio src='https://firebasestorage.googleapis.com/v0/b/whatisit-72c26.appspot.com/o/success.mp3?alt=media'></audio>")
		tts.WriteString(w)
		tts.WriteString(" lag näher am tatsächlichen Preis von ")
		tts.WriteString(df.PriceInEuroTTS(gs.CurrentArticlePrice))
		tts.WriteString(" und bekommt <say-as interpret-as='cardinal'>")
		tts.WriteString(wp)
		tts.WriteString("</say-as> Punkte! <break time='500ms'/>")
		tts.WriteString(l)
		tts.WriteString(" erhält <say-as interpret-as='cardinal'>")
		tts.WriteString(lp)
		tts.WriteString("</say-as> Punkte!")

		text.WriteString("Tatsächlicher Preis: ")
		text.WriteString(df.PriceInEuroText(gs.CurrentArticlePrice))
		text.WriteString(".\n\n")
		text.WriteString(w)
		text.WriteString(" +")
		text.WriteString(wp)
		text.WriteString(" Punkte! \n")
		text.WriteString(l)
		text.WriteString(" +")
		text.WriteString(lp)
		text.WriteString(" Punkte!")

	} else {
		tts.WriteString("<speak>")
		tts.WriteString(" Beide Spieler lagen gleich nah am tatsächlichen Preis von ")
		tts.WriteString(df.PriceInEuroTTS(gs.CurrentArticlePrice))
		tts.WriteString("! Beide Spieler erhalten <say-as interpret-as='cardinal'>")
		tts.WriteString(wp)
		tts.WriteString("</say-as> Punkte! ")

		text.WriteString("Beide Spieler lagen gleich nah am tatsächlichen Preis von ")
		text.WriteString(df.PriceInEuroText(gs.CurrentArticlePrice))
		text.WriteString("! Beide Spieler erhalten ")
		text.WriteString(wp)
		text.WriteString("Punkte! ")
	}

	return tts, text
}

func calcAndSaveResults(gs *df.GameSession) (winner string, loser string, pWinner string, pLoser string) {

	pointsWinner := gs.GetPointsOfCurrentTurn()
	pointsLoser := pointsWinner / 2

	diffPlayerOne := math.Abs(gs.CurrentArticlePrice - gs.GetPriceGuessOfCurrentTurnPlayerOne())
	diffPlayerTwo := math.Abs(gs.CurrentArticlePrice - gs.GetPriceGuessOfCurrentTurnPlayerTwo())

	if diffPlayerOne > diffPlayerTwo {
		gs.PointsPlayerTwo = gs.PointsPlayerTwo + pointsWinner
		gs.PointsPlayerOne = gs.PointsPlayerOne + pointsLoser
		return gs.NamePlayerTwo, gs.NamePlayerOne, strconv.Itoa(int(pointsWinner)), strconv.Itoa(int(pointsLoser))
	} else if diffPlayerOne < diffPlayerTwo {
		gs.PointsPlayerOne = gs.PointsPlayerOne + pointsWinner
		gs.PointsPlayerTwo = gs.PointsPlayerTwo + pointsLoser
		return gs.NamePlayerOne, gs.NamePlayerTwo, strconv.Itoa(int(pointsWinner)), strconv.Itoa(int(pointsLoser))
	} else {
		gs.PointsPlayerOne = gs.PointsPlayerOne + pointsWinner
		gs.PointsPlayerTwo = gs.PointsPlayerTwo + pointsWinner
		return "", "", strconv.Itoa(int(pointsWinner)), strconv.Itoa(int(pointsWinner))
	}

}
