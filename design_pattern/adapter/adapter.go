package adapter

import "fmt"

type transport interface {
	Move()
}

type person struct {
	name string
}

func (p *person) Do(tr transport) {

	fmt.Printf("person name is %v start do()", p.name)

	tr.Move()
}

type car struct {
}

type plane struct {
}

func (car) Do() {
	fmt.Println(" car do() ")
}

func (plane) Move() {
	fmt.Println(" plane is move()")

}

func Adapter() {
	personValue := &person{name: "chel"}

	planeValue := &plane{}

	personValue.Do(planeValue)

	carValue := &car{}

	carValueAdapter := &carAdapter{carTransport: carValue}

	personValue.Do(carValueAdapter)
}
