package day_02

import "bufio"
import "strings"

func solve01(input string) int {
	gameScores := map[string]int{
		"AX": 4, //1 rock + 4 draw
		"AY": 8, //2 paper + 6 win
		"AZ": 3, //3 scissors + 0 loss
		"BX": 1, //1 rock + 0 loss
		"BY": 5, //2 paper + 5 draw
		"BZ": 9, //3 scissors + 6 win
		"CX": 7, //1 rock + 6 win
		"CY": 2, //2 paper + 0 loss
		"CZ": 6, //3 scissors + 3 draw
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
