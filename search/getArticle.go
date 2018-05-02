package search

import (
	"log"
	"regexp"
	"net/http"
	"github.com/BloodyRainer/articlePrice/model"
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
func GetRandomArticle(req *http.Request) (*model.Article, error) {
	aNr := model.RandomArticleNr()

	a, err:= requestNameAndPriceByArctileNr(req, aNr)
	if err != nil {
		log.Print("no article with number: ", aNr)
		return nil, err
	}

	return a, nil
}

func GetArticleByArticleNr(req *http.Request, aNr string) (*model.Article, error){

	a, err:= requestNameAndPriceByArctileNr(req, aNr)
	if err != nil {
		log.Print("no article with number: ", aNr)
		return nil, err
	}

	return a, nil
}

func requestNameAndPriceByArctileNr(req *http.Request, articleNr string) (*model.Article, error) {
	respBody, err := searchArticle(req, articleNr)

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
