package services

import (
	"fmt"
	"github.com/guythatdrinkscoffee/CirculationApp/models"
	"net/http"
	"os"
)

var (
	endpoints *models.Endpoint
)

func init() {
	endpoints = models.NewEndpoints()
}

// GetAllRatesFor GetAllRates Returns all the exchange rates for the provided currency code.
func GetAllRatesFor(code string) (*models.Result, error) {
	//Build the request url
	reqUrl := fmt.Sprintf("%s&from=%s&amount=1", endpoints.Convert, code)

	//Define the request
	req, _ := http.NewRequest("GET", reqUrl, nil)

	//Add the application specific headers
	req.Header.Add("x-rapidapi-host", "currency-converter5.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", os.Getenv("APP_KEY"))

	//Make the request throught the default http client.
	res, err := http.DefaultClient.Do(req)

	//If the request results in an error then return
	if err != nil {
		return nil, err
	}

	//Close the res
	defer res.Body.Close()

	results := &models.Result{}

	//Decode the body into a Result
	err = results.FromJSON(res.Body)

	if err != nil {
		return nil, err
	}

	return results, nil
}
