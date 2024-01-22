package helloworldfurther

import "fmt"

type HelloInterface interface {
	String() string
}

// Using Ok Idiom
func UsingOkIdiomWithInterface() {
	var str interface{} = "testing sentence"
	s, ok := str.(string)
	fmt.Println(s, ok)
}
