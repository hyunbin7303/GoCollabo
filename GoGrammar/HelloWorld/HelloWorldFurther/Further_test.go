package helloworldfurther

import "testing"

func TestHelloName(t *testing.T) {
	input := 20

	SwitchStatement(input)
}
func TestForRangeGetSum(t *testing.T) {
	input := [5]int{10, 20, 30, 40, 50}

	got := ForRangeGetSum(input)
	want := 150

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
func TestGetScoreOfUserFromMap_WhenExistingUser_ReturnValid(t *testing.T) {
	input := "kevin"

	got, _ := GetScoreOfUserFromMap(input)
	want := 80

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestGetScoreOfUserFromMap_WhenNotExistingUser_ErrorReturn(t *testing.T) {
	input := "wronguser"

	got, _ := GetScoreOfUserFromMap(input)
	want := 80

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
