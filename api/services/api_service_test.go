package services

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/guythatdrinkscoffee/CirculationApp/config"
	"github.com/guythatdrinkscoffee/CirculationApp/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type APITestSuite struct {
	suite.Suite
	APIClient *resty.Client
	Config    *config.Config
}

func (suite *APITestSuite) SetupTest() {
	suite.APIClient = resty.New()
	suite.Config = config.GetConfig()
}

func TestAPISuite(t *testing.T) {
	suite.Run(t, new(APITestSuite))
}

//Make a request with a base currency and convert to all possible currencies
func (suite *APITestSuite) TestMakeRequestWith() {
	key := suite.Config.APP_KEY

	req := suite.APIClient.R()
	req.Header.Add("x-rapidapi-host", "currency-converter5.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", key)

	res, _ := req.Get("https://currency-converter5.p.rapidapi.com/currency/convert?format=json&from=usd")

	var apiRes *models.APIResponse

	err := json.Unmarshal(res.Body(), &apiRes)
	assert.Nil(suite.T(), err)

	//Check the response corresponds to the currency code passed to the query
	assert.Equal(suite.T(), "USD", apiRes.BaseCurrencyCode)
	assert.NotNil(suite.T(), apiRes)
	assert.Greaterf(suite.T(), len(apiRes.Rates), 0, "error message %s", "formatted")

	//Timer is needed if the API plan is on basic. The basic plan only allows one call per second
	//Therefore, the testing suite makes a call and waits a second between the followi
	time.Sleep(1 * time.Second)
}

//Make a request with a base, destination and amount.
func (suite *APITestSuite) TestMakeRequestWith2() {
	key := suite.Config.APP_KEY
	req := suite.APIClient.R()

	req.Header.Add("x-rapidapi-host", "currency-converter5.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", key)

	res, _ := req.Get("https://currency-converter5.p.rapidapi.com/currency/convert?format=json&from=CAD&to=USD&amount=20")

	var apiRes *models.APIResponse

	err := json.Unmarshal(res.Body(), &apiRes)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), len(apiRes.Rates), 1)
	assert.NotNil(suite.T(), apiRes.Rates["USD"])
}
