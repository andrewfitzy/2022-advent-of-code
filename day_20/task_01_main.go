package day_12

import "bufio"
import "strings"
import "fmt"
import "strconv"

// import "github.com/emirpasic/gods/sets/linkedhashset"

type item struct {
	startLocation int
	value         int
}

func trackerToString(tracker []item) string {
	staged := make([]string, len(tracker))
	for i := 0; i < len(tracker); i++ {
		staged[i] = strconv.Itoa(tracker[i].value)
	}
	stagedString := strings.Join(staged, ", ")
	return stagedString
}

func solve01(input string) int {
	tracker := make([]item, 0)
	scanner := bufio.NewScanner(strings.NewReader(input))
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		value, _ := strconv.Atoi(line)
		entry := item{
			value:         value,
			startLocation: i,
		}
		i++
		tracker = append(tracker, entry)
	}

	fmt.Println(trackerToString(tracker))
	fmt.Println("Processing")

	for i := 0; i < len(tracker); i++ {
		fmt.Printf("Index %v \n", i)
	}

	// // -------- old code --------

	// //now process
	// for j := 0; j < len(tracker) ; j++ {
	// //for j := 0; j < 3000 ; j++ {
	// 	index := j % len(tracker)
	// 	fmt.Printf("Process index %v with value %v\n", index, tracker[index].value)
	// 	//shuffle means adding and removing 1
	// 	if tracker[index].value >= 0 {
	// 		//shuffle right
	// 		fmt.Printf("Shuffle %v right by %v \n", tracker[index].value, tracker[index].value)
	// 		for k := 0;k < tracker[index].value;k++ {
	// 			fmt.Printf("iteration %v for value %v current loc %v\n", k, tracker[index].value, tracker[index].currentLocation)
	// 			if tracker[index].currentLocation+1 == len(tracker){
	// 				tracker[index].currentLocation = 0
	// 			} else {
	// 				tracker[index].currentLocation++
	// 			}
	// 			fmt.Printf("iteration %v new current loc %v\n", k, tracker[index].currentLocation)
	// 			fmt.Println(tracker)

	// 			shuffleNumber := tracker[index].value
	// 			for l := 0;l < len(tracker);l++ {
	// 				fmt.Printf("iteration l %v checking %v\n", l, tracker[index].currentLocation)
	// 				if l == index {
	// 					continue
	// 				}

	// 				if tracker[l].currentLocation > index && tracker[l].currentLocation <= index + shuffleNumber {
	// 					fmt.Printf("iteration l %v current %v > index %v\n", l, tracker[l], tracker[index])
	// 					tracker[l].currentLocation--
	// 				}
	// 			}
	// 			fmt.Println(tracker)
	// 		}
	// 	} else {
	// 		//shuffle left
	// 		fmt.Printf("Shufle left by %v \n", tracker[index].value)
	// 		// count := -1 * tracker[index].value
	// 		// for k := 0;k < count;k++ {
	// 		// 	tracker[index].currentLocation--
	// 		// 	for l := 0;l < len(tracker);l++ {
	// 		// 		if l == tracker[index].currentLocation {
	// 		// 			continue
	// 		// 		}
	// 		// 		if tracker[l].currentLocation < tracker[index].currentLocation {
	// 		// 			tracker[l].currentLocation++
	// 		// 		}
	// 		// 	}
	// 		// }
	// 	}

	// 	fmt.Println(trackerToString(tracker))
	// }
	return 3
}
