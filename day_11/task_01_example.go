package day_11

func task01ExampleMonkey0(value int) result {
	// Operation: new = old * 19
	// Test: divisible by 23
	// 	If true: throw to monkey 2
	// 	If false: throw to monkey 3
	worryLevel := value * 19
	return MonkeyInspection(worryLevel, 23, 2, 3)
}

func task01ExampleMonkey1(value int) result {
	// Operation: new = old + 6
	// Test: divisible by 19
	// 	If true: throw to monkey 2
	// 	If false: throw to monkey 0
	worryLevel := value + 6
	return MonkeyInspection(worryLevel, 19, 2, 0)
}

func task01ExampleMonkey2(value int) result {
	// Operation: new = old * old
	// Test: divisible by 13
	// 	If true: throw to monkey 1
	// 	If false: throw to monkey 3
	worryLevel := value * value
	return MonkeyInspection(worryLevel, 13, 1, 3)
}

func task01ExampleMonkey3(value int) result {
	// Operation: new = old + 3
	// Test: divisible by 17
	// 	If true: throw to monkey 0
	// 	If false: throw to monkey 1
	worryLevel := value + 3
	return MonkeyInspection(worryLevel, 17, 0, 1)
}

func solveExample01(input string) int {
	monkeys := SetupMonkeys(input)
	monkeyInspections := make(map[int]int)
	for i := 0; i < 20; i++ {
		for j := 0; j < len(monkeys); j++ {
			value, moreItems := monkeys[j].Dequeue()

			for moreItems {
				monkeyInspections[j] = monkeyInspections[j] + 1
				if j == 0 {
					result := task01ExampleMonkey0(value.(int))
					monkeys[result.nextMonkey].Enqueue(result.worryLevel)
				} else if j == 1 {
					result := task01ExampleMonkey1(value.(int))
					monkeys[result.nextMonkey].Enqueue(result.worryLevel)
				} else if j == 2 {
					result := task01ExampleMonkey2(value.(int))
					monkeys[result.nextMonkey].Enqueue(result.worryLevel)
				} else if j == 3 {
					result := task01ExampleMonkey3(value.(int))
					monkeys[result.nextMonkey].Enqueue(result.worryLevel)
				}

				value, moreItems = monkeys[j].Dequeue()
			}
		}
	}
	monkeyBusiness := GetMonkeyBusiness(monkeyInspections)
	return monkeyBusiness
}
