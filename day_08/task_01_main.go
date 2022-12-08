package day_08

import "bufio"
import "strings"
import "strconv"

func isGreaterThanEntries(lineOfSight []int, tree int) bool {
	if len(lineOfSight) == 0 { // it's an edge space, always visible
		return true
	}
	for _, tmpTree := range lineOfSight {
		if tree <= tmpTree {
			return false
		}
	}
	return true
}

func treeIsVisible(forest [][]int, irow int, icolumn int) bool {
	treesToLeft := make([]int, 0)
	for i := 0; i < icolumn; i++ {
		treesToLeft = append(treesToLeft, forest[irow][i])
	}

	treesToRight := make([]int, 0)
	for i := icolumn + 1; i < len(forest[irow]); i++ {
		treesToRight = append(treesToRight, forest[irow][i])
	}

	treesAbove := make([]int, 0)
	for i := 0; i < irow; i++ {
		treesAbove = append(treesAbove, forest[i][icolumn])
	}

	treesBelow := make([]int, 0)
	for i := irow + 1; i < len(forest); i++ {
		treesBelow = append(treesBelow, forest[i][icolumn])
	}

	tree := forest[irow][icolumn]

	visibleFromLeft := isGreaterThanEntries(treesToLeft, tree)
	visibleFromRight := isGreaterThanEntries(treesToRight, tree)
	visibleFromAbove := isGreaterThanEntries(treesAbove, tree)
	visibleFromBelow := isGreaterThanEntries(treesBelow, tree)
	isVisible := false
	if visibleFromLeft || visibleFromRight || visibleFromAbove || visibleFromBelow {
		isVisible = true
	}
	return isVisible
}

func solve01(input string) int {
	//process input into 2d array
	forest := make([][]int, 0)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		strTrees := strings.Split(line, "")
		trees := make([]int, 0, len(strTrees))
		for _, strTree := range strTrees {
			tree, _ := strconv.Atoi(strTree)
			trees = append(trees, tree)
		}
		forest = append(forest, trees)
	}

	visibleTrees := 0
	for irow, row := range forest {
		for icolumn := range row {
			if treeIsVisible(forest, irow, icolumn) {
				visibleTrees = visibleTrees + 1
			}
		}
	}

	return visibleTrees
}
