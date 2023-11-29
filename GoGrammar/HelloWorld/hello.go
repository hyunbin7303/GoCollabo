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
	// helloworldfurther.testing()

	s1 := puppy.Bark()
	fmt.Println(s1)

	s2 := dog.WhenGrownUp("Test")

	fmt.Println(s2)
}

func PrintMe() {
	fmt.Println(x)
}

// Struct iteration - using range keyword.
//

//Package Visibility and Scope
// using upper case of the first character - public, visible ooutside of package
// using lower case of the first character - private,not visible outside of package
// Exported and unexported items.
// In golang, it uses the world Exported and unexported rather than public or private.
