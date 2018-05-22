package df

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestModifyForTTS1(t *testing.T) {
	text := `Was kostet der Artikel 'UPPERCASE' mit Extrafuntkion ABC auf otto.de?`

	tl := ModifyForTTS(text)

	assert.Equal(t, `Was kostet der Artikel 'uppercase' mit Extrafuntkion ABC auf otto.de?`, tl)
}

func TestModifyForTTS2(t *testing.T) {
	text := `Was kostet der Artikel UPPERCASE mit Extrafuntkion CD, mit Extrafunktion ABC auf otto.de?`

	tl := ModifyForTTS(text)

	assert.Equal(t, `Was kostet der Artikel uppercase mit Extrafuntkion CD, mit Extrafunktion ABC auf otto.de?`, tl)
}

func TestModifyForTTS3(t *testing.T) {
	text := `Was kostet der Artikel UPPERCASE 180/190 cm auf otto.de?`

	tl := ModifyForTTS(text)

	assert.Equal(t, `Was kostet der Artikel uppercase <say-as interpret-as='cardinal'>180</say-as><say-as interpret-as='cardinal'>190</say-as> cm auf otto.de?`, tl)
}

func TestPriceToString1(t *testing.T) {
	ps := PriceInEuroText(64.00)

	assert.Equal(t, "64 Euro", ps)
}

func TestPriceToString2(t *testing.T) {
	ps := PriceInEuroText(64.15)

	assert.Equal(t, "64.15 Euro", ps)
}

func TestPriceInEuroTTS1(t *testing.T) {
	ps := PriceInEuroTTS(64.00)

	assert.Equal(t, "64 Euro", ps)
}

func TestPriceInEuroTTS2(t *testing.T) {
	ps := PriceInEuroTTS(64.15)

	assert.Equal(t, "64 Euro und 15 Cent", ps)
}

func TestPriceInEuroTTS3(t *testing.T) {
	ps := PriceInEuroTTS(64.99)

	assert.Equal(t, "64 Euro 99", ps)
}

