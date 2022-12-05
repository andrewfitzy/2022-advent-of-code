package day_05

import "bufio"
import "strconv"
import "strings"
import "regexp"
import "github.com/emirpasic/gods/stacks/arraystack"
import "github.com/emirpasic/gods/stacks"

func solve01(startPosition map[int]string, input string) string {
	stacks := make(map[int]stacks.Stack)

	// Process the initial config into a map of stacks.
	for index, config := range startPosition {
		stack := arraystack.New()
		configArray := strings.Split(config, "")
		for _, item := range configArray {
			stack.Push(item)
		}
		stacks[index] = stack
	}

	matcher := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	///iterate through the input, extract from, to and number using a regex
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		matches := matcher.FindStringSubmatch(line)
		numberOfCrates, _ := strconv.Atoi(matches[1])
		moveFrom, _ := strconv.Atoi(matches[2])
		moveTo, _ := strconv.Atoi(matches[3])

		for i := 0; i < numberOfCrates; i++ {
			val, _ := stacks[moveFrom].Pop()
			stacks[moveTo].Push(val)
		}
	}

	result := ""
	for i := 1; i <= len(stacks); i++ {
		peek, _ := stacks[i].Peek()
		result = result + peek.(string)
	}

	return result
}
