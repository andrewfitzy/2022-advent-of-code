package day_03

import "testing"

func Test_solve02_with_input_01(t *testing.T) {
	input := get_file_content("task_01_input_01.txt")

	want := 70

	got := solve02(input)

	if want != got {
		t.Errorf("solve02() = %v, want %v", got, want)
	}
}

func Test_solve02_with_input_02(t *testing.T) {
	input := get_file_content("task_01_input_02.txt")

	want := 2738

	got := solve02(input)

	if want != got {
		t.Errorf("solve02() = %v, want %v", got, want)
	}
}
