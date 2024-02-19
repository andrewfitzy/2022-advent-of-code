package day_09

import (
	"testing"
)

func Test_solve01_with_demo_data_01(t *testing.T) {
	input := "R 4\n" +
		"U 4\n" +
		"L 3\n" +
		"D 1\n" +
		"R 4\n" +
		"D 1\n" +
		"L 5\n" +
		"R 2"

	want := 13

	got := solve01(input)

	if want != got {
		t.Errorf("solve01() = %v, want %v", got, want)
	}
}

// func Test_solve01_with_real_data_01(t *testing.T) {
// 	input := util.GetFileContent("puzzle_input.txt")

// 	want := 6026

// 	got := solve01(input)

// 	if want != got {
// 		t.Errorf("solve01() = %v, want %v", got, want)
// 	}
// }
