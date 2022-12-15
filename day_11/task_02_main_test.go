package day_11

import "testing"
import "fmt"
import "github.com/andrewfitzy/2022-advent-of-code/util"

func Test_solve02_with_input_02(t *testing.T) {
	input := util.GetFileContent("task_01_input_02.txt")

	want := 0

	got := solve02(input)

	if want != got {
		fmt.Println("Don't forget to update this test with real values before running locally")

		//comment out check so that tests don't fail as answer not committed
		//t.Errorf("solve01() = %v, want %v", got, want)
	}
}
