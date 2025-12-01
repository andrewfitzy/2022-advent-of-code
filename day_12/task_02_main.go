package day_12

import (
	"bufio"
	"strings"

	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/sets/hashset"
)

// this assumes we're going from start to finish so we pass next to currenct in isValidStep
func getNextMovesBackwards(location point, landscape [][]string) []point {
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

func getShortestPathToAFromEnd(start point, space [][]string) []point {
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
		nextLocations := getNextMovesBackwards(currentState.(entry).location, space)
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

	shortestPath := getShortestPathToAFromEnd(eLocation, landscape)

	// return transitions not steps visited.
	return len(shortestPath) - 1
}
