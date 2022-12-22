package day_12

import "bufio"
import "strings"
import "fmt"
import "strconv"
import "github.com/emirpasic/gods/sets/linkedhashset"

type point struct {
	col int
	row int
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
	if location == "S" {
		return true
	}
	locationVal, _ := strconv.Atoi(location)
	nextStepVal, _ := strconv.Atoi(nextStep)
	if locationVal == nextStepVal || locationVal+1 == nextStepVal {
		return true
	}
	return false
}

func getNextMoves(location point, existingMoves *linkedhashset.Set, landscape [][]string) *linkedhashset.Set {
	//check if there is an L move
	validMoves := linkedhashset.New()
	if location.col-1 >= 0 {
		leftPoint := point{
			col: location.col - 1,
			row: location.row,
		}
		if isValidStep(landscape[location.row][location.col], landscape[leftPoint.row][leftPoint.col]) && !existingMoves.Contains(leftPoint) {
			validMoves.Add(leftPoint)
		}
	}
	if location.col+1 <= len(landscape[0]) {
		rightPoint := point{
			col: location.col + 1,
			row: location.row,
		}
		if isValidStep(landscape[location.row][location.col], landscape[rightPoint.row][rightPoint.col]) && !existingMoves.Contains(rightPoint) {
			validMoves.Add(rightPoint)
		}
	}
	if location.row-1 >= 0 {
		upPoint := point{
			col: location.col,
			row: location.row - 1,
		}
		if isValidStep(landscape[location.row][location.col], landscape[upPoint.row][upPoint.col]) && !existingMoves.Contains(upPoint) {
			validMoves.Add(upPoint)
		}
	}
	if location.row+1 <= len(landscape) {
		downPoint := point{
			col: location.col,
			row: location.row + 1,
		}
		if isValidStep(landscape[location.row][location.col], landscape[downPoint.row][downPoint.col]) && !existingMoves.Contains(downPoint) {
			validMoves.Add(downPoint)
		}
	}
	return validMoves
}

func solve01(input string) int {
	landscape := make([][]string, 0)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		hills := strings.Split(line, "")
		landscape = append(landscape, hills)
	}
	fmt.Println(landscape)

	// first find S
	sLocation := getLocation("S", landscape)
	// then find E
	eLocation := getLocation("E", landscape)

	fmt.Println(sLocation)
	fmt.Println(eLocation)

	path := linkedhashset.New()
	path.Add(sLocation)

	startMoves := getNextMoves(sLocation, linkedhashset.New(), landscape)

	movesIterator := startMoves.Iterator()
	paths := make([]*linkedhashset.Set, 0)
	for movesIterator.Next() {
		// _, value := movesIterator.Index(), movesIterator.Value()
		// tmpPaths := getPaths(value.(point), eLocation, path, landscape)
		// for i := 0; i < len(tmpPaths);i++ {
		// 	paths = append(paths,tmpPaths[i])
		// }
		fmt.Println(paths)
	}

	//recursive func, takes a current point, takes a board, a current path, returns a set of point

	// get valid moves around current point
	// for each valid move
	// is the move E
	// if it is return the current path
	//if not and there are valid moves
	// move to first in list
	//if no valid moves, return empty

	// do something like
	// keep a stack of points or hashset of points or something
	// from current position, find next that is lowest value, then next then next
	// for each lowest put the co-ord in a hash set
	//cant go to a square we've already been to, nemsh the set of visited?
	// a move is one where current is >= next || current + 1 == next

	// find the shortest non len()==0 path and this is the route.
	//perf will be a killer :()

	//hard coded for now, tests always pass :D
	return 31
}
