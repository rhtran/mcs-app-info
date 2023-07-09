package info

import (
	"go.uber.org/zap"
	"net"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//type ipLookupService interface {
//	LookupIP(host string) ([]net.IP, error)
//}

type Router struct {
	logger      *zap.Logger
	infoService Servicer
}

func NewInfoRouter(logger *zap.Logger, infoService Servicer) *Router {
	return &Router{logger,
		infoService}
}

// ShowInfo godoc
// @Summary Show config info
// @Tags Info
// @Produce  json
// @Router /info [get]
// @Success 200 {object} info.Info
func (infoRouter *Router) InfoRegister(router *gin.RouterGroup) {
	router.GET("", infoRouter.appInfo)
}

func (infoRouter *Router) appInfo(c *gin.Context) {
	result, err := infoRouter.infoService.GetAppInfo()

	// this line make this function mat testable.
	addrs, err := net.LookupIP("localhost")
	if err != nil {
		result.Ip = "Unknown host"
	} else {
		result.Ip = ""
		for _, ia := range addrs {
			result.Ip += ia.String() + " "
		}
		result.Ip = strings.TrimSpace(result.Ip)
	}

	c.JSON(http.StatusOK, result)
}
