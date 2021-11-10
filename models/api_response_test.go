package models

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAPIResponse_FromJSON(t *testing.T) {
	resp := &APIResponse{
		BaseCurrencyCode: "USD",
		BaseCurrencyName: "United States Dollar",
		Amount:           "",
		UpdatedDate:      "",
		Rates:            nil,
		Status:           "",
	}

	bodyBytes := new(bytes.Buffer)
	encErr := json.NewEncoder(bodyBytes).Encode(resp)

	assert.Nil(t, encErr)

	testRep := &APIResponse{}
	err := testRep.FromJSON(bodyBytes)

	assert.Nil(t, err)
	assert.Equal(t, resp, testRep)
}
