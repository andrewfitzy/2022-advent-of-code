package day_07

import (
	"testing"
)

func Test_solve02_with_demo_date_01(t *testing.T) {
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

	want := 24933642

	got := solve02(input)

	if want != got {
		t.Errorf("solve02() = %v, want %v", got, want)
	}
}

// func Test_solve02_with_real_data_01(t *testing.T) {
// 	input := util.GetFileContent("puzzle_input.txt")

// 	want := 942298

// 	got := solve02(input)

// 	if want != got {
// 		t.Errorf("solve02() = %v, want %v", got, want)
// 	}
// }
