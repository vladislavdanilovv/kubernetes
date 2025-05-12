package adapter

import "fmt"

type carAdapter struct {
	carTransport *car
}

func (c *carAdapter) Move() {

	c.carTransport.Do()
	fmt.Println("carAdapter is move()")
}
