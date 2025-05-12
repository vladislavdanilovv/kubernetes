package decorator

import "fmt"

type ExecFunc func(string)

func Exec(exec ExecFunc, s string) {
	exec(s)
}

type decoratorI interface {
	Weight() float32
	Height() int
}

func DecoratorFunc() {
	avgPersonValue := avgPerson{
		73.5,
		180,
	}

	personInfo(avgPersonValue)

	athleticPersonValue := athleticPerson{
		avgValue: avgPersonValue,
	}

	personInfo(athleticPersonValue)
}

type athleticPerson struct {
	avgValue decoratorI
}

func (a athleticPerson) Weight() float32 {
	return a.avgValue.Weight() - 2.5
}

func (a athleticPerson) Height() int {
	return a.avgValue.Height() + 5
}

type avgPerson struct {
	weight float32
	height int
}

func (d avgPerson) Weight() float32 {
	return d.weight
}

func (d avgPerson) Height() int {
	return d.height
}

func personInfo(d decoratorI) {
	fmt.Printf("weight is %1.1fkg, height is %dm \n", d.Weight(), d.Height())
}
