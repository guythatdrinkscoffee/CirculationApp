package router

import (
	"github.com/gin-gonic/gin"
	"github.com/guythatdrinkscoffee/CirculationApp/api/services"
	"github.com/guythatdrinkscoffee/CirculationApp/internal"
	"github.com/guythatdrinkscoffee/CirculationApp/models"
	"log"
)

type CirculationRouter struct {
	Router         *gin.Engine
	Cache          *internal.TTLCache
	CacheValidator *CacheValidator
}

func NewCirculationRouter(s *internal.TTLCache) CirculationRouter {
	r := gin.Default()
	c := NewCacheValidator(s)
	return CirculationRouter{
		Router:         r,
		Cache:          s,
		CacheValidator: c,
	}
}

func (g *CirculationRouter) SetupRoutes() {
	convertToAllRoute := g.Router.Group("/api/v1/convert").Use(g.CacheValidator.CheckCache)
	{
		//Convert a currency to all the available rates with a code parameter in the path
		convertToAllRoute.GET("/:code", func(ctx *gin.Context) {
			//Print that the request was made.
			log.Println("Request was made")

			//Grab the currency code from the Params
			code := ctx.Param("code")

			//Make the call to the api with the currency code
			results, err := services.ConvertFromToAll(code)

			//If the api call results in an error then abort
			if err != nil {
				ctx.JSON(500, err)
				ctx.Abort()
				return
			}

			_ = g.Cache.Set(code, results)

			res := &models.CirculationResponse{
				BaseCurrencyCode: results.BaseCurrencyCode,
				Rates:            results.Rates,
			}

			ctx.AbortWithStatusJSON(200, res)
		})

	}

}
