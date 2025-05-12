package handler

import (
	"gis/ouath2/filter"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	url = "api/test/v1"
)

func Handler(engine *gin.Engine, logger *zap.Logger) {
	engine.GET(url, filter.NewResult(HandleFunc, logger))
}

func HandleFunc(c *gin.Context) (any, error) {

	return "lets goo", nil
}
