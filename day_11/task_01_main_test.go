package day_11

import (
	"testing"

	"github.com/andrewfitzy/2022-advent-of-code/util"
)

func Test_solve01_with_demo_data_01(t *testing.T) {
	input := util.GetFileContent("example_input.txt")

	want := 10605

	got := solve01(input, 20)

	if want != got {
		t.Errorf("solveExample01() = %v, want %v", got, want)
	}
}

// func Test_solve01_with_real_data_01(t *testing.T) {
// 	input := util.GetFileContent("puzzle_input.txt")

// 	want := 50830

// 	got := solve01(input, 20)

// 	if want != got {
// 		t.Errorf("solveExample01() = %v, want %v", got, want)
// 	}
// }
