package day_12

import (
	"testing"
)

func Test_solve01_with_demo_data_01(t *testing.T) {
	input := "1\n" +
		"2\n" +
		"-3\n" +
		"3\n" +
		"-2\n" +
		"0\n" +
		"4"

	want := 3

	got := solve01(input)

	if want != got {
		t.Errorf("solve01() = %v, want %v", got, want)
	}
}

// func Test_solve01_with_real_data_01(t *testing.T) {
// 	input := util.GetFileContent("puzzle_input.txt")

// 	want := 0

// 	got := solve01(input)

// 	if want != got {
// 		t.Errorf("solve01() = %v, want %v", got, want)
// 	}
// }
