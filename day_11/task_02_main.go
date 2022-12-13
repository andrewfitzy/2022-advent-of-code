package day_11

func task02Monkey0(value int, worryLimit int) result {
	// Operation: new = old * 19
	// Test: divisible by 7
	//   If true: throw to monkey 2
	//   If false: throw to monkey 3
	worryLevel := value * 19
	return MonkeyInspectionTask02(worryLevel, worryLimit, 7, 2, 3)
}

func task02Monkey1(value int, worryLimit int) result {
	// Operation: new = old + 1
	// Test: divisible by 19
	//   If true: throw to monkey 4
	//   If false: throw to monkey 6
	worryLevel := value + 1
	return MonkeyInspectionTask02(worryLevel, worryLimit, 19, 4, 6)
}

func task02Monkey2(value int, worryLimit int) result {
	// Operation: new = old + 6
	// Test: divisible by 5
	//   If true: throw to monkey 7
	//   If false: throw to monkey 5
	worryLevel := value + 6
	return MonkeyInspectionTask02(worryLevel, worryLimit, 5, 7, 5)
}

func task02Monkey3(value int, worryLimit int) result {
	// Operation: new = old + 5
	// Test: divisible by 11
	//   If true: throw to monkey 5
	//   If false: throw to monkey 2
	worryLevel := value + 5
	return MonkeyInspectionTask02(worryLevel, worryLimit, 11, 5, 2)
}

func task02Monkey4(value int, worryLimit int) result {
	// 	Monkey 4:
	// Operation: new = old * old
	// Test: divisible by 17
	//   If true: throw to monkey 0
	//   If false: throw to monkey 3
	worryLevel := value * value
	return MonkeyInspectionTask02(worryLevel, worryLimit, 17, 0, 3)
}

func task02Monkey5(value int, worryLimit int) result {
	// Monkey 5:
	// Operation: new = old + 7
	// Test: divisible by 13
	// If true: throw to monkey 1
	// If false: throw to monkey 7
	worryLevel := value + 7
	return MonkeyInspectionTask02(worryLevel, worryLimit, 13, 1, 7)
}

func task02Monkey6(value int, worryLimit int) result {
	// Monkey 6:
	// Operation: new = old * 7
	// Test: divisible by 2
	//   If true: throw to monkey 0
	//   If false: throw to monkey 4
	worryLevel := value * 7
	return MonkeyInspectionTask02(worryLevel, worryLimit, 2, 0, 4)
}

func task02Monkey7(value int, worryLimit int) result {
	// Monkey 7:
	// Operation: new = old + 2
	// Test: divisible by 3
	//   If true: throw to monkey 1
	//   If false: throw to monkey 6
	worryLevel := value + 2
	return MonkeyInspectionTask02(worryLevel, worryLimit, 3, 1, 6)
}

func solve02(input string) int {
	monkeys := SetupMonkeys(input)
	monkeyInspections := make(map[int]int)
	//multiply all divisors together to find the worry limit (2 * 3 * 5 * 7 * 11 * 13 * 17 * 19)
	worryLimit := 9699690
	for i := 0; i < 10000; i++ {
		for j := 0; j < len(monkeys); j++ {
			value, moreItems := monkeys[j].Dequeue()

			for moreItems {
				monkeyInspections[j] = monkeyInspections[j] + 1
				result := result{}
				if j == 0 {
					result = task02Monkey0(value.(int), worryLimit)
				} else if j == 1 {
					result = task02Monkey1(value.(int), worryLimit)
				} else if j == 2 {
					result = task02Monkey2(value.(int), worryLimit)
				} else if j == 3 {
					result = task02Monkey3(value.(int), worryLimit)
				} else if j == 4 {
					result = task02Monkey4(value.(int), worryLimit)
				} else if j == 5 {
					result = task02Monkey5(value.(int), worryLimit)
				} else if j == 6 {
					result = task02Monkey6(value.(int), worryLimit)
				} else if j == 7 {
					result = task02Monkey7(value.(int), worryLimit)
				}
				monkeys[result.nextMonkey].Enqueue(result.worryLevel)
				value, moreItems = monkeys[j].Dequeue()
			}
		}
	}
	monkeyBusiness := GetMonkeyBusiness(monkeyInspections)
	return monkeyBusiness
}
