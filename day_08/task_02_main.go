package day_08

import "bufio"
import "strings"
import "strconv"

func getViewingDistanceLeftOrUp(lineOfSight []int, tree int) int {
	distance := 0
	if len(lineOfSight) == 0 { // it's an edge space, always visible
		return distance
	}
	for i := len(lineOfSight) - 1; i >= 0; i-- {
		distance = distance + 1
		if lineOfSight[i] >= tree {
			break
		}
	}
	return distance
}

func getViewingDistanceRightOrDown(lineOfSight []int, tree int) int {
	distance := 0
	if len(lineOfSight) == 0 { // it's an edge space, always visible
		return distance
	}
	for i := 0; i < len(lineOfSight); i++ {
		distance = distance + 1
		if lineOfSight[i] >= tree {
			break
		}
	}
	return distance
}

func getScenicScore(forest [][]int, irow int, icolumn int) int {
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

	visibleFromLeft := getViewingDistanceLeftOrUp(treesToLeft, tree)
	visibleFromRight := getViewingDistanceRightOrDown(treesToRight, tree)
	visibleFromAbove := getViewingDistanceLeftOrUp(treesAbove, tree)
	visibleFromBelow := getViewingDistanceRightOrDown(treesBelow, tree)
	scenicScore := visibleFromLeft * visibleFromRight * visibleFromAbove * visibleFromBelow

	return scenicScore
}

func solve02(input string) int {

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

	bestScenicScore := 0
	for irow, row := range forest {
		for icolumn := range row {
			scenicScore := getScenicScore(forest, irow, icolumn)
			if scenicScore > bestScenicScore {
				bestScenicScore = scenicScore
			}
		}
	}

	return bestScenicScore
}
