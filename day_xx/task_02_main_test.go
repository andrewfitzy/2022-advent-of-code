package day_xx

import (
	"testing"

	"github.com/andrewfitzy/2022-advent-of-code/util"
)

func Test_solve02_with_demo_data_01(t *testing.T) {
	input := "Hola Mundo"

	want := "Hola Mundo"

	got := solve02(input)

	if want != got {
		t.Errorf("solve02() = %v, want %v", got, want)
	}
}

func Test_solve02_with_real_data_01(t *testing.T) {
	input := util.GetFileContent("puzzle_input.txt")

	want := "Hello World"

	got := solve02(input)

	if want != got {
		t.Errorf("solve02() = %v, want %v", got, want)
	}
}
