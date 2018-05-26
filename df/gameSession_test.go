package df

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFloatArray2String(t *testing.T) {
	fa := []float64{59.99, 79.89,30.00}

	result := floatArray2String(fa)

	assert.Equal(t, "[59.99,79.89,30.00]", result)
}
