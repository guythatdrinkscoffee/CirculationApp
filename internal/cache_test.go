package internal

import (
	"github.com/guythatdrinkscoffee/CirculationApp/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStorage(t *testing.T) {
	st := NewTLLCache()
	assert.NotNil(t, st)
}

func TestTTLCache_Set(t *testing.T) {
	code := "USD"

	//Initialize a new storage
	st := NewTLLCache()

	//Set a new ApiResponse test model
	res := &models.APIResponse{
		BaseCurrencyCode: code,
		BaseCurrencyName: "United States Dollar",
		Rates: map[string]models.Rate{
			"EUR": {
				CurrencyName:  "Euro",
				Rate:          "0.86",
				RateForAmount: "0.86",
			},
		},
	}

	setErr := st.Set(code, res)
	assert.Nil(t, setErr)
}

func TestTTLCache_Get(t *testing.T) {
	code := "USD"

	//Initialize a new storage
	st := NewTLLCache()

	//Set a new ApiResponse test model
	res := &models.APIResponse{
		BaseCurrencyCode: code,
		BaseCurrencyName: "United States Dollar",
		Rates: map[string]models.Rate{
			"EUR": {
				CurrencyName:  "Euro",
				Rate:          "0.86",
				RateForAmount: "0.86",
			},
		},
	}

	//Insert the ApiResponse test case into the Cache
	setErr := st.Set(code, res)
	assert.Nil(t, setErr)

	//Test the GET returns the correct results
	val, err := st.Get(code)

	assert.Nil(t, err)
	assert.NotNil(t, val)
	assert.Equal(t, res.Rates, val.Rates)
}
