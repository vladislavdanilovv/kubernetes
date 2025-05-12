package domain

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

func GeneratorStart(c *gin.Context) (any, error) {
	done := make(chan struct{})
	defer close(done)
	in := make(chan int)

	numbers := []int{1, 2, 3, 4, 5, 6}

	go func() {
		for _, number := range numbers {
			fmt.Println("in", number)
			in <- number
		}
		close(in)
	}()

	for result := range AddMultiply(done, Add(done, in)) {
		fmt.Println("result", result)
	}

	return nil, nil
}

func Add(done <-chan struct{}, in <-chan int) chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for num := range in {
			select {
			case <-done:
				return
			case out <- num + 1:
			}
		}

	}()
	return out
}

func AddMultiply(done <-chan struct{}, in <-chan int) chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for num := range in {
			select {
			case <-done:
				return
			case out <- num * num:
			}
		}
	}()
	return out
}

func Foo() {
	now := time.Now()
	doneSleep := make(chan struct{})
	done := make(chan struct{})

	go sleep(doneSleep)

	go g(doneSleep, done)

	<-done

	fmt.Printf("done in %v seconds\n", time.Since(now).Seconds())

}

func sleep(doneSleep chan<- struct{}) {
	defer close(doneSleep)

	s := rand.Intn(10)

	fmt.Printf("some operation for %v seconds\n", s)

	select {
	case <-time.After(time.Duration(s) * time.Second):
		doneSleep <- struct{}{}
		return
	}
}

func g(doneSleep <-chan struct{}, done chan<- struct{}) {
	defer close(done)
	bs := "a b c"
	//bs := []byte("◺")
	select {

	case <-doneSleep:
		done <- struct{}{}
		fmt.Println([]byte(bs), bs)
		fmt.Println("routine is start")
		return
	}

}
