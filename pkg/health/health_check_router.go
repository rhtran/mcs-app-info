package health

import (
	"go.uber.org/zap"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CheckRouter struct {
	checkService CheckServicer
	logger       *zap.Logger
}

func NewCheckRouter(checkService CheckServicer, logger *zap.Logger) *CheckRouter {
	return &CheckRouter{checkService, logger}
}

func (checkRouter *CheckRouter) CheckRegister(router *gin.RouterGroup) {
	router.GET("", checkRouter.healthCheck)
}

func (checkRouter *CheckRouter) healthCheck(c *gin.Context) {
	result, err := checkRouter.checkService.Check()

	if err != nil {
		c.JSON(http.StatusBadRequest, result)
	} else {
		c.JSON(http.StatusOK, result)
	}
}
