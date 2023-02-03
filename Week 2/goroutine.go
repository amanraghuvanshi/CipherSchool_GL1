package main

import (
	"fmt"
	"sync"
)

// func main() {
// 	var WG sync.WaitGroup
// 	counter := 2000
// 	WG.Add(counter)
// 	defer WG.Wait()

// 	for i := 0; i < counter; i++ {
// 		go Hello(&WG, i)
// 	}
// 	go Hello(&WG, counter)
// 	// go Hello()
// 	// time.Sleep(time.Second * 1)
// }

func Hello(wait *sync.WaitGroup, counter int) {
	fmt.Print(counter, " ")
	wait.Done()
}

// NOTES:

// A goroutine are the program that run independently and simultaneously with other goroutines

// In normal scenario, the controller will go through the hello message	function, where it will return the message and then the program will end

// While when we are using this go routine, the controller will end directly. Why, because the go routine spawns a thread that is assigned to the control of the function and since it took a while to complete, the main function will end immediately

// The completion of main will be done, before the controller other goroutines, because main only have the responsibility to create the goroutine, its work is done. So, it will terminate the control, and will kill all the other processes going on

// The sync package have some things that is knwon as the wait group, theses group keeps the main goroutine alive till all the other goroutines are done. This basically adds all the goroutines and lets them complete with having a signal sent when they are finished

// Goroutines doesn't executes in any particular order, they execute concurrently. They execute in their own pace.

// Race Conditions: When two processes are running concurrently and are trying to perform some action on same resource or variable, then that condition is knwon as the race condition.

// To avoid this race condition, we use the mutex function from the same sync package, from which we can implement lock on the particular resource or variable
