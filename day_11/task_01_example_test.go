package day_11

import "testing"
import "github.com/andrewfitzy/2022-advent-of-code/util"

func Test_solveExample01_with_input_01(t *testing.T) {
	input := util.GetFileContent("task_01_input_01.txt")

	want := 10605

	got := solveExample01(input)

	if want != got {
		t.Errorf("solveExample01() = %v, want %v", got, want)
	}
}
