package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEndpoints(t *testing.T) {
	ep := NewEndpoints()
	convertEp := "https://currency-converter5.p.rapidapi.com/currency/convert?format=json"
	assert.Equal(t, ep.Convert, convertEp)
}
