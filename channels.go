package main

import (
	"fmt"
	"time"
)

// Things to remember
// Sending/Receiving from Go Channels are blocking calls.
func main() {
	// this is how we create a channel.
	msgChannel := make(chan string)

	// creating an anonymous function to call as a routine
	go func() {
		//passing the value from go-routine to a channel (msgChannel)
		msgChannel <- "hello from go-routine"
	}()// calling the go routine

	// blocking the main() thread until it receives the msg from the msgChannel.
	msg := <-msgChannel
	// once the channel sends out a msg, we can continue.
	fmt.Println(msg)

	// -----------------------------------------------------------------------------
	// Buffered Channels
	bufferedChannel := make(chan string, 2)

	bufferedChannel <- "message from 1st"
	bufferedChannel <- "message from 2nd"
	// if we uncomment below line we will get error
	// because channels are by-default unbuffered and for every msg sent by channel there must be a
	// equivalent receiving that msg, otherwise we get here. In the case of buffered we can send n number
	// inside the channels before taking them out.
	// bufferedChannel <- "message from 3rd"


	fmt.Println(<- bufferedChannel)
	fmt.Println(<- bufferedChannel)

	// -----------------------------------------------------------------------------
	// Channel Synchronization. Wait for one-goroutine to finsih the task
	done := make(chan string, 1)
	another := make(chan string, 1)
	worker(another)
	worker(done)
	fmt.Println("Main is doing its own thing")
	<- another
	<- done
	fmt.Println("Main is doing its own thing")
}

func worker(done chan string) {
	for i := 0; i< 10; i++ {
		fmt.Print("( working..", i, ")")
	}
	fmt.Println("\nWork is Done")
	time.Sleep(time.Second *10 )
	done <- "true"
}
