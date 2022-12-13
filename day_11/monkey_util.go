package day_11

import "sort"
import "bufio"
import "strings"
import "strconv"
import "github.com/emirpasic/gods/queues"
import "github.com/emirpasic/gods/queues/arrayqueue"

type result struct {
	nextMonkey int
	worryLevel int
}

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

func MonkeyInspection(worryLevel int, testDivisor int, passMonkey int, failMonkey int) result {
	reducedWorry := worryLevel / 3
	if reducedWorry%testDivisor == 0 {
		return result{
			nextMonkey: passMonkey,
			worryLevel: reducedWorry,
		}
	}
	return result{
		nextMonkey: failMonkey,
		worryLevel: reducedWorry,
	}
}

func SetupMonkeys(input string) map[int]queues.Queue {
	monkeys := make(map[int]queues.Queue)
	scanner := bufio.NewScanner(strings.NewReader(input))
	moneyPointer := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "Starting items: ") {
			entry := line[16:]
			items := strings.Split(entry, ", ")
			monkey := arrayqueue.New()
			for _, item := range items {

				if s, err := strconv.Atoi(item); err == nil {
					monkey.Enqueue(s)
				}

			}
			monkeys[moneyPointer] = monkey
			moneyPointer++
		}
	}
	return monkeys
}
