package day_11

func solve01(input string, iterations int) int {
	monkeys := SetupMonkeys(input)

	monkeyInspections := make(map[int]int)
	for i := 0; i < iterations; i++ {
		for j := 0; j < len(monkeys); j++ {
			monkey := monkeys[j]

			value, moreItems := monkey.items.Dequeue()

			for moreItems {
				monkeyInspections[j] = monkeyInspections[j] + 1

				worryLevel := monkey.MonkeyPlay(value.(int))
				result := monkey.MonkeyBored(worryLevel)

				nextMonkey := monkeys[result.nextMonkey]
				nextMonkey.items.Enqueue(result.worryLevel)

				value, moreItems = monkey.items.Dequeue()
			}
		}
	}
	monkeyBusiness := GetMonkeyBusiness(monkeyInspections)
	return monkeyBusiness
}
