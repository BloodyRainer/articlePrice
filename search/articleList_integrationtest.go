package search

import (
	"time"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine"
	"net/http"
	"github.com/BloodyRainer/articlePrice/model"
)

func IntegrationTestArticleListHandler(w http.ResponseWriter, r *http.Request) {

	ctx := appengine.NewContext(r)
	articles := model.GetArticles()

	badNumbers := make([]string, 0)

	for _, aNr := range articles {
		_, err := requestArticleByArticleNr(ctx, aNr)
		if err != nil {
			badNumbers = append(badNumbers, aNr)
			log.Errorf(ctx, "Error for Article-Number %v", aNr)
		} else {
			log.Infof(ctx, "Success: Found Article-Number %v", aNr)
		}

		time.Sleep(3 * time.Second)
	}

	if len(badNumbers) == 0 {
		log.Infof(ctx, "SUCCESS: Testing of %d articles completed, all articles are available", len(articles))
	} else {
		log.Errorf(ctx, "ERROR: Testing of %d articles completed, list of Article-Numbers with errors is %v", len(articles), badNumbers)
	}

	w.WriteHeader(http.StatusOK)

}
