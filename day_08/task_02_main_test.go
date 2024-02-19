package day_08

import (
	"testing"
)

func Test_solve02_with_demo_data_01(t *testing.T) {
	input := "30373\n" +
		"25512\n" +
		"65332\n" +
		"33549\n" +
		"35390"

	want := 8

	got := solve02(input)

	if want != got {
		t.Errorf("solve02() = %v, want %v", got, want)
	}
}

// func Test_solve02_with_real_data_01(t *testing.T) {
// 	input := util.GetFileContent("puzzle_input.txt")

// 	want := 330786

// 	got := solve02(input)

// 	if want != got {
// 		t.Errorf("solve02() = %v, want %v", got, want)
// 	}
// }
