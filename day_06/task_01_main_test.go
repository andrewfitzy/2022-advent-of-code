package day_06

import (
	"testing"
)

// import "github.com/andrewfitzy/2022-advent-of-code/util"

func Test_solve01_with_demo_data_01(t *testing.T) {

	input := "mjqjpqmgbljsphdztnvjfqwrcgsmlb"

	want := 7

	got := solve01(input)

	if want != got {
		t.Errorf("solve01() = %v, want %v", got, want)
	}
}

func Test_solve01_with_demo_data_02(t *testing.T) {

	input := "bvwbjplbgvbhsrlpgdmjqwftvncz"

	want := 5

	got := solve01(input)

	if want != got {
		t.Errorf("solve01() = %v, want %v", got, want)
	}
}

func Test_solve01_with_demo_data_03(t *testing.T) {
	input := "nppdvjthqldpwncqszvftbrmjlhg"

	want := 6

	got := solve01(input)

	if want != got {
		t.Errorf("solve01() = %v, want %v", got, want)
	}
}

func Test_solve01_with_demo_data_04(t *testing.T) {
	input := "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"

	want := 10

	got := solve01(input)

	if want != got {
		t.Errorf("solve01() = %v, want %v", got, want)
	}
}

func Test_solve01_with_demo_data_05(t *testing.T) {
	input := "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"

	want := 11

	got := solve01(input)

	if want != got {
		t.Errorf("solve01() = %v, want %v", got, want)
	}
}

// func Test_solve01_with_real_data_01(t *testing.T) {
// 	input := util.GetFileContent("puzzle_input.txt")

// 	want := 1804

// 	got := solve01(input)

// 	if want != got {
// 		t.Errorf("solve01() = %v, want %v", got, want)
// 	}
// }
