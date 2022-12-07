package day_01

import "testing"
import "github.com/andrewfitzy/2022-advent-of-code/util"

func Test_solve01_with_input_01(t *testing.T) {
	input := util.GetFileContent("task_01_input_01.txt")

	want := elf{
		index:    3,
		calories: 24000,
	}

	got := solve01(input)

	if want != got {
		t.Errorf("solve01() = %v, want %v", got, want)
	}
}

// func Test_solve01_with_input_02(t *testing.T) {
// 	input := util.GetFileContent("task_01_input_02.txt")

// 	want := elf{
// 		index:    28,
// 		calories: 0,
// 	}

// 	got := solve01(input)

// 	if want != got {
// 		t.Errorf("solve01() = %v, want %v", got, want)
// 	}
// }
