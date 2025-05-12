package main

import (
	"fmt"
	"gis/grpc/server_2/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		err := func() error {
			defer wg.Done()
			log.Println("SERVER")
			return RunServer(log.Default()).ListenAndServe()
		}()

		if err != nil {
			fmt.Println(err)
		}

	}()
	wg.Wait()
}

func RunServer(logger *log.Logger) *http.Server {
	logger.Println("OK:8081")
	gin.SetMode("release")
	engine := gin.New()

	Controller(engine, logger)

	return &http.Server{
		Addr:    ":8081",
		Handler: engine,
	}
}

type controllerFunc func(c *gin.Context) (interface{}, error)

type result struct {
	Res interface{}
	err error
}

func handleFunc(f controllerFunc, logger *log.Logger) gin.HandlerFunc {

	return func(c *gin.Context) {
		dto, err := f(c)

		logger.Println("request", c.Request.RequestURI)
		logger.Println("dto", dto)

		c.JSON(http.StatusOK, result{
			dto,
			err,
		})
	}
}

func Controller(engine *gin.Engine, logger *log.Logger) {
	engine.GET(`/api/v1`, handleFunc(v1, logger))
}

func v1(c *gin.Context) (interface{}, error) {

	conn, err := grpc.NewClient("localhost:8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Ошибка подключения: %v", err)
	}
	defer conn.Close()

	client := pb.NewServiceClient(conn)

	response, err := client.Get(c, &pb.Request{Name: "тестовый"})

	fmt.Println(response, err)

	return client.Get(c, &pb.Request{Name: "1"})
}
