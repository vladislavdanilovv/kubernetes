package main

import (
	"context"
	"errors"
	"fmt"
	"gis/controllers"
	"gis/design_pattern/adapter"
	"gis/design_pattern/decorator"
	"gis/design_pattern/singleton"
	"gis/function"
	"gis/infrastructure"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

type Square struct {
	Width  float64
	Height float64
}

func (s Square) Area() float64 {
	return s.Width * s.Height
}

type Sizer interface {
	Area() float64
}

func Less(s1, s2 Sizer) Sizer {
	if s1.Area() < s2.Area() {
		return s1
	}
	return s2
}

type Calculator interface {
	Add() any
}

type numberI interface {
	int8 | float64
}

type number numberI

type Params struct {
	a, b int
}

func (t Params) Add() any {
	return t.a + t.b
}

type myStruct struct {
	param interface{}
}

func testFunc(arg ...int) {
	for _, value := range arg {
		fmt.Println(value)
	}
}

//
//func main() {
//	arg := []int{1, 2, 3, 4}
//
//	testFunc(arg...)
//
//	var arg1 []myStruct
//	var arg2 []myStruct
//	arg2 = append(arg2, arg1...)
//
//	//c := Params{1, 2}
//	//fmt.Println(c.Add())
//	//c := Circle{Radius: 10}
//	//s := Square{Height: 10, Width: 5}
//	//
//	//l := Less(c, s)
//	//fmt.Printf("%+v is the smallest\n", l)
//}

// func main() {
//
//	calc := CalcStruct{}.calc.Add(1, 2)
//	fmt.Println(calc)
//	//tariffRec := Rec{}
//	//tariffRec.CalcTariff(1000)
//	//
//	//fmt.Println(tariffRec)
//
//	//memento.Handler()
//	//ChainOfResponsibility.TestChainOfResponsibility()
//
// }
var wg sync.WaitGroup
var Database *gorm.DB

func main() {
	//connectionString := "slava:Qwerty123!@tcp(localhost:3306)/slava?charset=utf8mb4&parseTime=True&loc=UTC"
	//Database, _ = gorm.Open(mysql.Open(connectionString))

	zapLogger, err := infrastructure.RegisterLogger()
	if err != nil {
		zapLogger.Warn("Ошибка при инициализации логгера")
		os.Exit(1)
	}

	defer zapLogger.Sync()

	gin.SetMode("release")

	server := Handlers(zapLogger)
	wg.Add(1)
	go func() {
		err := func() error {
			defer wg.Done()
			zapLogger.Info("SERVER")
			return server.ListenAndServe()
		}()

		if err != nil {
			fmt.Println(err)
		}

	}()
	wg.Wait()
}

func Handlers(logger *zap.Logger) *http.Server {

	logger.Info("HANDLERS OK")

	engine := gin.New()

	controllers.HandlerTest(engine, logger)

	return &http.Server{
		Addr:         ":8080",
		Handler:      engine,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

}

type appHandler func(c *gin.Context) (any, error)

func HandleError(h appHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := h(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		c.JSON(http.StatusOK, res)
		logs := log.Default()
		logs.Println(res, err)
	}

}

func LoggerTest(c *gin.Context) (any, error) {
	count := 1000
	var mutex sync.RWMutex
	var wg sync.WaitGroup
	//
	for i := 0; i < count; i++ {
		wg.Add(1)

		go info(i+1, &wg, &mutex)

	}

	wg.Wait()

	infrastructure.Logger.Info("test", zap.Field{Key: "ROUTINE IS END", Type: zapcore.StringType, String: "true"})
	return nil, nil
}

func info(number int, wg *sync.WaitGroup, mutex *sync.RWMutex) {

	//mutex.RLock()
	infrastructure.Logger.Info("test", zap.Field{Key: "test", Type: zapcore.StringType, String: strconv.Itoa(number)})

	//mutex.RUnlock()
	//fmt.Println("I = ", number)
	wg.Done()
}

func TestHandler1(c *gin.Context) (any, error) {
	res, err := function.Test1()
	return res, err
}

func TestHandler2(c *gin.Context) (any, error) {
	res, err := function.Test2()

	return res, err
}

func TestFunc(c *gin.Context) (any, error) {
	logs := log.Default()
	logs.Println("TestFunc")
	TestContext(c)

	return nil, nil
}

func TestContext(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)

	defer cancel()

	//select {
	//case <-ctx.Done():
	//	fmt.Println("IM DONE")
	//}
	err := SlowFunc(ctx)
	if errors.Is(err, ctx.Err()) {
		fmt.Println("ctx.Err() ", err)
	}
	if errors.Is(err, context.DeadlineExceeded) {
		fmt.Println("context.DeadlineExceeded ", err)
	}
	if err != nil {
		fmt.Println(err)
	}

}

func SlowFunc(ctx context.Context) error {

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://www.youtube.com", nil)
	if err != nil {
		return err
	}

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(5 * time.Second):
		fmt.Println(res)
	}

	//for i := 0; i < 5; i++ {
	//	select {
	//	case <-ctx.Done():
	//		return fmt.Errorf("Context cancelled: %v\n", ctx.Err())
	//	default:
	//		fmt.Printf("index:%v \n", i+1)
	//
	//		time.Sleep(1 * time.Second)
	//	}
	//
	//}

	return nil
}

func Decorator(c *gin.Context) (any, error) {
	db := InfraStr{}

	decorator.Exec(newLocalDec(db), "test")

	return nil, nil
}

func localDec(s string) {

	fmt.Println(s)
}

type DB interface {
	Store(string) error
}

type InfraStr struct {
}

func (InfraStr) Store(s string) error {
	fmt.Println(s)
	return nil
}

func newLocalDec(db DB) decorator.ExecFunc {
	return func(s string) {
		fmt.Println("db")

		err := db.Store(s)
		if err != nil {
			return
		}

	}
}

func decoratorHand(c *gin.Context) (interface{}, error) {
	//arg := 0
	asd := func() int {
		return rand.Intn(100) + 1
	}

	localFunc := func(number int) int {
		return number
	}(asd())

	fmt.Println(localFunc)

	//decorator.DecoratorFunc()

	//return map[string]string{"test/v5": "response"}, nil
	return nil, errors.New("test/v5")
}

func singletonHand(c *gin.Context) (interface{}, error) {
	singleton.SingletonFunc()

	return nil, nil
}

func adapterHand(c *gin.Context) (interface{}, error) {
	adapter.Adapter()

	return nil, nil
}

func testV5(c *gin.Context) (interface{}, error) {
	in := make(chan int)
	out := make(chan int)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			in <- i + 1
		}
		close(in)
	}()

	go func() {
		defer wg.Done()

		for num := range in {
			out <- doubling(num)
		}
		close(out)
	}()

	for num := range out {
		fmt.Println("out", num)
	}

	wg.Wait()

	return nil, nil
}

func doubling(num int) int {
	return num * num
}
