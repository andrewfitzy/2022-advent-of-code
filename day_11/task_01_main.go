package day_11

import "fmt"

func task01Monkey0(value int) result {
	// Operation: new = old * 19
	// Test: divisible by 7
	//   If true: throw to monkey 2
	//   If false: throw to monkey 3
	worryLevel := value * 19
	return MonkeyInspection(worryLevel, 7, 2, 3)
}

func task01Monkey1(value int) result {
	// Operation: new = old + 1
	// Test: divisible by 19
	//   If true: throw to monkey 4
	//   If false: throw to monkey 6
	worryLevel := value + 1
	return MonkeyInspection(worryLevel, 19, 4, 6)
}

func task01Monkey2(value int) result {
	// Operation: new = old + 6
	// Test: divisible by 5
	//   If true: throw to monkey 7
	//   If false: throw to monkey 5
	worryLevel := value + 6
	return MonkeyInspection(worryLevel, 5, 7, 5)
}

func task01Monkey3(value int) result {
	// Operation: new = old + 5
	// Test: divisible by 11
	//   If true: throw to monkey 5
	//   If false: throw to monkey 2
	worryLevel := value + 5
	return MonkeyInspection(worryLevel, 11, 5, 2)
}

func task01Monkey4(value int) result {
	// 	Monkey 4:
	// Operation: new = old * old
	// Test: divisible by 17
	//   If true: throw to monkey 0
	//   If false: throw to monkey 3
	worryLevel := value * value
	return MonkeyInspection(worryLevel, 17, 0, 3)
}

func task01Monkey5(value int) result {
	// Monkey 5:
	// Operation: new = old + 7
	// Test: divisible by 13
	// If true: throw to monkey 1
	// If false: throw to monkey 7
	worryLevel := value + 7
	return MonkeyInspection(worryLevel, 13, 1, 7)
}

func task01Monkey6(value int) result {
	// Monkey 6:
	// Operation: new = old * 7
	// Test: divisible by 2
	//   If true: throw to monkey 0
	//   If false: throw to monkey 4
	worryLevel := value * 7
	return MonkeyInspection(worryLevel, 2, 0, 4)
}

func task01Monkey7(value int) result {
	// Monkey 7:
	// Operation: new = old + 2
	// Test: divisible by 3
	//   If true: throw to monkey 1
	//   If false: throw to monkey 6
	worryLevel := value + 2
	return MonkeyInspection(worryLevel, 3, 1, 6)
}

func solve01(input string) int {
	monkeys := SetupMonkeys(input)
	fmt.Println(monkeys)
	monkeyInspections := make(map[int]int)
	for i := 0; i < 20; i++ {
		for j := 0; j < len(monkeys); j++ {
			value, moreItems := monkeys[j].Dequeue()

			for moreItems {
				monkeyInspections[j] = monkeyInspections[j] + 1
				if j == 0 {
					result := task01Monkey0(value.(int))
					monkeys[result.nextMonkey].Enqueue(result.worryLevel)
				} else if j == 1 {
					result := task01Monkey1(value.(int))
					monkeys[result.nextMonkey].Enqueue(result.worryLevel)
				} else if j == 2 {
					result := task01Monkey2(value.(int))
					monkeys[result.nextMonkey].Enqueue(result.worryLevel)
				} else if j == 3 {
					result := task01Monkey3(value.(int))
					monkeys[result.nextMonkey].Enqueue(result.worryLevel)
				} else if j == 4 {
					result := task01Monkey4(value.(int))
					monkeys[result.nextMonkey].Enqueue(result.worryLevel)
				} else if j == 5 {
					result := task01Monkey5(value.(int))
					monkeys[result.nextMonkey].Enqueue(result.worryLevel)
				} else if j == 6 {
					result := task01Monkey6(value.(int))
					monkeys[result.nextMonkey].Enqueue(result.worryLevel)
				} else if j == 7 {
					result := task01Monkey7(value.(int))
					monkeys[result.nextMonkey].Enqueue(result.worryLevel)
				}

				value, moreItems = monkeys[j].Dequeue()
			}
		}
	}
	monkeyBusiness := GetMonkeyBusiness(monkeyInspections)
	return monkeyBusiness
}
