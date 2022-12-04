package day_04

import "testing"

func Test_solve01_with_input_01(t *testing.T) {
	input := get_file_content("task_01_input_01.txt")

	want := 2

	got := solve01(input)

	if want != got {
		t.Errorf("solve01() = %v, want %v", got, want)
	}
}

func Test_solve01_with_input_02(t *testing.T) {
	input := get_file_content("task_01_input_02.txt")

	want := 513

	got := solve01(input)

	if want != got {
		t.Errorf("solve01() = %v, want %v", got, want)
	}
}
