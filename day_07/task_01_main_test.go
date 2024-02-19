package day_07

import (
	"testing"
)

func Test_solve01_with_demo_data_01(t *testing.T) {
	input := "$ cd /\n" +
		"$ ls\n" +
		"dir a\n" +
		"14848514 b.txt\n" +
		"8504156 c.dat\n" +
		"dir d\n" +
		"$ cd a\n" +
		"$ ls\n" +
		"dir e\n" +
		"29116 f\n" +
		"2557 g\n" +
		"62596 h.lst\n" +
		"$ cd e\n" +
		"$ ls\n" +
		"584 i\n" +
		"$ cd ..\n" +
		"$ cd ..\n" +
		"$ cd d\n" +
		"$ ls\n" +
		"4060174 j\n" +
		"8033020 d.log\n" +
		"5626152 d.ext\n" +
		"7214296 k"

	want := 95437

	got := solve01(input)

	if want != got {
		t.Errorf("solve01() = %v, want %v", got, want)
	}
}

// func Test_solve01_with_real_data_01(t *testing.T) {
// 	input := util.GetFileContent("puzzle_input.txt")

// 	want := 1443806

// 	got := solve01(input)

// 	if want != got {
// 		t.Errorf("solve01() = %v, want %v", got, want)
// 	}
// }
