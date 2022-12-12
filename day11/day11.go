package day11

const LCM = 9699690

type Monkey struct {
	Items       []int
	Operation   func(int) int
	Test        func(int) int
	Inspections int
}

func EvalMonkey(monkey int, monkeys []Monkey, worrying bool) []Monkey {
	for _, item := range monkeys[monkey].Items {
		monkeys[monkey].Inspections++

		// Calc new item value
		newItemVal := monkeys[monkey].Operation(item)
		if !worrying {
			newItemVal = newItemVal / 3
		} else if newItemVal >= LCM {
			newItemVal %= LCM
		}
		// Remove item from current monkey
		monkeys[monkey].Items = monkeys[monkey].Items[1:]

		// Determine new monkey
		newMonkey := monkeys[monkey].Test(newItemVal)
		// Add item to new monkey
		monkeys[newMonkey].Items = append(monkeys[newMonkey].Items, newItemVal)
	}

	return monkeys
}

func monkeyBusiness(monkeys []Monkey) int {
	var top1 int
	var top2 int

	for _, monkey := range monkeys {
		if int(monkey.Inspections) > top1 {
			top2 = top1
			top1 = int(monkey.Inspections)
		} else if int(monkey.Inspections) > top2 {
			top2 = int(monkey.Inspections)
		}
	}

	return top1 * top2
}

func Part1() int {
	monkeys := loadData()

	for round := 1; round <= 20; round++ {
		for j := range monkeys {
			monkeys = EvalMonkey(j, monkeys, false)
		}
	}

	return monkeyBusiness(monkeys)
}

func Part2() int {
	monkeys := loadData()

	for round := 1; round <= 10000; round++ {
		for j := range monkeys {
			monkeys = EvalMonkey(j, monkeys, true)
		}
	}

	return monkeyBusiness(monkeys)
}
