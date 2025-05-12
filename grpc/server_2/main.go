package main

import (
	"context"
	"fmt"
	"gis/grpc/server_2/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net"
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

			return RunServer(log.Default())
		}()

		if err != nil {
			fmt.Println(err)
		}

	}()
	wg.Wait()
}

type logging struct {
	logger *log.Logger
}

func RunServer(logger *log.Logger) error {
	logger.Println("OK:8082")

	logs := logging{logger}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(logs.grpcLogging))

	pb.RegisterServiceServer(grpcServer, &Service{})

	listener, err := net.Listen("tcp", ":8082")
	if err != nil {
		return err
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		return err
	}

	return nil
}

func (l logging) grpcLogging(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	l.logger.Printf("gRPC method: %s, request: %v", info.FullMethod, req)

	return handler(ctx, req)
}

type controllerFunc func(c *gin.Context) (interface{}, error)

type result struct {
	res interface{}
	err error
}

func handleFunc(f controllerFunc, logger *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		dto, err := f(c)

		c.JSON(http.StatusOK, result{
			dto,
			err,
		})
	}
}

type Service struct {
	pb.UnimplementedServiceServer
}

func (s *Service) Get(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	log.Printf("Получен запрос на пользователя с именем: %s", req.Name)

	return &pb.Response{
		Id:    "1",
		Name:  req.GetName(),
		Email: "ivan@example.com",
	}, nil
}

func Controller(engine *gin.Engine, logger *log.Logger) {

	logger.Println("8082:controller")
	engine.GET(`/api/v1`, handleFunc(v1, logger))
}

func v1(c *gin.Context) (interface{}, error) {

	return nil, nil
}
