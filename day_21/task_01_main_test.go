package day_12

import "testing"
import "github.com/andrewfitzy/2022-advent-of-code/util"

func Test_solve01_with_input_01(t *testing.T) {
	input := util.GetFileContent("task_01_input_01.txt")

	want := 152

	got := solve01(input, "root")

	if want != got {
		t.Errorf("solve01() = %v, want %v", got, want)
	}
}

// func Test_solve01_with_input_02(t *testing.T) {
// 	input := util.GetFileContent("task_01_input_02.txt")

// 	want := 0

// 	got := solve01(input, "root")

// 	if want != got {
// 		t.Errorf("solve01() = %v, want %v", got, want)
// 	}
// }
