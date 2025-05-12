package filter

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
)

type Result func(c *gin.Context) (any, error)

type filterStr struct {
	Error  error `json:"error,omitempty"`
	Result any   `json:"result,omitempty"`
}

func NewResult(do Result, logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		a, err := do(c)
		if err != nil {
			logger.Error(c.FullPath(), zap.Error(err))
			c.JSON(http.StatusBadRequest, filterStr{Error: err})
		}

		logger.Info(c.FullPath(), zap.Field{Key: "message", Type: zapcore.StringType, String: "success"})

		c.JSON(http.StatusOK, filterStr{Result: a})
	}
}
