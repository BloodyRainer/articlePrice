package search

import (
	"github.com/BloodyRainer/articlePrice/model"
	"context"
	engLog "google.golang.org/appengine/log"
	"sync"
	"github.com/BloodyRainer/articlePrice/df"
)

const domain = "www.otto.de"
const path = "/p/search/"
const protocol = "https"
const queryPrefix = "articlenumber="

// Appengine needs the original request.
func GetRandomArticle(ctx context.Context) (*df.Article, error) {
	aNr := model.RandomArticleNr()

	a, err := requestArticleByArticleNr(ctx, aNr)
	if err != nil {
		engLog.Errorf(ctx, err.Error())
		engLog.Errorf(ctx, "error finding article with number "+aNr)
		return nil, err
	}

	return a, nil
}

func requestArticleByArticleNr(ctx context.Context, articleNr string) (*df.Article, error) {
	respBody, url, err := searchArticleByNr(ctx, articleNr)
	if err != nil {
		return nil, err
	}

	a, err := createArticleFromHtml(ctx, respBody)
	if err != nil {
		return nil, err
	}

	a.ArticleNr = articleNr
	a.Link = url

	return a, nil
}

func createArticleFromHtml(ctx context.Context, html string) (*df.Article, error) {

	var err error
	var name string
	var price string
	var imgUrl string

	wg := &sync.WaitGroup{}

	wg.Add(3)

	go func() {
		name, err = getName(html)
		if err != nil {
			engLog.Errorf(ctx, "failed to parse name: "+err.Error())
		}
		wg.Done()
	}()

	go func() {
		price, err = getPrice(html)
		if err != nil {
			engLog.Errorf(ctx, "failed to parse price: "+err.Error())
		}
		wg.Done()
	}()

	go func() {
		imgUrl, err = getImageUrl(html)
		if err != nil {
			engLog.Errorf(ctx, "failed to parse url: "+err.Error())
		}
		wg.Done()
	}()

	wg.Wait()

	if err != nil {
		return nil, err
	}

	a := &df.Article{
		Name:   name,
		Price:  price,
		ImgUrl: imgUrl,
	}

	return a, nil
}
