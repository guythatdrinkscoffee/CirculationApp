package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/guythatdrinkscoffee/CirculationApp/api/controllers"
)

type CirculationRouter struct {
	Router     *gin.Engine
	Controller *controllers.CirculationController
}

func NewCirculationRouter() CirculationRouter {
	r := gin.Default()
	c := controllers.NewCirculationController()
	return CirculationRouter{
		Router:     r,
		Controller: c,
	}
}

func (g *CirculationRouter) SetupRoutes() {
	api := g.Router.Group("/api/v1")
	{
		//Convert a currency to all the available rates with a code parameter in the path
		api.GET("/convert/:code", nil)

	}

}
