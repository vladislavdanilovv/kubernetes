package controllers

import (
	"gis/infrastructure"
	"gis/internal/domain"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
)

type filterStr struct {
	Error  string      `json:"error,omitempty"`
	Result interface{} `json:"result,omitempty"`
}
type dcFunc func(c *gin.Context) (interface{}, error)

func filter(do dcFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		arg := filterStr{}

		result, err := do(c)

		if err != nil {
			arg.Error = err.Error()

			c.JSON(http.StatusInternalServerError, arg)

			infrastructure.Logger.Error(c.FullPath(), zap.Error(err))
			return
		}

		arg.Result = result

		infrastructure.Logger.Info("success", zap.Field{Key: "api", Type: zapcore.StringType, String: c.FullPath()})
		c.AbortWithStatusJSON(http.StatusOK, arg)

		return
	}

}

const url = "test"

func HandlerTest(controller *gin.Engine, logger *zap.Logger) {
	//controller.GET(main.url+"/v1", filter(main.TestFunc))
	//controller.GET(main.url+"/v2", filter(main.TestHandler2))
	//controller.GET(main.url+"/v3", filter(main.LoggerTest))
	//controller.GET(main.url+"/v4", filter(main.Decorator))
	//controller.GET(main.url+"/v5", filter(main.testV5))
	//controller.GET(main.url+"/v6", filter(main.singletonHand))
	//controller.GET(main.url+"/v7", filter(main.adapterHand))
	controller.GET(url+"/v8", filter(domain.GeneratorStart))
	controller.GET(url+"/v9", filter(a))
	controller.GET(url+"/v10", filter(test))

	logger.Info("HANDLERS READY")
}

func a(c *gin.Context) (any, error) {
	domain.Foo()
	return nil, nil
}

func test(c *gin.Context) (any, error) {
	return "HELLO FROM KUBE", nil
}
