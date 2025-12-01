package day_12

import (
	"testing"

	"github.com/andrewfitzy/2022-advent-of-code/util"
)

func Test_solve02_with_demo_data_01(t *testing.T) {
	input := "Sabqponm\n" +
		"abcryxxl\n" +
		"accszExk\n" +
		"acctuvwj\n" +
		"abdefghi"

	want := 29

	got := solve02(input)

	if want != got {
		t.Errorf("solve01() = %v, want %v", got, want)
	}
}

func Test_solve02_with_real_data_01(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping day_12.Test_solve01_with_real_data_01 in short mode.")
	}
	input := util.GetFileContent("puzzle_input.txt")

	want := 500

	got := solve02(input)

	if want != got {
		t.Errorf("solve01() = %v, want %v", got, want)
	}
}
