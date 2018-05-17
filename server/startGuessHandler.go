package server

import (
	"github.com/BloodyRainer/articlePrice/dialog"
	"context"
	"github.com/BloodyRainer/articlePrice/search"
	engLog "google.golang.org/appengine/log"
)

func askRandomArticle(ctx context.Context, dfReq *dialog.DfRequest) (*dialog.DfResponse, error) {
	a, err := search.GetRandomArticle(ctx)

	engLog.Infof(ctx, "random article: "+a.String())

	if err != nil {
		return nil, err
	}

	resp := dialog.MakeArticleNameResponse(ctx, *a, *dfReq)

	return resp, nil
}
