package day_01

import "testing"
import "os"

func get_file_content(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func Test_solve_with_input_01(t *testing.T) {
	input := get_file_content("task_01_input_01.txt")

	want := elf{
		index:    3,
		calories: 24000,
	}

	got := solve(input)

	if want != got {
		t.Errorf("solve() = %q, want %q", got, want)
	}
}

func Test_solve_with_input_02(t *testing.T) {
	input := get_file_content("task_01_input_02.txt")

	want := elf{
		index:    28,
		calories: 71300,
	}

	got := solve(input)

	if want != got {
		t.Errorf("solve() = %q, want %q", got, want)
	}
}
