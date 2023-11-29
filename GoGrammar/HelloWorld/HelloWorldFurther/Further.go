package helloworldfurther

import "fmt"

// Exported function since the first character is capitalized.
func Testing() {
	fmt.Print("Testing")
	testing()
}
func testing() {
	fmt.Print("testing")
}
