package day_02

import (
	"testing"
)

func Test_solve01_with_demo_data_01(t *testing.T) {
	input := "A Y\n" +
		"B X\n" +
		"C Z"

	want := 15

	got := solve01(input)

	if want != got {
		t.Errorf("solve01() = %v, want %v", got, want)
	}
}

// func Test_solve01_with_real_data_01(t *testing.T) {
// 	input := util.GetFileContent("puzzle_input.txt")

// 	want := 13268

// 	got := solve01(input)

// 	if want != got {
// 		t.Errorf("solve01() = %v, want %v", got, want)
// 	}
// }
