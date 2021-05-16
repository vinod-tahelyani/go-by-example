package main

import (
	"fmt"
	"time"
)

func GoRoutines()  {
	f("direct")

	go f("goroutine")

	go func(msg string)  {
		fmt.Println(msg)
	}("anonymous function")

	time.Sleep(time.Second)
	fmt.Println("done")
}

func f(str string)  {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d- %s\n", i, str)
	}
}