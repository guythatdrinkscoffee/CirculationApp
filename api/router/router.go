package router

import (
	"github.com/gin-gonic/gin"
	"github.com/guythatdrinkscoffee/CirculationApp/api/middlewares"
	"github.com/guythatdrinkscoffee/CirculationApp/api/services"
	"github.com/guythatdrinkscoffee/CirculationApp/internal"
	"log"
)

type CirculationRouter struct {
	Router         *gin.Engine
	Cache          *internal.TTLCache
	CacheValidator *middlewares.CacheValidator
}

func NewCirculationRouter(c *internal.TTLCache, mode string) CirculationRouter {
	//Set the mode for the engine(production or debug)
	gin.SetMode(mode)

	//Define the Router with a default gin engine
	r := gin.Default()

	//Define cache and pass it into the CacheValidator
	cv := middlewares.NewCacheValidator(c)

	return CirculationRouter{
		Router:         r,
		Cache:          c,
		CacheValidator: cv,
	}
}

func (g *CirculationRouter) SetupRoutes() {
	convert := g.Router.Group("/api/v1").Use(g.CacheValidator.CheckCache)
	{
		//Convert a currency to all the available rates with a code parameter in the path
		convert.POST("/convert", func(ctx *gin.Context) {
			log.Println("The uri did not exist in the cache. Make the request")

			//Get the uri
			uri := ctx.Request.RequestURI

			//Get the query params
			base := ctx.Query("from")
			dest := ctx.Query("to")
			amount := ctx.Query("amount")

			//Make the request to the api
			res, err := services.MakeRequestWith(base, dest, amount)

			if err != nil {
				ctx.JSON(500, err)
				ctx.Abort()
				return
			}

			//Set the value in the cache for the uri
			err = g.Cache.Set(uri, res)
			if err != nil {
				ctx.JSON(500, err)
				ctx.Abort()
				return
			}

			ctx.AbortWithStatusJSON(200, res)
		})
	}

}
