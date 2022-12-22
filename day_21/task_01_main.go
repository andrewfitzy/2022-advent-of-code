package day_12

import "bufio"
import "strings"
import "fmt"
import "strconv"

// import "github.com/emirpasic/gods/sets/linkedhashset"

type monkey struct {
	name      string
	left      string
	right     string
	operation string
	value     int
}

func makeMonkey(input string) monkey {
	monkey := monkey{}

	monkeyParts := strings.Split(input, " ")

	monkey.name = strings.Replace(monkeyParts[0], ":", "", -1)

	if len(monkeyParts) == 2 {
		//dbpl: 5
		monkey.value, _ = strconv.Atoi(monkeyParts[1])
	}
	if len(monkeyParts) == 4 {
		//root: pppw + sjmn
		monkey.left = monkeyParts[1]
		monkey.operation = monkeyParts[2]
		monkey.right = monkeyParts[3]
	}

	return monkey
}

func getMonkeyResult(monkey monkey) int {
	left := 0
	right := 0
	result := 0

	if len(monkey.left) == 0 && len(monkey.right) == 0 {
		//was always a number money
		return monkey.value
	}
	if tmpLeft, err := strconv.Atoi(monkey.left); err == nil {
		left = tmpLeft
	} else {
		return result
	}

	if tmpRight, err := strconv.Atoi(monkey.right); err == nil {
		right = tmpRight
	} else {
		return result
	}
	switch monkey.operation {
	case "+":
		result = left + right
	case "-":
		result = left - right
	case "*":
		result = left * right
	case "/":
		result = left / right
	}
	return result
}

func monkeyIsComplete(monkey monkey) bool {
	if len(monkey.left) == 0 && len(monkey.right) == 0 {
		//was always a number money
		return true
	}
	if _, err := strconv.Atoi(monkey.left); err != nil {
		return false
	}
	if _, err := strconv.Atoi(monkey.right); err != nil {
		return false
	}

	return getMonkeyResult(monkey) == monkey.value
}

func resolveMonkey(monkeyName string, tracker map[string]monkey) string {
	if _, err := strconv.Atoi(monkeyName); err == nil {
		return monkeyName
	}
	monkey := tracker[monkeyName]
	if monkeyIsComplete(monkey) {
		result := getMonkeyResult(monkey)
		return fmt.Sprint(result)
	}
	return monkeyName
}

func isSolved(tracker map[string]monkey) bool {
	solved := true
	for _, monkey := range tracker {
		solved = solved && monkeyIsComplete(monkey)
	}
	return solved
}

func solve01(input string, targetMonkey string) int {
	tracker := make(map[string]monkey)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		monkey := makeMonkey(line)
		tracker[monkey.name] = monkey
	}
	fmt.Println(tracker)
	//repeat until all are true
	allAreSolved := false
	for !allAreSolved {
		for j := range tracker {
			tmpMonkey := tracker[j]
			if !monkeyIsComplete(tmpMonkey) {
				left := resolveMonkey(tmpMonkey.left, tracker)
				right := resolveMonkey(tmpMonkey.right, tracker)
				value := getMonkeyResult(tmpMonkey)

				newMonkey := monkey{
					name:      tmpMonkey.name,
					left:      left,
					right:     right,
					operation: tmpMonkey.operation,
					value:     value,
				}
				tracker[tmpMonkey.name] = newMonkey
			}
		}
		allAreSolved = isSolved(tracker)
	}
	fmt.Println(tracker)

	return tracker[targetMonkey].value
}
