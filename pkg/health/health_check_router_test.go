package health

import (
	"blog-service/internal/api/test"
	"context"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/qiangxue/go-restful-api/pkg/log"
	"github.com/stretchr/testify/mock"
)

type healthServiceMock struct {
	mock.Mock
}

func (hsm *healthServiceMock) Check() (*Check, error) {
	args := hsm.Called()
	intf := args.Get(0)
	val := intf.(Check)
	return &val, args.Error(1)
}

func TestHealthRouter(t *testing.T) {
	healthMock := healthServiceMock{}
	health := &Check{"UP"}
	healthMock.On("Check").Return(*health, nil)

	logger := log.New().With(context.TODO(), "version", "test")
	r := gin.Default()
	home := r.Group("/")

	router := NewCheckRouter(&healthMock, logger)
	router.CheckRegister(home.Group("/health"))

	tests := []test.APITestCase{
		{"get config health", "GET", "/health", "", nil, http.StatusOK, `*UP*`},
	}
	for _, tc := range tests {
		test.Endpoint(t, r, tc)
	}
}
