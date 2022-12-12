package day_09

import "bufio"
import "strings"
import "strconv"
import "fmt"
import "github.com/emirpasic/gods/sets/hashset"

func headIsAdjacentToTail(headLocationCol int, headLocationRow int, tailLocationCol int, tailLocationRow int) bool {

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

func solve01(input string) int {
	visited := hashset.New() // empty

	scanner := bufio.NewScanner(strings.NewReader(input))

	// next move head around the board
	headLocationCol := 0
	headLocationRow := 0
	tailLocationCol := 0
	tailLocationRow := 0

	visited.Add(fmt.Sprintf("%v,%v", tailLocationCol, tailLocationRow))
	for scanner.Scan() {
		line := scanner.Text()
		entry := strings.Split(line, " ")
		direction := entry[0]
		spaces, _ := strconv.Atoi(entry[1])
		for i := 0; i < spaces; i++ {
			// move head in the direction it needs to move one space at a time
			if "U" == direction {
				headLocationRow--
			}
			if "D" == direction {
				headLocationRow++
			}
			if "L" == direction {
				headLocationCol--
			}
			if "R" == direction {
				headLocationCol++
			}

			//check if tail is still within distance of head
			if headIsAdjacentToTail(headLocationCol, headLocationRow, tailLocationCol, tailLocationRow) {
				continue
			}
			//move tail
			if headLocationCol == tailLocationCol {
				//if in the same vertical move the row
				if headLocationRow > tailLocationRow {
					tailLocationRow++
				} else {
					tailLocationRow--
				}
			} else if headLocationRow == tailLocationRow {
				//if in the same horizon move the col
				if headLocationCol > tailLocationCol {
					tailLocationCol++
				} else {
					tailLocationCol--
				}
			} else if headLocationRow < tailLocationRow && headLocationCol < tailLocationCol {
				//move diagonally up and left
				tailLocationCol--
				tailLocationRow--
			} else if headLocationRow < tailLocationRow && headLocationCol > tailLocationCol {
				//move diagonally up and right
				tailLocationCol++
				tailLocationRow--
			} else if headLocationRow > tailLocationRow && headLocationCol < tailLocationCol {
				//move diagonally down and left
				tailLocationCol--
				tailLocationRow++
			} else if headLocationRow > tailLocationRow && headLocationCol > tailLocationCol {
				//move diagonally down and left
				tailLocationCol++
				tailLocationRow++
			}

			visited.Add(fmt.Sprintf("%v,%v", tailLocationCol, tailLocationRow))
		}
	}
	return visited.Size()
}
