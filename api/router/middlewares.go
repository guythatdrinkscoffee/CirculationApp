package router

import (
	"github.com/gin-gonic/gin"
	"github.com/guythatdrinkscoffee/CirculationApp/Utils"
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
	code := ctx.Param("code")

	if !Utils.CodeIsValid(code) {
		ctx.JSON(400, "Invalid currency code")
		ctx.Abort()
		return
	}

	res, err := c.Storage.Get(code)

	if err != nil {
		log.Println("Response not found in cache")
		ctx.Next()
		return
	}

	ctx.AbortWithStatusJSON(200, res.Rates)
}
