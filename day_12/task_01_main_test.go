package day_12

import (
	"testing"
)

func Test_solve01_with_demo_data_01(t *testing.T) {
	input := "Sabqponm\n" +
		"abcryxxl\n" +
		"accszExk\n" +
		"acctuvwj\n" +
		"abdefghi"

	want := 31

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
