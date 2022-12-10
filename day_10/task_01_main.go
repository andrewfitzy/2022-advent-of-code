package day_10

import "bufio"
import "strings"
import "strconv"

func solve01(input string) int {
	result := 0
	register := 1
	reportCycle := 20
	currentCycle := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		currentCycle++
		if reportCycle == currentCycle {
			stepResult := register * reportCycle
			result = result + stepResult
			reportCycle = reportCycle + 40
		}

		entry := strings.Split(line, " ")
		op := entry[0]
		if op != "noop" {
			currentCycle++
			if reportCycle == currentCycle {
				stepResult := register * reportCycle
				result = result + stepResult
				reportCycle = reportCycle + 40
			}
			value, _ := strconv.Atoi(entry[1])
			register = register + value
		}
	}
	return result
}
