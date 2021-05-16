package main

import "fmt"

func Functions()  {
	a := add(2, 3)
	b, c := multipleReturnValuesfunc()
	d := variadicFunc([]int{2,3}...)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

func add(a, b int) int {
	return a + b
}

func multipleReturnValuesfunc() (int, int) {
	return 3, 6
}

func variadicFunc(a... int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}