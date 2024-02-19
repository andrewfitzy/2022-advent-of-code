package day_03

import (
	"testing"
)

func Test_solve02_with_demo_data_01(t *testing.T) {
	input := "vJrwpWtwJgWrhcsFMMfFFhFp\n" +
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\n" +
		"PmmdzqPrVvPwwTWBwg\n" +
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\n" +
		"ttgJtRGJQctTZtZT\n" +
		"CrZsJsPPZsGzwwsLwLmpwMDw"

	want := 70

	got := solve02(input)

	if want != got {
		t.Errorf("solve02() = %v, want %v", got, want)
	}
}

// func Test_solve02_with_real_data_01(t *testing.T) {
// 	input := util.GetFileContent("puzzle_input.txt")

// 	want := 2738

// 	got := solve02(input)

// 	if want != got {
// 		t.Errorf("solve02() = %v, want %v", got, want)
// 	}
// }
