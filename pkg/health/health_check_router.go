package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	checkService CheckServicer
}

func NewCheckRouter(checkService CheckServicer) *Router {
	return &Router{checkService}
}

func (checkRouter *Router) CheckRegister(router *gin.RouterGroup) {
	router.GET("/health", checkRouter.healthCheck)
}

func (checkRouter *Router) healthCheck(c *gin.Context) {
	result, err := checkRouter.checkService.Check()

	if err != nil {
		c.JSON(http.StatusBadRequest, result)
	} else {
		c.JSON(http.StatusOK, result)
	}
}
