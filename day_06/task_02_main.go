package day_06

import "strings"
import "github.com/emirpasic/gods/sets/hashset"

func solve02(input string) int {
	result := 0
	datastream := strings.Split(input, "")
	datastreamLength := len(datastream)
	// start at index 13 as a word is 14 in length
	for i := 13; i < datastreamLength; i++ {
		set := hashset.New()
		for j := 0; j < 14; j++ {
			//working backwards from i-0 to i-13
			set.Add(datastream[i-j])
		}
		if set.Size() == 14 {
			result = i + 1
			break
		}
	}

	return result
}
