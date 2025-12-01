package day_12

import "fmt"

type point struct {
	col int
	row int
}

type entry struct {
	location point
	path     []point
	cost     int
}

func getLocation(itemToFind string, space [][]string) point {
	for i := 0; i < len(space); i++ {
		for j := 0; j < len(space[i]); j++ {
			if space[i][j] == itemToFind {
				return point{
					col: j,
					row: i,
				}
			}
		}
	}
	panic(fmt.Sprintf("Cannot find %v", itemToFind))
}

func isValidStep(location string, nextStep string) bool {
	// validate lengths as we'll use index access in a mo
	if len(location) != 1 {
		panic(fmt.Sprintf("Invalid content in location %v", location))
	}

	if len(nextStep) != 1 {
		panic(fmt.Sprintf("Invalid content in nextStep %v", nextStep))
	}

	// From the start, assume S is at elevation a
	var locationVal rune
	if location == "S" {
		locationVal = rune('a')
	} else {
		locationVal = rune(location[0])
	}

	// At the end, assume E is at elevation z
	var nextStepVal rune
	if nextStep == "E" {
		nextStepVal = rune('z')
	} else {
		nextStepVal = rune(nextStep[0])
	}

	// check if we're going up 1 point of elevation, staying the same, or going down n steps
	if locationVal+1 >= nextStepVal {
		return true
	}
	return false
}

// Comparator function (sort by smallest number of spaces)
func byCost(a, b interface{}) int {
	priorityA := a.(entry).cost
	priorityB := b.(entry).cost
	if priorityA > priorityB {
		return 1
	} else if priorityA < priorityB {
		return -1
	} else {
		return 0
	}
}
