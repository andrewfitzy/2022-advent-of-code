package day_04

import "bufio"
import "strings"
import "strconv"

func isOneContainedInTheOther(left string, right string) bool {
	leftMidpoint := strings.Index(left, "-")
	leftStart, _ := strconv.Atoi(left[0:leftMidpoint])
	leftStop, _ := strconv.Atoi(left[leftMidpoint+1:])

	rightMidpoint := strings.Index(right, "-")
	rightStart, _ := strconv.Atoi(right[0:rightMidpoint])
	rightStop, _ := strconv.Atoi(right[rightMidpoint+1:])

	if leftStart <= rightStart && rightStop <= leftStop {
		return true
	} else if rightStart <= leftStart && leftStop <= rightStop {
		return true
	}

	return false
}

func solve01(input string) int {
	total := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		midpoint := strings.Index(line, ",")
		left := line[0:midpoint]
		right := line[midpoint+1:]

		if isOneContainedInTheOther(left, right) {
			total = total + 1
		}
	}

	return total
}
