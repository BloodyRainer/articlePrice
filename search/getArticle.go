package search

import (
	"net/url"
	"log"
	"io/ioutil"
	"errors"
	"regexp"
	"net/http"
	"strings"
	"github.com/BloodyRainer/articlePrice/model"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
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

func GetRandomArticle(req *http.Request) (*model.Article, error) {
	aNr := model.RandomArticleNr()

	name, price, err := requestNameAndPriceByArctileNr(req, aNr)
	if err != nil {
		log.Print("no product with number: ", aNr)
		return nil, err
	}

	a := &model.Article{
		ArticleNr: aNr,
		Name:      name,
		Price:     price,
	}

	return a, nil

}

func requestNameAndPriceByArctileNr(req *http.Request, articleNr string) (string, string, error) {
	respBody, err := searchArticle(req, articleNr)

	name, err := getName(respBody)
	if err != nil {
		return "", "", err
	}

	price, err := getPrice(respBody)
	if err != nil {
		return "", "", err
	}

	return name, price, nil
}

func getName(body string) (string, error) {
	nameMatch := nameReg.FindStringSubmatch(string(body))
	if len(nameMatch) < 2 || nameMatch[1] == "" {
		return "", errors.New("no name found")
	}

	// TODO: dirty hacks
	name := strings.Replace(nameMatch[1], "&quot;", "\"", -1)
	name = strings.Replace(name, "&amp;", "&", -1)

	return name, nil
}

func getPrice(body string) (string, error) {

	priceMatch := priceReg.FindStringSubmatch(string(body))
	if len(priceMatch) < 2 || priceMatch[1] == "" {
		return "", errors.New("no price found")
	}

	return priceMatch[1], nil
}

func searchArticle(orgReq *http.Request, nr string) (string, error) {
	ctx := appengine.NewContext(orgReq)

	client := urlfetch.Client(ctx)

	query := queryPrefix + nr

	url := &url.URL{
		Scheme:   protocol,
		Host:     domain,
		Path:     path,
		RawQuery: query,
	}

	req := &http.Request{
		Method: "GET",
		URL:    url,
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("request error", err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	return string(body), err
}
