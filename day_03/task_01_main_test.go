package day_03

import (
	"testing"
)

func Test_solve01_with_demo_data_01(t *testing.T) {
	input := "vJrwpWtwJgWrhcsFMMfFFhFp\n" +
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\n" +
		"PmmdzqPrVvPwwTWBwg\n" +
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\n" +
		"ttgJtRGJQctTZtZT\n" +
		"CrZsJsPPZsGzwwsLwLmpwMDw"

	want := 157

	got := solve01(input)

	if want != got {
		t.Errorf("solve01() = %v, want %v", got, want)
	}
}

// func Test_solve01_with_real_data_01(t *testing.T) {
// 	input := util.GetFileContent("puzzle_input.txt")

// 	want := 8109

// 	got := solve01(input)

// 	if want != got {
// 		t.Errorf("solve01() = %v, want %v", got, want)
// 	}
// }
