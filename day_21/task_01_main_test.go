package day_21

import (
	"testing"
)

func Test_solve01_with_demo_data_01(t *testing.T) {
	input := "root: pppw + sjmn\n" +
		"dbpl: 5\n" +
		"cczh: sllz + lgvd\n" +
		"zczc: 2\n" +
		"ptdq: humn - dvpt\n" +
		"dvpt: 3\n" +
		"lfqf: 4\n" +
		"humn: 5\n" +
		"ljgn: 2\n" +
		"sjmn: drzm * dbpl\n" +
		"sllz: 4\n" +
		"pppw: cczh / lfqf\n" +
		"lgvd: ljgn * ptdq\n" +
		"drzm: hmdt - zczc\n" +
		"hmdt: 32"

	want := 152

	got := solve01(input, "root")

	if want != got {
		t.Errorf("solve01() = %v, want %v", got, want)
	}
}

// func Test_solve01_with_real_data_01(t *testing.T) {
// 	input := util.GetFileContent("puzzle_input.txt")

// 	want := 160274622817992

// 	got := solve01(input, "root")

// 	if want != got {
// 		t.Errorf("solve01() = %v, want %v", got, want)
// 	}
// }
