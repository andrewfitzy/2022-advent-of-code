package task_02

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/sets/hashset"
)

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

func getNextMoves(location point, landscape [][]string) []point {
	rows := len(landscape)
	cols := len(landscape[0])

	var validMoves []point

	// Movement Deltas (Up, Down, Left, Right)
	movements := [4][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

	for i := 0; i < len(movements); i++ {
		next_x := location.col + movements[i][0]
		next_y := location.row + movements[i][1]
		is_in_bounds := next_y >= 0 && next_y < rows && next_x >= 0 && next_x < cols

		// flip next and current so that the logic is still correct even though we are going back from the end
		if is_in_bounds && isValidStep(landscape[next_y][next_x], landscape[location.row][location.col]) {
			nextPoint := point{
				col: next_x,
				row: next_y,
			}
			validMoves = append(validMoves, nextPoint)
		}
	}
	return validMoves
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

func getShortestPath(start point, space [][]string) []point {
	queue := priorityqueue.NewWith(byCost)
	visited := hashset.New()

	first_step := entry{
		location: start,
		path:     []point{start},
		cost:     1,
	}
	queue.Enqueue(first_step)

	for !queue.Empty() {
		currentState, _ := queue.Dequeue()
		// First check that we haven't already been to the place
		if visited.Contains(currentState.(entry).location) {
			continue
		}
		visited.Add(currentState.(entry).location)

		// Second check if we've found an 'a'
		if space[currentState.(entry).location.row][currentState.(entry).location.col] == "a" {
			return currentState.(entry).path
		}

		// Finally, add valid next locations
		nextLocations := getNextMoves(currentState.(entry).location, space)
		for _, val := range nextLocations {
			newPath := append(currentState.(entry).path, val)
			nextStep := entry{
				location: val,
				path:     newPath,
				cost:     currentState.(entry).cost + 1,
			}
			queue.Enqueue(nextStep)
		}
	}

	return []point{}
}

func solve02(input string) int {
	landscape := make([][]string, 0)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		hills := strings.Split(line, "")
		landscape = append(landscape, hills)
	}

	// first find E
	eLocation := getLocation("E", landscape)

	shortestPath := getShortestPath(eLocation, landscape)

	// return transitions not steps visited.
	return len(shortestPath) - 1
}
