package day_00

import "testing"
import "github.com/andrewfitzy/2022-advent-of-code/util"

func Test_sayHello_02(t *testing.T) {
	name := util.GetFileContent("main_input_01.txt")
	want := "Hello Beth"

	if got := sayHello(name); got != want {
		t.Errorf("hello() = %q, want %q", got, want)
	}
}
