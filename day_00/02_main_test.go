package day_00

import "testing"

func Test_sayHello_02(t *testing.T) {
	name := "Beth"
	want := "Hello Beth"

	if got := sayHello(name); got != want {
		t.Errorf("hello() = %q, want %q", got, want)
	}
}
