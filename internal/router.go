package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/guythatdrinkscoffee/CirculationApp/api"
)

type CirculationRouter struct {
	Router     *gin.Engine
	Controller *api.CirculationController
}

func NewCirculationRouter() CirculationRouter {
	r := gin.Default()
	c := api.NewCirculationController()
	return CirculationRouter{
		Router:     r,
		Controller: c,
	}
}

func (g *CirculationRouter) SetupRoutes() {
	public := g.Router.Group("/")
	{
		//Signup a user
		public.GET("/", nil)
	}

}
