package day_01

import (
	"testing"
)

func Test_solve02_with_demo_data_01(t *testing.T) {
	input := "1000\n" +
		"2000\n" +
		"3000\n" +
		"\n" +
		"4000\n" +
		"\n" +
		"5000\n" +
		"6000\n" +
		"\n" +
		"7000\n" +
		"8000\n" +
		"9000\n" +
		"\n" +
		"10000"

	want := 45000

	got := solve02(input)

	if want != got {
		t.Errorf("solve02() = %v, want %v", got, want)
	}
}

// func Test_solve02_with_real_data_01(t *testing.T) {
// 	input := util.GetFileContent("puzzle_input.txt")

// 	want := 209691

// 	got := solve02(input)

// 	if want != got {
// 		t.Errorf("solve02() = %v, want %v", got, want)
// 	}
// }
