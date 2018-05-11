package search

import (
	"testing"
	"github.com/stretchr/testify/assert"
)



func TestGetImgUrl(t *testing.T) {
	img, err := getImageUrl(testHtmlMay2018)

	assert.Nil(t, err)
	assert.Equal(t, "https://i.otto.de/i/otto/26390776/red-dead-redemption-2-xbox-one.jpg?$formatz$", img)
}

func TestGetPrice(t *testing.T) {
	price, err := getPrice(testHtmlMay2018)

	assert.Nil(t, err)
	assert.Equal(t, "69.99", price)
}

func TestGetName(t *testing.T) {
	name, err := getName(testHtmlMay2018)

	assert.Nil(t, err)
	assert.Equal(t, "Red Dead Redemption 2 Xbox One", name)
}

func TestCustomizeImgUrl(t *testing.T) {
	oldUrl := "https://i.otto.de/i/otto/26390776/red-dead-redemption-2-xbox-one.jpg?$formatz$"

	newUrl := customizeImgUrl(oldUrl)

	assert.Equal(t, "https://i.otto.de/i/otto/26390776?h=520&amp;w=384&amp;sm=clamp", newUrl)
}


