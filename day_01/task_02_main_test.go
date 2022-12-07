package day_01

import "testing"
import "github.com/andrewfitzy/2022-advent-of-code/util"

func Test_solve02_with_input_01(t *testing.T) {
	input := util.GetFileContent("task_01_input_01.txt")

	want := 45000

	got := solve02(input)

	if want != got {
		t.Errorf("solve02() = %v, want %v", got, want)
	}
}

// func Test_solve02_with_input_02(t *testing.T) {
// 	input := util.GetFileContent("task_01_input_02.txt")

// 	want := 0

// 	got := solve02(input)

// 	if want != got {
// 		t.Errorf("solve02() = %v, want %v", got, want)
// 	}
// }
