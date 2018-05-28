package intentHandlers

import (
	"github.com/BloodyRainer/articlePrice/df"
	"context"
	"github.com/BloodyRainer/articlePrice/search"
	engLog "google.golang.org/appengine/log"
	"strconv"
)

func AskArticleQuestionFirstPlayer(ctx context.Context, dfReq df.Request) (*df.Response, error) {

	gs, err := df.MakeGameSessionFromDfRequest(dfReq)
	if err != nil {
		return nil, err
	}

	a, err := search.GetRandomArticle(ctx)
	if err != nil {
		return nil, err
	}

	engLog.Infof(ctx, "asking random article: "+a.String())

	gs.IncrementTrun()
	engLog.Infof(ctx, "current turn: %v ", gs.Turn)

	gs.CurrentArticleName = a.Name
	gs.CurrentArticleNumber = a.ArticleNr
	gs.CurrentArticlePrice, err = strconv.ParseFloat(a.Price, 64)
	gs.CurrentArticleImgUrl = a.ImgUrl
	if err != nil {
		return nil, err
	}

	resp := makeArticleQuestionFirstPlayer(dfReq, gs)

	return resp, nil
}

func makeArticleQuestionFirstPlayer(dfReq df.Request, gs *df.GameSession) *df.Response {

	points := strconv.Itoa(int(gs.GetPointsOfCurrentTurn()))
	pName := gs.GetFirstPlayerName()

	payload := df.MakeSimpleRespPayload(true,
		"<speak>Wie ist der Preis von "+df.ModifyForTTS(gs.CurrentArticleName)+" auf otto D E? <break time='2000ms'/> " + pName+", du bist dran, für "+ points +" Punkte, wie ist dein Tipp?</speak>",
		"Wie ist der Preis von "+gs.CurrentArticleName+" auf otto.de? " + pName+", du bist dran, für "+ points +" Punkte, wie ist dein Tipp?")

	bc := df.Item{
		BasicCard: &df.BasicCard{
			Title: gs.CurrentArticleName,
			Image: &df.Image{
				Url:               gs.CurrentArticleImgUrl,
				AccessibilityText: "zu diesem Artikel konnte keine Vorschau gefunden werden",
			},
		},
	}

	payload.Google.RichResponse.Items = append(payload.Google.RichResponse.Items, bc)

	resp := &df.Response{
		Source:  source,
		Payload: payload,
		OutputContexts: []df.Context{
			df.MakeOutputContext("first_player_answer", 3, nil, dfReq),
			df.MakeOutputContext("game_session", 5, gs.ToParameters(), dfReq),
			df.MakeOutputContext("ask_first_player", 0, nil, dfReq),
		},
	}

	return resp
}
