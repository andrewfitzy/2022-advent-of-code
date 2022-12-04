package day_04

import "bufio"
import "strings"
import "strconv"
import "github.com/emirpasic/gods/sets/hashset"

func buildSet(start int, stop int) *hashset.Set {
	set := hashset.New() // empty

	for i := start; i <= stop; i++ {
		set.Add(i)
	}
	return set
}

func isOneOverlappingTheOther(left string, right string) bool {
	leftMidpoint := strings.Index(left, "-")
	leftStart, _ := strconv.Atoi(left[0:leftMidpoint])
	leftStop, _ := strconv.Atoi(left[leftMidpoint+1:])
	leftSet := buildSet(leftStart, leftStop)

	rightMidpoint := strings.Index(right, "-")
	rightStart, _ := strconv.Atoi(right[0:rightMidpoint])
	rightStop, _ := strconv.Atoi(right[rightMidpoint+1:])
	rightSet := buildSet(rightStart, rightStop)

	intersection := leftSet.Intersection(rightSet)

	return intersection.Size() > 0
}

func solve02(input string) int {
	total := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		midpoint := strings.Index(line, ",")
		left := line[0:midpoint]
		right := line[midpoint+1:]

		if isOneOverlappingTheOther(left, right) {
			total = total + 1
		}
	}

	return total
}
