package day_11

import (
	"testing"

	"github.com/andrewfitzy/2022-advent-of-code/util"
)

func Test_solve02_with_demo_data_01(t *testing.T) {
	input := util.GetFileContent("example_input.txt")

	want := 2713310158

	got := solve02(input, 10000)

	if want != got {
		t.Errorf("solveExample01() = %v, want %v", got, want)
	}
}

// func Test_solve02_with_real_data_01(t *testing.T) {
// 	input := util.GetFileContent("puzzle_input.txt")

// 	want := 14399640002

// 	got := solve02(input, 10000)

// 	if want != got {
// 		t.Errorf("solve01() = %v, want %v", got, want)
// 	}
// }
