package day_12

import "bufio"
import "strings"
import "fmt"

func solve01(input string) int {
	result := 31
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		// first find S
		// then find E

		// do something like
		// keep a stack of points or hashset of points or something
		// from current position, find next that is lowest value, then next then next
		// for each lowest put the co-ord in a hash set
		//cant go to a square we've already been to, nemsh the set of visited?
		// a move is one where current is >= next || current + 1 == next

	}
	return result
}
