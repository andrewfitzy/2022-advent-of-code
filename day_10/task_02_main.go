package day_10

import "bufio"
import "strings"
import "strconv"

func solve02(input string) string {
	result := ""
	register := 1
	currentCycle := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		if currentCycle >= register-1 && currentCycle <= register+1 {
			result = result + "#"
		} else {
			result = result + "."
		}
		currentCycle++
		if currentCycle%40 == 0 {
			result = result + "\n"
			currentCycle = 0
		}

		entry := strings.Split(line, " ")
		op := entry[0]
		if op != "noop" {
			if currentCycle >= register-1 && currentCycle <= register+1 {
				result = result + "#"
			} else {
				result = result + "."
			}
			currentCycle++
			if currentCycle%40 == 0 {
				result = result + "\n"
				currentCycle = 0
			}

			value, _ := strconv.Atoi(entry[1])
			register = register + value
		}
	}
	return string(result)
}
