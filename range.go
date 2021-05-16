package main

import "fmt"

func Range() {
	arr := []int{2,3,4,5}
	sum := 0
	for ind, num := range arr {
		fmt.Println(ind)
		sum+=num
	}
	fmt.Println(sum)


	maps := map[string]int{"a": 1, "b": 2}

	for k, v := range maps {
		fmt.Printf("%s -> %d\n", k, v)
	}

	str := "go"

	for _, c := range str {
		fmt.Println(c)
	}
}