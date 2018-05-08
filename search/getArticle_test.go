package search

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"time"
	"fmt"
)

func TestCreateArticleFromUrl(t *testing.T) {

	start := time.Now()

	a, err := createArticleFromHtml(nil, testHtmlMay2018)

	end := time.Since(start)

	assert.Nil(t, err)
	assert.Equal(t, "69.99", a.Price)
	assert.Equal(t, "Red Dead Redemption 2 Xbox One", a.Name)
	assert.Equal(t, `https://i.otto.de/i/otto/26390776/red-dead-redemption-2-xbox-one.jpg?$formatz$`, a.ImgUrl)

	fmt.Printf("regex parsing html took %v ms", end.Seconds() * 1000)

}
