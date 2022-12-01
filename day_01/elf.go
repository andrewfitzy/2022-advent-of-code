package day_01

import "fmt"

type elf struct {
	index    int
	calories int
}

func (a elf) String() string {
	return fmt.Sprintf("{ index: %d, calories: %d}", a.index, a.calories)
}
