package main

import (
	"fmt"
	"time"
)

// func main() {
// 	// Sending multiple data on the channel
// 	channel := make(chan int)
// 	go func() {
// 		channel <- 1
// 		time.Sleep(time.Second * 1)
// 		channel <- 2
// 		time.Sleep(time.Second * 1)
// 		channel <- 3
// 		close(channel)
// 		// here the main function is waiting for the channel to receive something
// 		// So, if we are sending multiple messages in a channel, then you need to close the channel, this allows the main function not to wait for any recieving messages
// 	}()
// 	for ch := range channel {
// 		fmt.Println(ch)
// 	}
// }

// func main() {
// 	helloCh := make(chan string, 1)
// 	goodByeCh := make(chan string, 1)
// 	quitCh := make(chan bool)

// 	go ReceiveMsg(helloCh, goodByeCh, quitCh)
// 	go SendMsg(helloCh, "Hello, world!")
// 	time.Sleep(time.Second * 1)
// 	go SendMsg(goodByeCh, "Goodbye, world!")
// 	res := <-quitCh
// 	fmt.Println("Message from quitCh:", res)
// }

func SendMsg(ch chan<- string, message string) {
	// This is a sender goroutine
	ch <- message
}

func ReceiveMsg(helloCh, goodByeCh <-chan string, quitCh chan<- bool) {
	for {
		select {
		case message := <-helloCh:
			fmt.Println("Message from helloCh:", message)
		case message := <-goodByeCh:
			fmt.Println("Goodbye from goodByeCh:", message)
		case <-time.After(time.Second * 2):
			fmt.Println("Nothing Received in 2seconds, Exiting the receiveMsg goroutine")
			quitCh <- true
			return
		}
	}
}
