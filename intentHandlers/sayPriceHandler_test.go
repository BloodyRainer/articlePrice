package intentHandlers

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDifferenceGuessActualInPercentPositiv(t *testing.T) {
	actual := 100.00
	guess := 115.00

	diff := differenceGuessActualInPercent(guess, actual)

	assert.Equal(t, 15.00, diff)

}

func TestDifferenceGuessActualInPercentNegativ(t *testing.T) {
	actual := 100.00
	guess := 80.00

	diff := differenceGuessActualInPercent(guess, actual)

	assert.Equal(t, -20.00, diff)

}

func TestDifferenceGuessActualInPercentHighNegativ(t *testing.T) {
	actual := 100.00
	guess := 10.00

	diff := differenceGuessActualInPercent(guess, actual)

	assert.Equal(t, -90.00, diff)
}