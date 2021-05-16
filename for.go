package main

import "fmt"

func For()  {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for {
		fmt.Println("Infinite loop")
		break
	}

	for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}