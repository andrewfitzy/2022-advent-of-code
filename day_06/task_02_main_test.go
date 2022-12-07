package day_06

import "testing"

// import "github.com/andrewfitzy/2022-advent-of-code/util"

func Test_solve02_with_input_01(t *testing.T) {

	input := "mjqjpqmgbljsphdztnvjfqwrcgsmlb"

	want := 19

	got := solve02(input)

	if want != got {
		t.Errorf("solve02() = %v, want %v", got, want)
	}
}

func Test_solve02_with_input_02(t *testing.T) {

	input := "bvwbjplbgvbhsrlpgdmjqwftvncz"

	want := 23

	got := solve02(input)

	if want != got {
		t.Errorf("solve02() = %v, want %v", got, want)
	}
}

func Test_solve02_with_input_03(t *testing.T) {
	input := "nppdvjthqldpwncqszvftbrmjlhg"

	want := 23

	got := solve02(input)

	if want != got {
		t.Errorf("solve02() = %v, want %v", got, want)
	}
}

func Test_solve02_with_input_04(t *testing.T) {
	input := "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"

	want := 29

	got := solve02(input)

	if want != got {
		t.Errorf("solve02() = %v, want %v", got, want)
	}
}

func Test_solve02_with_input_05(t *testing.T) {
	input := "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"

	want := 26

	got := solve02(input)

	if want != got {
		t.Errorf("solve02() = %v, want %v", got, want)
	}
}

// func Test_solve02_with_input_06(t *testing.T) {
// 	input := util.GetFileContent("task_01_input_01.txt")

// 	want := 0

// 	got := solve02(input)

// 	if want != got {
// 		t.Errorf("solve02() = %v, want %v", got, want)
// 	}
// }
