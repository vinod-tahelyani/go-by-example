package main

import "fmt"

func Arrays()  {
	var a[4] int

	a[0] = 2

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	b := [2]int{2,3}

	for i := 0; i < len(b); i++ {
		fmt.Println(b[i])
	}


	var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}