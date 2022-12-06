package day_06

import "strings"
import "github.com/emirpasic/gods/sets/hashset"

func solve01(input string) int {
	result := 0
	datastream := strings.Split(input, "")
	datastreamLength := len(datastream)
	// start at index 3 as a marker is 4 in length
	for i := 3; i < datastreamLength; i++ {
		set := hashset.New()
		for j := 0; j < 4; j++ {
			//working backwards from i-0 to i-3
			set.Add(datastream[i-j])
		}
		if set.Size() == 4 {
			result = i + 1
			break
		}
	}

	return result
}
