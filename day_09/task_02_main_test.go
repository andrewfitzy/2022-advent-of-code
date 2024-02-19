package day_09

import (
	"testing"
)

func Test_solve02_with_demo_data_01(t *testing.T) {
	input := "R 4\n" +
		"U 4\n" +
		"L 3\n" +
		"D 1\n" +
		"R 4\n" +
		"D 1\n" +
		"L 5\n" +
		"R 2"

	want := 1

	got := solve02(input)

	if want != got {
		t.Errorf("solve02() = %v, want %v", got, want)
	}
}

func Test_solve02_with_demo_data_02(t *testing.T) {
	input := "R 5\n" +
		"U 8\n" +
		"L 8\n" +
		"D 3\n" +
		"R 17\n" +
		"D 10\n" +
		"L 25\n" +
		"U 20"

	want := 36

	got := solve02(input)

	if want != got {
		t.Errorf("solve02() = %v, want %v", got, want)
	}
}

// func Test_solve02_with_real_data_01(t *testing.T) {
// 	input := util.GetFileContent("puzzle_input.txt")

// 	want := 2273

// 	got := solve02(input)

// 	if want != got {
// 		t.Errorf("solve02() = %v, want %v", got, want)
// 	}
// }
