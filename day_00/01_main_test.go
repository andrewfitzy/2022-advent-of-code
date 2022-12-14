package day_00

import "testing"

func Test_sayHello_01(t *testing.T) {
	name := "Bob"
	want := "Hello Bob"

	if got := sayHello(name); got != want {
		t.Errorf("hello() = %q, want %q", got, want)
	}
}
