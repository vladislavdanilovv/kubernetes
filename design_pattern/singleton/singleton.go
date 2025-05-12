package singleton

import (
	"fmt"
	"sync"
)

type SingleVar map[string]string

var (
	once sync.Once

	instance SingleVar
)

func New(key, value string) SingleVar {
	once.Do(func() {
		instance = make(SingleVar)
		instance[key] = value
	})

	return instance
}

func SingletonFunc() {
	key := "key"
	value := "value"
	arg := New(key, value)

	fmt.Println(arg)

	key2 := "key2"
	value2 := "value2"

	arg2 := New(key2, value2)

	//arg["asd"] = "variant_two"
	fmt.Println(arg2)
}
