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

	switch helloworldfurther.PrintOutFoo(20) {
	case helloworldfurther.PrintOutFoo(10), helloworldfurther.PrintOutFoo(20), helloworldfurther.PrintOutFoo(30):
		fmt.Println("First number say it.")
		fallthrough

	case helloworldfurther.PrintOutFoo(40):
		fmt.Println("Print out the second number.")
	}

	fmt.Println(puppy.Bark())
	fmt.Println(dog.WhenGrownUp("Test"))

	helloworldfurther.IfStatement(40)
	helloworldfurther.SwitchStatement(30)
	helloworldfurther.GoRoutineUsage()
	// helloworldfurther.ForRangeLoopStatement()

	helloworldfurther.UsingOkIdiomWithInterface()

}

func PrintMe() {
	fmt.Println(x)
}
