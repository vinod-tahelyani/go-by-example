package main

import (
	"fmt"
	"time"
)

func selectWithChannel()  {
	chan1 := make(chan string, 1)
	chan2 := make(chan string, 1)

	go goRoutien(chan1, 5, "first message")
	go goRoutien(chan2, 15, "second message")

	for i := 0; i < 2; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
		if i == 16 {
			break
		}
		select {
			case msg := <- chan1:
				fmt.Println(msg)
			case msg := <- chan2:
				fmt.Println(msg)
		}
	}
}


func goRoutien(channel chan<- string, sleep int64, msg string) {
	time.Sleep(time.Duration(sleep * 1e9))
	channel <- msg
}