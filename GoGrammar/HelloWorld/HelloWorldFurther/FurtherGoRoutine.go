package helloworldfurther

import (
	"fmt"
	"math/rand"
	"time"
)

// Golang Concurrency
// Goroutines and Channels

// Allow concurrent execution of functions

// Channels-  used for communication and synchronization between goroutines.
//  They can be created using make and can be used to send and receive values.

// To start a goroutine, simply prefix a function call with the go keyword. This will create a new goroutine that executes the function concurrently. Channels can be used to share data between goroutines, allowing them to communicate and synchronize.
// By using goroutines and channels, Go provides a simple and efficient way to implement concurrency in programs.

func GoRoutineUsage() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	d1 := time.Duration(rand.Int63n(250))
	d2 := time.Duration(rand.Int63n(250))
	go func() {
		time.Sleep(d1 * time.Millisecond)
		ch1 <- 41
	}()

	go func() {
		time.Sleep(d2 * time.Millisecond)
		ch2 <- 42
	}()

	// Select is for when we're writing concurrent code and dealing with channels.
	select {
	case v1 := <-ch1:
		fmt.Println("value from channel 1", v1)
	case v2 := <-ch2:
		fmt.Println("value from channel 2", v2)
	}
}
