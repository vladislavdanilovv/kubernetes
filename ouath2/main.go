package main

import (
	"fmt"
	"gis/ouath2/handler"
	middleware2 "gis/ouath2/middleware"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"net/http"
)

type Server struct {
	main   *errgroup.Group
	engine *gin.Engine
	logger *zap.Logger
}

func newServer() *Server {
	logger, err := registerLogger()
	if err != nil {
		panic(fmt.Sprintf("logger error %v", err))
	}
	return &Server{
		main:   &errgroup.Group{},
		engine: gin.New(),
		logger: logger,
	}
}

func (r *Server) registerHandlers() *Server {
	handler.Handler(r.engine, r.logger)

	return r
}
func (r *Server) register() *Server {
	r.logger.Info("Server:Register")
	r.engine.Use(middleware2.CorsCredentials())

	r.registerHandlers()

	serve := &http.Server{
		Addr:    ":8080",
		Handler: r.engine,
	}

	r.main.Go(func() error {
		return serve.ListenAndServe()
	})

	return r
}

func (r *Server) run() error {
	r.logger.Info("Server:Run")
	return r.main.Wait()
}

func main() {

	gin.SetMode("release")

	newServer().register().run()
}
