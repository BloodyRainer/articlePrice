package intentHandlers

import (
	"github.com/BloodyRainer/articlePrice/df"
	"context"
	"github.com/BloodyRainer/articlePrice/search"
	engLog "google.golang.org/appengine/log"
)

const source = "Der Preis ist heiss"

func AskRandomArticle(ctx context.Context, dfReq *df.Request) (*df.Response, error) {
	a, err := search.GetRandomArticle(ctx)

	engLog.Infof(ctx, "random article: "+a.String())

	if err != nil {
		return nil, err
	}

	resp := makeArticleNameResponse(ctx, *a, *dfReq)

	return resp, nil
}

func makeArticleNameResponse(ctx context.Context, a df.Article, dfReq df.Request) *df.Response {

	payload := df.MakeSimpleRespPayload(true,
		"<speak>Wie ist der Preis von "+df.ModifyForTTS(a.Name)+" auf otto D E?</speak>",
		"Wie ist der Preis von "+a.Name+" auf otto.de?")

	bc := df.Item{
		BasicCard: &df.BasicCard{
			Title: a.Name,
			Image: &df.Image{
				Url:               a.ImgUrl,
				AccessibilityText: "zu diesem Artikel konnte keine Vorschau gefunden werden",
			},
		},
	}

	payload.Google.RichResponse.Items = append(payload.Google.RichResponse.Items, bc)

	resp := &df.Response{
		Source:  source,
		Payload: payload,
		OutputContexts: []df.Context{
			df.MakeOutputContext("question_was_asked", 3, a.ToParameters(), dfReq),
		},
	}

	engLog.Infof(ctx, resp.Payload.Google.RichResponse.Items[0].SimpleResponse.DisplayText)

	return resp
}
