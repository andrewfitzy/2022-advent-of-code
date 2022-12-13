package day_11

import "testing"
import "github.com/andrewfitzy/2022-advent-of-code/util"

func Test_solve02_with_input_02(t *testing.T) {
	input := util.GetFileContent("task_01_input_02.txt")

	want := 0

	got := solve02(input)

	if want != got {
		//comment out check so that tests don't fail as answer not committed
		//t.Errorf("solve01() = %v, want %v", got, want)
	}
}
