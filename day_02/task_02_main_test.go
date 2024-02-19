package day_02

import (
	"testing"
)

func Test_solve02_with_demo_data_01(t *testing.T) {
	input := "A Y\n" +
		"B X\n" +
		"C Z"

	want := 12

	got := solve02(input)

	if want != got {
		t.Errorf("solve02() = %v, want %v", got, want)
	}
}

// func Test_solve02_real_data_01(t *testing.T) {
// 	input := util.GetFileContent("puzzle_input.txt")

// 	want := 15508

// 	got := solve02(input)

// 	if want != got {
// 		t.Errorf("solve02() = %v, want %v", got, want)
// 	}
// }
