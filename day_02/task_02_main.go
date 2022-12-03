package day_02

import "bufio"
import "strings"

func solve02(input string) int {
	gameScores := map[string]int{
		"AX": 3, //3 scissors + 0 loss
		"AY": 4, //1 rock + 3 draw
		"AZ": 8, //2 paper + 6 win
		"BX": 1, //1 rock + 0 loss
		"BY": 5, //2 paper + 3 draw
		"BZ": 9, //3 scissors + 6 win
		"CX": 2, //2 paper + 0 loss
		"CY": 6, //3 scissors + 3 draw
		"CZ": 7, //1 rock + 6 win
	}

	total := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.Replace(line, " ", "", -1)
		total = total + gameScores[trimmed]
	}

	return total
}
