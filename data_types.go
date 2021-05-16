package main

import "fmt"

func VariousDataTypes() {
	fmt.Println("go1" + "go2")

	fmt.Println(1 + 2.0)
	fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)
}