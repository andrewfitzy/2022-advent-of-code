package day_11

func task02ExampleMonkey0(value int, worryLimit int) result {
	// Operation: new = old * 19
	// Test: divisible by 23
	// 	If true: throw to monkey 2
	// 	If false: throw to monkey 3
	worryLevel := value * 19
	return MonkeyInspectionTask02(worryLevel, worryLimit, 23, 2, 3)
}

func task02ExampleMonkey1(value int, worryLimit int) result {
	// Operation: new = old + 6
	// Test: divisible by 19
	// 	If true: throw to monkey 2
	// 	If false: throw to monkey 0
	worryLevel := value + 6
	return MonkeyInspectionTask02(worryLevel, worryLimit, 19, 2, 0)
}

func task02ExampleMonkey2(value int, worryLimit int) result {
	// Operation: new = old * old
	// Test: divisible by 13
	// 	If true: throw to monkey 1
	// 	If false: throw to monkey 3
	worryLevel := value * value
	return MonkeyInspectionTask02(worryLevel, worryLimit, 13, 1, 3)
}

func task02ExampleMonkey3(value int, worryLimit int) result {
	// Operation: new = old + 3
	// Test: divisible by 17
	// 	If true: throw to monkey 0
	// 	If false: throw to monkey 1
	worryLevel := value + 3
	return MonkeyInspectionTask02(worryLevel, worryLimit, 17, 0, 1)
}

func solveExample02(input string) int {
	monkeys := SetupMonkeys(input)
	monkeyInspections := make(map[int]int)
	//multiply all divisors together to find the worry limit (13 * 17 * 19 * 23)
	worryLimit := 96577
	for i := 0; i < 10000; i++ {
		for j := 0; j < len(monkeys); j++ {
			value, moreItems := monkeys[j].Dequeue()

			for moreItems {
				monkeyInspections[j] = monkeyInspections[j] + 1
				result := result{}
				if j == 0 {
					result = task02ExampleMonkey0(value.(int), worryLimit)
				} else if j == 1 {
					result = task02ExampleMonkey1(value.(int), worryLimit)
				} else if j == 2 {
					result = task02ExampleMonkey2(value.(int), worryLimit)
				} else if j == 3 {
					result = task02ExampleMonkey3(value.(int), worryLimit)
				}
				monkeys[result.nextMonkey].Enqueue(result.worryLevel)
				value, moreItems = monkeys[j].Dequeue()
			}
		}
	}
	monkeyBusiness := GetMonkeyBusiness(monkeyInspections)
	return monkeyBusiness
}
