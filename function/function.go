package function

import "fmt"

func Test1() (string, error) {
	return "TEST11111", nil
}

func Test2() (int, error) {
	arr := []int{1, 2, 3, 4}
	src := arr[:1]

	src = append(src, 11)
	//foo(src)

	fmt.Println(arr)
	fmt.Println(src)
	return 22222222, nil
}

func foo(src []int) {
	src = append(src, 10)
}
