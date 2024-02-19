package day_05

import (
	"testing"
)

func Test_solve02_with_demo_data_01(t *testing.T) {
	/*
	       [D]
	   [N] [C]
	   [Z] [M] [P]
	    1   2   3
	*/
	input := "move 1 from 2 to 1\n" +
		"move 3 from 1 to 3\n" +
		"move 2 from 2 to 1\n" +
		"move 1 from 1 to 2"

	startPosition := map[int]string{
		1: "ZN",
		2: "MCD",
		3: "P",
	}

	want := "MCD"

	got := solve02(startPosition, input)

	if want != got {
		t.Errorf("solve02() = %v, want %v", got, want)
	}
}

// func Test_solve02_with_real_data_01(t *testing.T) {
// 	/*
// 	                   [B]     [L]     [S]
// 	           [Q] [J] [C]     [W]     [F]
// 	       [F] [T] [B] [D]     [P]     [P]
// 	       [S] [J] [Z] [T]     [B] [C] [H]
// 	       [L] [H] [H] [Z] [G] [Z] [G] [R]
// 	   [R] [H] [D] [R] [F] [C] [V] [Q] [T]
// 	   [C] [J] [M] [G] [P] [H] [N] [J] [D]
// 	   [H] [B] [R] [S] [R] [T] [S] [R] [L]
// 	    1   2   3   4   5   6   7   8   9
// 	*/

// 	input := util.GetFileContent("puzzle_input.txt")
// 	startPosition := map[int]string{
// 		1: "HCR",
// 		2: "BJHLSF",
// 		3: "RMDHJTQ",
// 		4: "SGRHZBJ",
// 		5: "RPFZTDCB",
// 		6: "THCG",
// 		7: "SNVZBPWL",
// 		8: "RJQGC",
// 		9: "LDTRHPFS",
// 	}

// 	want := "CDTQZHBRS"

// 	got := solve02(startPosition, input)

// 	if want != got {
// 		t.Errorf("solve02() = %v, want %v", got, want)
// 	}
// }
