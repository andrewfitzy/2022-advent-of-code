package day_11

import "github.com/emirpasic/gods/queues"

type monkey struct {
	number            int
	items             queues.Queue
	operation         func(int) int
	testDivisor       int
	trueTargetMonkey  int
	falseTargetMonkey int
}

func (m monkey) MonkeyPlay(value int) int {
	return m.operation(value)
}

func (m monkey) MonkeyBored(value int, limit ...int) result {
	reducedWorry := value / 3
	if len(limit) == 1 {
		reducedWorry = value % limit[0]
	}
	if reducedWorry%m.testDivisor == 0 {
		return result{
			nextMonkey: m.trueTargetMonkey,
			worryLevel: reducedWorry,
		}
	}
	return result{
		nextMonkey: m.falseTargetMonkey,
		worryLevel: reducedWorry,
	}
}
