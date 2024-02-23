package day_11

func solve02(input string, iterations int) int {
	monkeys := SetupMonkeys(input)
	monkeyInspections := make(map[int]int)

	worryLimit := GetWorryLimit(monkeys)

	for i := 0; i < iterations; i++ {
		for j := 0; j < len(monkeys); j++ {
			monkey := monkeys[j]

			value, moreItems := monkey.items.Dequeue()

			for moreItems {
				monkeyInspections[j] = monkeyInspections[j] + 1

				worryLevel := monkey.MonkeyPlay(value.(int))
				result := monkey.MonkeyBored(worryLevel, worryLimit)

				nextMonkey := monkeys[result.nextMonkey]
				nextMonkey.items.Enqueue(result.worryLevel)

				value, moreItems = monkey.items.Dequeue()
			}
		}
	}
	monkeyBusiness := GetMonkeyBusiness(monkeyInspections)
	return monkeyBusiness
}

func GetWorryLimit(monkeys map[int]monkey) int {
	worryLimit := 1
	for _, monkey := range monkeys {
		worryLimit = worryLimit * monkey.testDivisor
	}
	return worryLimit
}
