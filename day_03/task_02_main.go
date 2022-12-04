package day_03

import "bufio"
import "strings"
import "github.com/emirpasic/gods/sets/hashset"

func buildSet(input string) *hashset.Set {
	set := hashset.New() // empty
	slice := strings.Split(input, "")

	for _, item := range slice {
		set.Add(item)
	}
	return set
}

func getCommonElfItem(input1 string, input2 string, input3 string) rune {
	set1 := buildSet(input1)
	set2 := buildSet(input2)
	set3 := buildSet(input3)

	intersection1and2 := set1.Intersection(set2)
	intersection1and2and3 := intersection1and2.Intersection(set3)
	commonItems := intersection1and2and3.Values()

	commonItem := ' '
	if len(commonItems) > 1 {
		panic("Found more than one item")
	}
	for _, item := range commonItems {
		itemAsString := item.(string)
		stringAsRunes := []rune(itemAsString)
		commonItem = stringAsRunes[0]

	}
	return commonItem
}

func solve02(input string) int {
	total := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for {
		item1 := ""
		item2 := ""
		item3 := ""

		if scanner.Scan() {
			item1 = scanner.Text()
		} else {
			break
		}

		if scanner.Scan() {
			item2 = scanner.Text()
		} else {
			break
		}

		if scanner.Scan() {
			item3 = scanner.Text()
		} else {
			break
		}
		commonItem := getCommonElfItem(item1, item2, item3)
		itemPriority := getItemPriority(commonItem)
		total = total + itemPriority
	}

	return total
}
