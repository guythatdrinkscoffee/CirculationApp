package middlewares

import (
	"github.com/guythatdrinkscoffee/CirculationApp/internal"
	"github.com/guythatdrinkscoffee/CirculationApp/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCacheValidator(t *testing.T) {
	//Init a new TTLCache
	s := internal.NewTLLCache()

	//The CacheValidator takes in a TTLCache
	c := NewCacheValidator(&s)

	assert.NotNil(t, c)
}

func TestCacheValidator_CheckCache(t *testing.T) {
	s := internal.NewTLLCache()

	c := NewCacheValidator(&s)

	r := &models.APIResponse{
		BaseCurrencyCode: "EUR",
		BaseCurrencyName: "Euro",
		Amount:           "20.00",
		UpdatedDate:      "",
		Rates:            nil,
		Status:           "",
	}

	//Set the value in the cache
	e := c.Storage.Set("EUR", r)

	assert.Nil(t, e)

	//Get the value in the cache
	res, err := c.Storage.Get("EUR")

	assert.Nil(t, err)
	assert.Equal(t, res.BaseCurrencyName, r.BaseCurrencyName)
	assert.Equal(t, res.BaseCurrencyCode, r.BaseCurrencyCode)
}
