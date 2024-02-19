package day_10

import (
	"testing"

	"github.com/andrewfitzy/2022-advent-of-code/util"
)

func Test_solve02_with_demo_data_01(t *testing.T) {
	input := util.GetFileContent("example_input.txt")
	want := "##..##..##..##..##..##..##..##..##..##..\n" +
		"###...###...###...###...###...###...###.\n" +
		"####....####....####....####....####....\n" +
		"#####.....#####.....#####.....#####.....\n" +
		"######......######......######......####\n" +
		"#######.......#######.......#######.....\n"

	got := solve02(input)

	if want != got {
		t.Errorf("solve02() = \n%v, want \n%v", got, want)
	}
}

// func Test_solve02_with_real_data_01(t *testing.T) {
// 	input := util.GetFileContent("puzzle_input.txt")

// 	// If all rows are aligned, the letters are FPGPHFGH
// 	want := "####.###...##..###..#..#.####..##..#..#.\n" +
// 		"#....#..#.#..#.#..#.#..#.#....#..#.#..#.\n" +
// 		"###..#..#.#....#..#.####.###..#....####.\n" +
// 		"#....###..#.##.###..#..#.#....#.##.#..#.\n" +
// 		"#....#....#..#.#....#..#.#....#..#.#..#.\n" +
// 		"#....#.....###.#....#..#.#.....###.#..#.\n"

// 	got := solve02(input)

// 	if want != got {
// 		t.Errorf("solve02() = \n%v, want \n%v", got, want)
// 	}
// }
