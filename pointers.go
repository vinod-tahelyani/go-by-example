package main

import "fmt"

func Pointers()  {
	val, newValue := 2, 10

	fmt.Println(val)

	changeValueNormally(val, newValue)

	fmt.Println(val)

	changeValueWithPointer(&val, newValue)

	fmt.Println(val)

}

func changeValueWithPointer(p *int, newValue int)  {
	*p = newValue
}

func changeValueNormally(val int, newValue int)  {
	val++
	val--
	val = newValue
	val++
	val--
}