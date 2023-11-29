package main

import (
	helloworldfurther "HelloWorld/HelloWorld/HelloWorldFurther"
	"fmt"
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
