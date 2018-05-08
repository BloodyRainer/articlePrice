package search

import (
	"google.golang.org/appengine/urlfetch"
	"net/url"
	"io/ioutil"
	"net/http"
	"errors"
	"strings"
	"context"
	"regexp"
	"time"
	"google.golang.org/appengine/log"
	"fmt"
)

var priceReg *regexp.Regexp
var nameReg *regexp.Regexp
var imgUrlReg *regexp.Regexp

func init() {
	priceReg = regexp.MustCompile(`itemprop="price".*content="(\d+\.?\d+)"`)
	nameReg = regexp.MustCompile(`<h1.*itemprop="name".*>(.*)</h1>`)
	//TODO: actual img-url: https://i.otto.de/i/otto/26390776/red-dead-redemption-2-xbox-one.jpg?$formatz$
	//TODO: example of how to customize format https://i.otto.de/i/otto/26390776?h=520&amp;w=384&amp;sm=clamp
	imgUrlReg = regexp.MustCompile(`<meta.*name="twitter:image".*content="(.*?)".*>`)
	//imgUrlReg = regexp.MustCompile(`<img.*id="prd_mainProductImage".*src="(.*?)".*>`)
}

// Appengine needs the original request.
func searchArticle(ctx context.Context, nr string) (string, error) {

	client := urlfetch.Client(ctx)

	query := queryPrefix + nr

	url := &url.URL{
		Scheme:   protocol,
		Host:     domain,
		Path:     path,
		RawQuery: query,
	}

	//GET to otto.de
	start := time.Now()
	req := &http.Request{
		Method: "GET",
		URL:    url,
	}
	durStr := fmt.Sprintf("otto GET took %v ms", time.Since(start).Seconds() * 1000)
	log.Infof(ctx, durStr)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	return string(body), err
}

func getImageUrl(body string) (string, error) {
	imgUrlMatch := imgUrlReg.FindStringSubmatch(string(body))

	if len(imgUrlMatch) < 2 || imgUrlMatch[1] == "" {
		return "", errors.New("no image-url found")
	}
	return imgUrlMatch[1], nil
}

func getPrice(body string) (string, error) {

	priceMatch := priceReg.FindStringSubmatch(string(body))
	if len(priceMatch) < 2 || priceMatch[1] == "" {
		return "", errors.New("no price found")
	}

	return priceMatch[1], nil
}

func getName(body string) (string, error) {
	nameMatch := nameReg.FindStringSubmatch(body)
	if len(nameMatch) < 2 || nameMatch[1] == "" {
		return "", errors.New("no name found")
	}

	// TODO: dirty hacks
	name := strings.Replace(nameMatch[1], "&quot;", `"`, -1)
	name = strings.Replace(name, "&amp;", "&", -1)
	name = strings.Replace(name, "ä", "ae", -1)
	name = strings.Replace(name, "ö", "oe", -1)
	name = strings.Replace(name, "ü", "ue", -1)
	name = strings.Replace(name, "ß", "ss", -1)
	name = strings.Replace(name, "«", "'", -1)
	name = strings.Replace(name, "»", "'", -1)
	name = strings.Replace(name, "™", "", -1)
	name = strings.Replace(name, "®", "", -1)

	return name, nil
}
