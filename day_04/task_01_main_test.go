package day_04

import (
	"testing"
)

func Test_solve01_with_demo_data_01(t *testing.T) {
	input := "2-4,6-8\n" +
		"2-3,4-5\n" +
		"5-7,7-9\n" +
		"2-8,3-7\n" +
		"6-6,4-6\n" +
		"2-6,4-8"

	want := 2

	got := solve01(input)

	if want != got {
		t.Errorf("solve01() = %v, want %v", got, want)
	}
}

// func Test_solve01_with_real_data_01(t *testing.T) {
// 	input := util.GetFileContent("puzzle_input.txt")

// 	want := 513

// 	got := solve01(input)

// 	if want != got {
// 		t.Errorf("solve01() = %v, want %v", got, want)
// 	}
// }
