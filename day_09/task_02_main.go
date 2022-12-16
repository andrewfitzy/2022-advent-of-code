package day_09

import "bufio"
import "strings"
import "strconv"
import "fmt"
import "github.com/emirpasic/gods/sets/hashset"

type point struct {
	col int
	row int
}

func nextIsAdjacentToCurrent(headLocationCol int, headLocationRow int, tailLocationCol int, tailLocationRow int) bool {

	// above left
	if headLocationCol-1 == tailLocationCol && headLocationRow-1 == tailLocationRow {
		return true
	}
	// above
	if headLocationCol == tailLocationCol && headLocationRow-1 == tailLocationRow {
		return true
	}
	// above right
	if headLocationCol+1 == tailLocationCol && headLocationRow-1 == tailLocationRow {
		return true
	}

	// left
	if headLocationCol-1 == tailLocationCol && headLocationRow == tailLocationRow {
		return true
	}
	// both on the same square
	if headLocationCol == tailLocationCol && headLocationRow == tailLocationRow {
		return true
	}
	// right
	if headLocationCol+1 == tailLocationCol && headLocationRow == tailLocationRow {
		return true
	}

	// below left
	if headLocationCol-1 == tailLocationCol && headLocationRow+1 == tailLocationRow {
		return true
	}
	// below
	if headLocationCol == tailLocationCol && headLocationRow+1 == tailLocationRow {
		return true
	}
	// below right
	if headLocationCol+1 == tailLocationCol && headLocationRow+1 == tailLocationRow {
		return true
	}

	return false
}

func solve02(input string) int {
	// choose a starting position
	headLocationCol := 0
	headLocationRow := 0

	rope := make([]point, 10)
	for i := 0; i < len(rope); i++ {
		rope[i] = point{
			col: headLocationCol,
			row: headLocationRow,
		}
	}

	visited := hashset.New() // empty

	scanner := bufio.NewScanner(strings.NewReader(input))

	visited.Add(fmt.Sprintf("%v,%v", rope[9].col, rope[9].row))

	for scanner.Scan() {
		line := scanner.Text()
		entry := strings.Split(line, " ")
		direction := entry[0]
		spaces, _ := strconv.Atoi(entry[1])
		for i := 0; i < spaces; i++ {
			// move head in the direction it needs to move one space at a time
			if direction == "U" {
				headLocationRow--
			}
			if direction == "D" {
				headLocationRow++
			}
			if direction == "L" {
				headLocationCol--
			}
			if direction == "R" {
				headLocationCol++
			}

			rope[0].col = headLocationCol
			rope[0].row = headLocationRow
			tmpCurrentCol := headLocationCol
			tmpCurrentRow := headLocationRow

			for i := 1; i < len(rope); i++ {
				nextLocationCol := rope[i].col
				nextLocationRow := rope[i].row

				//check if remaining knots are within distance
				if nextIsAdjacentToCurrent(tmpCurrentCol, tmpCurrentRow, nextLocationCol, nextLocationRow) {
					tmpCurrentCol = nextLocationCol
					tmpCurrentRow = nextLocationRow
					continue
				}
				//move tail
				if tmpCurrentCol == nextLocationCol {
					//if in the same vertical move the row
					if tmpCurrentRow > nextLocationRow {
						nextLocationRow++
					} else {
						nextLocationRow--
					}
				} else if tmpCurrentRow == nextLocationRow {
					//if in the same horizon move the col
					if tmpCurrentCol > nextLocationCol {
						nextLocationCol++
					} else {
						nextLocationCol--
					}
				} else if tmpCurrentRow < nextLocationRow && tmpCurrentCol < nextLocationCol {
					//move diagonally up and left
					nextLocationCol--
					nextLocationRow--
				} else if tmpCurrentRow < nextLocationRow && tmpCurrentCol > nextLocationCol {
					//move diagonally up and right
					nextLocationCol++
					nextLocationRow--
				} else if tmpCurrentRow > nextLocationRow && tmpCurrentCol < nextLocationCol {
					//move diagonally down and left
					nextLocationCol--
					nextLocationRow++
				} else if tmpCurrentRow > nextLocationRow && tmpCurrentCol > nextLocationCol {
					//move diagonally down and left
					nextLocationCol++
					nextLocationRow++
				}
				rope[i].col = nextLocationCol
				rope[i].row = nextLocationRow
				tmpCurrentCol = nextLocationCol
				tmpCurrentRow = nextLocationRow
			}
			visited.Add(fmt.Sprintf("%v,%v", rope[9].col, rope[9].row))
		}
	}
	return visited.Size()
}
