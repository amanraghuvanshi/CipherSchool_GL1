package main

// func main() {
// 	channel := make(chan string, 1)
// 	// buffered channel
// 	go func(ch chan<- string) {
// 		{
// 			ch <- "2" //sending the data in the channel
// 			fmt.Println("1")
// 		}
// 	}(channel)
// 	message := <-channel // listen for data
// 	time.Sleep(time.Second * 2)
// 	fmt.Println(message)
// }

//------------------------------------------------

// func Main1() {
// 	channel1 := make(chan string, 1)
// 	go func(ch <-chan string) {
// 		{
// 			msg := <-ch
// 			fmt.Println(msg)
// 			fmt.Println("1")
// 		}
// 	}(channel1)
// 	message := "Hello, from the main1 function" // This statement is being sent to the goroutine that will be sending this message
// 	channel1 <- message                         // There the channel is receiving the message
// 	time.Sleep(time.Second * 5)
// 	fmt.Println("Done")
// }

// ===============================================

// func main() {
// 	channel := make(chan string, 1) //initializing channel
// 	go func(ch chan string) {
// 		mess := <-ch
// 		ch <- "Hello from anonymous channel"
// 		fmt.Println(mess)
// 		fmt.Println("1")
// 	}(channel) //anonymous goroutine
// 	message := "Hello, from the main3 function"
// 	channel <- message
// 	time.Sleep(time.Second * 2)
// 	message = <-channel
// 	fmt.Println(message)
// }

// SELECT STATEMENTS
// This is kind of switch statement, but it works for the channels and determines which channel has to be executed

// NOTES:

//In the birectional channels, we can do either of three things, we can send data, we can receive data or we can do both

// -> The first function has a sending channel, and it's the first of kind in the unidirectional channels

// -> Being a buffered channel, the first function will still execute even if their is no listener, since the buffer has the capacity to hold one message, so it will hold one message and let the program do its thing

// For birectional channels, first we are sending a message to the channel, now inside the goroutine the message will be sent to the channel and afterwards, a message will be received by the channel from the anonymous function, followed by a message.
