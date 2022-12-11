package day_10

import "testing"
import "github.com/andrewfitzy/2022-advent-of-code/util"

func Test_solve02_with_input_01(t *testing.T) {
	input := util.GetFileContent("task_01_input_01.txt")
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

func Test_solve02_with_input_02(t *testing.T) {
	input := util.GetFileContent("task_01_input_02.txt")

	// leaving this one in, all answers are different for this challenge
	want := "####.###...##..###..#..#.####..##..#..#.\n" +
		"#....#..#.#..#.#..#.#..#.#....#..#.#..#.\n" +
		"###..#..#.#....#..#.####.###..#....####.\n" +
		"#....###..#.##.###..#..#.#....#.##.#..#.\n" +
		"#....#....#..#.#....#..#.#....#..#.#..#.\n" +
		"#....#.....###.#....#..#.#.....###.#..#.\n"

	got := solve02(input)

	if want != got {
		t.Errorf("solve02() = \n%v, want \n%v", got, want)
	}
}
