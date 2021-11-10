package services

import (
	"fmt"
	"github.com/guythatdrinkscoffee/CirculationApp/config"
	"github.com/guythatdrinkscoffee/CirculationApp/models"
	"io"
	"log"
	"net/http"
)

var (
	endpoints *models.Endpoint
)

func init() {
	endpoints = models.NewEndpoints()
}

func MakeRequestWith(base string, dest string, amount string) (*models.APIResponse, error) {
	if len(dest) == 0 && len(amount) == 0 {
		//No destination currency or amount was passed so
		//make an api request to convert a single currency to
		//all of the available currencies.
		return convertFromToAll(base)
	} else {
		if len(amount) == 0 {
			amount = "1"
		}
		return convertFromToWithAmount(base, dest, amount)
	}
}

// ConvertFromToAll Returns all the exchange rates for the provided currency code.
func convertFromToAll(code string) (*models.APIResponse, error) {
	//Build the request url`
	reqUrl := fmt.Sprintf("%s&from=%s&amount=1", endpoints.Convert, code)

	//Define the request
	req, err := buildRequest("GET", reqUrl)

	if err != nil {
		return nil, err
	}

	//Make the request through the default http client.
	res, reqErr := http.DefaultClient.Do(req)

	//If the request results in an error then return
	if reqErr != nil {
		return nil, reqErr
	}

	//Close the res
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(res.Body)

	results := &models.APIResponse{}

	//Decode the body into a APIResponse
	err = results.FromJSON(res.Body)

	if err != nil {
		return nil, err
	}

	return results, nil
}

// ConvertFromToWithAmount Converts one currency to another with a given amount
func convertFromToWithAmount(base string, derived string, amount string) (*models.APIResponse, error) {
	//Build the request url
	reqUrl := fmt.Sprintf("%s&from=%s&to=%s&amount=%s", endpoints.Convert, base, derived, amount)

	//Define the request
	req, err := buildRequest("GET", reqUrl)

	if err != nil {
		return nil, err
	}

	//Make the request
	res, reqErr := http.DefaultClient.Do(req)

	if reqErr != nil {
		return nil, reqErr
	}

	//Close the res
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(res.Body)

	results := &models.APIResponse{}

	err = results.FromJSON(res.Body)

	if err != nil {
		return nil, err
	}

	return results, nil
}

//A util function to apply the needed headers and build the request
func buildRequest(method string, url string) (*http.Request, error) {
	c := config.GetConfig()
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}

	//add the appropriate headers
	//Add the application specific headers
	req.Header.Add("x-rapidapi-host", "currency-converter5.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", c.APP_KEY)

	return req, nil
}
