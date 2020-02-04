package main

import (
	"fmt"
	"time"
)

func parallelFunction(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// normal call
	parallelFunction("Direct-Call")
	// call function as a go routine
	go parallelFunction("Go-Routine")
	// anonymous function and run as go-routine
	go func(from string) {
		for i := 0; i < 3; i++ {
			fmt.Println(from, ":", i)
		}
	}("Anonymous-Routine")

	// order of these two routines i.e. which one will be executed first is not guranteed..
	// even though we created the "Go-Routine" first there is still a equal chance that
	// the "Anonymous-Routine" might first start and finish.

	// If we remove the Sleep, the go-routines will never be fired.
	// because once the main() is finsihed all the goroutines are dropped from memory
	time.Sleep(time.Second)
	// we can use channels to stop that from happening

	fmt.Println("From Main")
}
