package router

import (
	"github.com/gin-gonic/gin"
	"github.com/guythatdrinkscoffee/CirculationApp/internal"
	"log"
)

type CacheValidator struct {
	Storage *internal.TTLCache
}

func NewCacheValidator(s *internal.TTLCache) *CacheValidator {
	return &CacheValidator{
		Storage: s,
	}
}

func (c CacheValidator) CheckCache(ctx *gin.Context) {
	//Grab the uri from the request
	uri := ctx.Request.RequestURI

	log.Println(uri)
	//Check the cache if the uri exists
	res, err := c.Storage.Get(uri)

	//If err is not nil, then the uri does not exist in the
	//cache. Call ctx.next() to make the appropriate request
	//for the uri.
	if err != nil {
		ctx.Next()
		return
	}

	//The uri did exist in the cache so return the value
	ctx.AbortWithStatusJSON(200, res)
}
