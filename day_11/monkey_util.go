package day_11

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/emirpasic/gods/queues/arrayqueue"
)

func GetMonkeyBusiness(monkeys map[int]int) int {
	monkeyValues := make([]int, 0, len(monkeys))
	for _, value := range monkeys {
		monkeyValues = append(monkeyValues, value)
	}
	sort.Ints(monkeyValues)
	highestValue := monkeyValues[len(monkeyValues)-1]
	secondHighestValue := monkeyValues[len(monkeyValues)-2]
	monkeyBusiness := highestValue * secondHighestValue
	return monkeyBusiness
}

func SetupMonkeys(input string) map[int]monkey {
	monkeys := make(map[int]monkey)

	monkeyStrings := strings.Split(input, "\n\n")
	for _, monkeyString := range monkeyStrings {
		monkeyLines := strings.Split(monkeyString, "\n")
		monkeyNumber := -1
		monkeyItems := arrayqueue.New()
		var monkeyOperation func(int) int
		testDivisor := 1
		trueTargetMonkey := -1
		falseTargetMonkey := -1
		for _, monkeyLine := range monkeyLines {
			trimmedLine := strings.TrimSpace(monkeyLine)
			if strings.HasPrefix(trimmedLine, "Monkey") {
				monkeyNumber = GetMonkeyNumber(trimmedLine)
			}
			if strings.HasPrefix(trimmedLine, "Starting items") {
				items := GetMonkeyItems(trimmedLine)
				for j := 0; j < len(items); j++ {
					monkeyItems.Enqueue(items[j])
				}
			}
			if strings.HasPrefix(trimmedLine, "Operation") {
				monkeyOperation = GetOperation(trimmedLine)
			}
			if strings.HasPrefix(trimmedLine, "Test") {
				testDivisor = GetNumberAtEndOfLine(trimmedLine)
			}
			if strings.HasPrefix(trimmedLine, "If true") {
				trueTargetMonkey = GetNumberAtEndOfLine(trimmedLine)
			}
			if strings.HasPrefix(trimmedLine, "If false") {
				falseTargetMonkey = GetNumberAtEndOfLine(trimmedLine)
			}
		}
		monkey := monkey{monkeyNumber, monkeyItems, monkeyOperation, testDivisor, trueTargetMonkey, falseTargetMonkey}
		monkeys[monkeyNumber] = monkey
	}
	return monkeys
}

func GetMonkeyNumber(input string) int {
	monkeyNumber := -1
	monkeyNumberStr := strings.ReplaceAll(input[len("Monkey "):], ":", "")
	if value, err := strconv.Atoi(monkeyNumberStr); err == nil {
		monkeyNumber = value
	}
	return monkeyNumber
}

func GetMonkeyItems(input string) []int {
	entry := strings.TrimSpace(input)[16:]
	items := strings.Split(entry, ", ")

	monkeyItems := make([]int, len(items))
	for index, item := range items {
		if value, err := strconv.Atoi(item); err == nil {
			monkeyItems[index] = value
		}
	}
	return monkeyItems
}

func GetNumberAtEndOfLine(input string) int {
	tmpNumber := -1
	strTokens := strings.Split(input, " ")
	if value, err := strconv.Atoi(strTokens[len(strTokens)-1]); err == nil {
		tmpNumber = value
	}
	return tmpNumber
}

func GetOperation(input string) func(int) int {
	var operator rune
	var value int = 0
	if input == "Operation: new = old * old" || input == "Operation: new = old + old" {
		fmt.Sscanf(input, "Operation: new = old %c old", &operator)
	} else {
		fmt.Sscanf(input, "Operation: new = old %c %d", &operator, &value)
	}

	operation := func(n int) int {
		var valueToUse int = value
		if valueToUse == 0 {
			valueToUse = n
		}
		if operator == '+' {
			return n + valueToUse
		}
		return n * valueToUse
	}
	return operation
}
