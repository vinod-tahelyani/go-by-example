package main

import "fmt"

func Closures()  {
	
	nextInt := intIterator(2)

	for i := 0; i < 5; i++ {
		fmt.Println(nextInt())
	}

}

func intIterator(start int) func() int {
	return func () int {
		start++
		return start-1
	}
}
