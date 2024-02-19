package day_06

import (
	"testing"

	"github.com/andrewfitzy/2022-advent-of-code/util"
)

// import "github.com/andrewfitzy/2022-advent-of-code/util"

func Test_solve02_with_demo_data_01(t *testing.T) {

	input := "mjqjpqmgbljsphdztnvjfqwrcgsmlb"

	want := 19

	got := solve02(input)

	if want != got {
		t.Errorf("solve02() = %v, want %v", got, want)
	}
}

func Test_solve02_with_demo_data_02(t *testing.T) {

	input := "bvwbjplbgvbhsrlpgdmjqwftvncz"

	want := 23

	got := solve02(input)

	if want != got {
		t.Errorf("solve02() = %v, want %v", got, want)
	}
}

func Test_solve02_with_demo_data_03(t *testing.T) {
	input := "nppdvjthqldpwncqszvftbrmjlhg"

	want := 23

	got := solve02(input)

	if want != got {
		t.Errorf("solve02() = %v, want %v", got, want)
	}
}

func Test_solve02_with_demo_data_04(t *testing.T) {
	input := "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"

	want := 29

	got := solve02(input)

	if want != got {
		t.Errorf("solve02() = %v, want %v", got, want)
	}
}

func Test_solve02_with_demo_data_05(t *testing.T) {
	input := "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"

	want := 26

	got := solve02(input)

	if want != got {
		t.Errorf("solve02() = %v, want %v", got, want)
	}
}

func Test_solve02_with_real_data_01(t *testing.T) {
	input := util.GetFileContent("puzzle_input.txt")

	want := 2508

	got := solve02(input)

	if want != got {
		t.Errorf("solve02() = %v, want %v", got, want)
	}
}
