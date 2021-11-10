package router

import (
	"github.com/guythatdrinkscoffee/CirculationApp/config"
	"github.com/guythatdrinkscoffee/CirculationApp/internal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCirculationRouter(t *testing.T) {
	c := internal.NewTLLCache()
	m := config.GetConfig()
	cR := NewCirculationRouter(c, m.GIN_MODE)

	assert.NotNil(t, cR)
	assert.Equal(t, m.GIN_MODE, "debug")
}
