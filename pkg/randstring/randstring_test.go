package randstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithLength(t *testing.T) {
	assert := assert.New(t)

	stringLength := 10
	randString := WithLength(10)	

	assert.Equal(
		stringLength,
		len(randString),
		"string's length should be of desired length",
	)

	randString2 := WithLength(10)
	assert.NotEqual(
		randString,
		randString2,
		"2 generated string should not be the same",
	)
}
