package day_03

import "bufio"
import "strings"

func getCommonItem(left string, right string) rune {
	for _, leftCharacter := range left {
		for _, rightCharacter := range right {
			if leftCharacter == rightCharacter {
				return leftCharacter
			}
		}
	}
	panic("No Matching character found")
}

func getItemPriority(item rune) int {
	itemPriority := 0
	itemAsciiValue := int(item)
	if itemAsciiValue >= 65 && itemAsciiValue <= 90 {
		itemPriority = itemAsciiValue - 38
	}
	if itemAsciiValue >= 97 && itemAsciiValue <= 122 {
		itemPriority = itemAsciiValue - 96
	}
	return itemPriority
}

func solve01(input string) int {
	total := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		midpoint := len(line) / 2
		left := line[0:midpoint]
		right := line[midpoint:]
		commonItem := getCommonItem(left, right)
		itemPriority := getItemPriority(commonItem)
		total = total + itemPriority
	}

	return total
}
