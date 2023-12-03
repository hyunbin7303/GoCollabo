package main

import (
	helloworldfurther "HelloWorld/HelloWorld/HelloWorldFurther"
	"fmt"

	"github.com/GoesToEleven/dog"
	"github.com/GoesToEleven/puppy"
)

// file block scope
var x = 20

func main() {
	fmt.Println("Hello, World!")
	PrintMe()
	x := 30
	fmt.Println(x)
	PrintMe()

	helloworldfurther.Testing()
	// helloworldfurther.Checking()
	fmt.Println(puppy.Bark())
	fmt.Println(dog.WhenGrownUp("Test"))

	helloworldfurther.IfStatement()

	helloworldfurther.SwitchStatement()

	helloworldfurther.GoSelectStatement()

	helloworldfurther.ForRangeLoopStatement()
}

func PrintMe() {
	fmt.Println(x)
}
