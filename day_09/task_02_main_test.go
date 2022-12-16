package day_09

import "testing"
import "github.com/andrewfitzy/2022-advent-of-code/util"

func Test_solve02_with_input_01(t *testing.T) {
	input := util.GetFileContent("task_01_input_01.txt")

	want := 1

	got := solve02(input)

	if want != got {
		t.Errorf("solve02() = %v, want %v", got, want)
	}
}

func Test_solve02_with_input_02(t *testing.T) {
	input := util.GetFileContent("task_02_input_02.txt")

	want := 36

	got := solve02(input)

	if want != got {
		t.Errorf("solve02() = %v, want %v", got, want)
	}
}

// func Test_solve02_with_input_03(t *testing.T) {
// 	input := util.GetFileContent("task_01_input_02.txt")

// 	want := 0

// 	got := solve02(input)

// 	if want != got {
// 		t.Errorf("solve02() = %v, want %v", got, want)
// 	}
// }
