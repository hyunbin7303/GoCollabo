package helloworldfurther

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// In golang, it uses the world Exported and unexported rather than public or private.
// Exported function since the first character is capitalized.
// Exported and unexported items.
// using upper case of the first character - public, visible ooutside of package
// using lower case of the first character - private,not visible outside of package

func Testing() {
	fmt.Print("Testing")
	testing()
}
func testing() {
	fmt.Print("testing")
}

func SwitchStatement() {
	x := 30
	switch {
	case x < 42:
		fmt.Println("x is less than 42")
		fallthrough
	default:

	}
}

func IfStatement() {
	fmt.Println("If statement testing-----")
	x := 40
	if z := 2 * rand.Intn(x); z >= x {
		fmt.Printf("Z value checking %v\n", z)
	} else {
		fmt.Printf("Z value and x value : %v %v \n", z, x)
	}

	mapping := map[string]int{
		"kevin": 80,
		"macy":  54,
		"adam":  76,
		"julio": 80,
	}
	if v, exists := mapping["kevin"]; exists {
		fmt.Printf("Kevin's score: %d \n", v)
	}

	// Using Ok Idiom
	var str interface{} = "testing sentence"
	s, ok := str.(string)
	fmt.Println(s, ok)

	fmt.Println("Testing number")
	num := 11
	if ok, err := isEvenNumber(num); ok {
		fmt.Printf("Given number %d is divisible. %v \n", num, ok)
	} else {
		fmt.Println(err)
	}
}

func GoSelectStatement() {
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

func ForRangeLoopStatement() {
	arr := []int{10, 20, 30, 40, 50}
	for i, v := range arr {
		fmt.Println("Ranging testing", i, v)
	}
}

// If statement with interface variable.

func isEvenNumber(num int) (bool, error) {
	if num%2 != 0 {
		return false, errors.New("Given number is not divisible by 2.")
	}
	return true, nil
}
