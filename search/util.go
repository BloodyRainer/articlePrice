package search

import (
	"google.golang.org/appengine/urlfetch"
	"net/url"
	"io/ioutil"
	"google.golang.org/appengine"
	"net/http"
	"log"
	"errors"
	"strings"
)

// Appengine needs the original request.
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

func getPrice(body string) (string, error) {

	priceMatch := priceReg.FindStringSubmatch(string(body))
	if len(priceMatch) < 2 || priceMatch[1] == "" {
		return "", errors.New("no price found")
	}

	return priceMatch[1], nil
}

func getName(body string) (string, error) {
	nameMatch := nameReg.FindStringSubmatch(string(body))
	if len(nameMatch) < 2 || nameMatch[1] == "" {
		return "", errors.New("no name found")
	}

	// TODO: dirty hacks
	name := strings.Replace(nameMatch[1], "&quot;", `"`, -1)
	name = strings.Replace(name, "&amp;", "&", -1)

	return name, nil
}
