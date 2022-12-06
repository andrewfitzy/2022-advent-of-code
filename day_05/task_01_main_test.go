package day_05

import "testing"
import "github.com/andrewfitzy/2022-advent-of-code/util"

func Test_solve01_with_input_01(t *testing.T) {
	/*
	       [D]
	   [N] [C]
	   [Z] [M] [P]
	    1   2   3
	*/
	input := util.GetFileContent("task_01_input_01.txt")

	startPosition := map[int]string{
		1: "ZN",
		2: "MCD",
		3: "P",
	}

	want := "CMZ"

	got := solve01(startPosition, input)

	if want != got {
		t.Errorf("solve01() = %v, want %v", got, want)
	}
}

func Test_solve01_with_input_02(t *testing.T) {
	/*
	                   [B]     [L]     [S]
	           [Q] [J] [C]     [W]     [F]
	       [F] [T] [B] [D]     [P]     [P]
	       [S] [J] [Z] [T]     [B] [C] [H]
	       [L] [H] [H] [Z] [G] [Z] [G] [R]
	   [R] [H] [D] [R] [F] [C] [V] [Q] [T]
	   [C] [J] [M] [G] [P] [H] [N] [J] [D]
	   [H] [B] [R] [S] [R] [T] [S] [R] [L]
	    1   2   3   4   5   6   7   8   9
	*/

	input := util.GetFileContent("task_01_input_02.txt")
	startPosition := map[int]string{
		1: "HCR",
		2: "BJHLSF",
		3: "RMDHJTQ",
		4: "SGRHZBJ",
		5: "RPFZTDCB",
		6: "THCG",
		7: "SNVZBPWL",
		8: "RJQGC",
		9: "LDTRHPFS",
	}

	want := "SHQWSRBDL"

	got := solve01(startPosition, input)

	if want != got {
		t.Errorf("solve01() = %v, want %v", got, want)
	}
}
