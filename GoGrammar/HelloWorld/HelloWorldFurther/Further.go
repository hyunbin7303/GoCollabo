package helloworldfurther

import (
	"errors"
	"fmt"
	"math/rand"
)

// In golang, it uses the world Exported and unexported rather than public or private.
// Exported function since the first character is capitalized.Exported and unexported items.
// using upper case of the first character - public, visible ooutside of package
// using lower case of the first character - private,not visible outside of package
func PrintOutFoo(num int) int {
	fmt.Printf("Printing out number:%d. ", num)
	return num
}
func checking() { // This cannot be used in hello.go or other go files.
	fmt.Print("testing")
}

func SwitchStatementMultiValues(char rune) bool {
	switch char {
	case ' ', '\t', '\n', '\f', '\r':
		return true
	}
	return false
}

// Fallthrough - transfers control to the next case.
func SwitchStatement(num int) {
	x := num
	switch {
	case x < 20:
		fmt.Println("x is less than 20")
		fallthrough

	case x >= 20 && x < 60:
		fmt.Println("x is bigger than or equal to 20, and less than 60.")
		fallthrough
	default:
		fmt.Println("Not in the criteria")
	}
}

func ValidateEvenNum() {
	fmt.Println("Type number:")
	var num int
	fmt.Scanln(&num)
	if ok, err := isEvenNumber(num); ok {
		fmt.Printf("Given number %d is divisible. %v \n", num, ok)
	} else {
		fmt.Println(err)
	}
}

func IfStatement(num int) {
	x := num
	if z := 2 * rand.Intn(x); z >= x {
		fmt.Printf("Z value checking %v\n", z)
	} else {
		fmt.Printf("Z value and x value : %v %v \n", z, x)
	}

}
func GetScoreOfUserFromMap(user string) (int, error) {
	mapping := map[string]int{
		"kevin": 80,
		"macy":  54,
		"adam":  76,
		"julio": 80,
	}
	if v, exists := mapping["kevin"]; exists {
		fmt.Printf("Kevin's score: %d \n", v)
		return v, nil
	}
	return 0, errors.New("User does not exist in the system")
}

func ForRangeGetSum(givenFiveNums [5]int) int {
	for i, v := range givenFiveNums {
		fmt.Println("Ranging testing", i, v)
	}
	var total = 0
	for _, price := range givenFiveNums {
		total += price
	}
	return total
}

func isEvenNumber(num int) (bool, error) {
	if num%2 != 0 {
		return false, errors.New("Given number is not divisible by 2.")
	}
	return true, nil
}
