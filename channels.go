package main

import (
	"fmt"
	"time"
)

func Channel()  {
	messageChannel := make(chan string)

	go func()  {
		for {
			if t := time.Now().Second(); t%7 == 0 {
				messageChannel <- "ping 1"
				messageChannel <- "ping 2"
				time.Sleep(time.Second * 10)
				// messageChannel <- "ping 3"
				fmt.Println(t)
				break				
			}
		}
	}()

	go func()  {
		fmt.Println("Waiting for message")
		msg := <- messageChannel
		fmt.Println("from f1 " + msg)
		fmt.Println("message recieved")
	}()
	go func()  {
		fmt.Println("Waiting for message")
		msg := <- messageChannel
		fmt.Println("from f2 " +msg)
		fmt.Println("message recieved")
	}()
	msg := <- messageChannel
	fmt.Println("from f0 " + msg)
	time.Sleep(time.Minute)
}