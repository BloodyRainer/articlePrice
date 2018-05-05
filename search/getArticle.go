package search

import (
	"log"
	"regexp"
	"github.com/BloodyRainer/articlePrice/model"
	"context"
)

var priceReg *regexp.Regexp
var nameReg *regexp.Regexp

const domain = "www.otto.de"
const path = "/p/search/"
const protocol = "https"
const queryPrefix = "articlenumber="

func init() {
	priceReg = regexp.MustCompile(`itemprop="price".*content="(\d+\.?\d+)"`)
	nameReg = regexp.MustCompile(`<h1.*itemprop="name".*>(.*)</h1>`)
}

// Appengine needs the original request.
func GetRandomArticle(ctx context.Context) (*model.Article, error) {
	aNr := model.RandomArticleNr()

	a, err := requestNameAndPriceByArctileNr(ctx, aNr)
	if err != nil {
		log.Print("no article with number: ", aNr)
		return nil, err
	}

	return a, nil
}

//func GetPriceByArticleNr(ctx context.Context, aNr string) (float64, error) {
//	respBody, err := searchArticle(ctx, aNr)
//	if err != nil {
//		return -1, err
//	}
//
//	price, err := getPrice(respBody)
//	if err != nil {
//		return -1, err
//	}
//
//	p, err := strconv.ParseFloat(price, 64)
//	if err != nil {
//		return -1, err
//	}
//
//	return p, nil
//}

//func GetArticleByArticleNr(ctx context.Context, aNr string) (*model.Article, error) {
//
//	a, err := requestNameAndPriceByArctileNr(ctx, aNr)
//	if err != nil {
//		log.Print("no article with number: ", aNr)
//		return nil, err
//	}
//
//	return a, nil
//}
//
func requestNameAndPriceByArctileNr(ctx context.Context, articleNr string) (*model.Article, error) {
	respBody, err := searchArticle(ctx, articleNr)

	name, err := getName(respBody)
	if err != nil {
		return nil, err
	}

	price, err := getPrice(respBody)
	if err != nil {
		return nil, err
	}

	a := &model.Article{
		ArticleNr: articleNr,
		Name:      name,
		Price:     price,
	}

	return a, nil
}
