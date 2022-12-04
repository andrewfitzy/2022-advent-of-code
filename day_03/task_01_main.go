package day_03

import "bufio"
import "strings"

func getCommonCompartmentItem(left string, right string) rune {
	for _, leftCharacter := range left {
		for _, rightCharacter := range right {
			if leftCharacter == rightCharacter {
				return leftCharacter
			}
		}
	}
	panic("No Matching character found")
}

func solve01(input string) int {
	total := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		midpoint := len(line) / 2
		left := line[0:midpoint]
		right := line[midpoint:]
		commonItem := getCommonCompartmentItem(left, right)
		itemPriority := getItemPriority(commonItem)
		total = total + itemPriority
	}

	return total
}
