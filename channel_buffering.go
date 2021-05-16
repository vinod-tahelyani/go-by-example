package main

import "fmt"

func channel_buffering() {
	bufferedChannel := make(chan string, 2)

	bufferedChannel <- "buffered"
	bufferedChannel <- "channel"

	fmt.Println(<- bufferedChannel)
	fmt.Println(<- bufferedChannel)
}