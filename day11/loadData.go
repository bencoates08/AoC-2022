package day11

func buildTestCallback(divisibleBy, trueMonkey, falseMonkey int) func(int) int {
	return func(val int) int {
		if val%divisibleBy == 0 {
			return trueMonkey
		}
		return falseMonkey
	}
}

func loadData() []Monkey {
	return []Monkey{
		// Monkey 0
		{
			Items:     []int{50, 70, 89, 75, 66, 66},
			Operation: func(old int) int { return old * 5 },
			Test:      buildTestCallback(2, 2, 1),
		},

		// Monkey 1
		{
			Items:     []int{85},
			Operation: func(old int) int { return old * old },
			Test:      buildTestCallback(7, 3, 6),
		},

		// Monkey 2
		{
			Items:     []int{66, 51, 71, 76, 58, 55, 58, 60},
			Operation: func(old int) int { return old + 1 },
			Test:      buildTestCallback(13, 1, 3),
		},

		// Monkey 3
		{
			Items:     []int{79, 52, 55, 51},
			Operation: func(old int) int { return old + 6 },
			Test:      buildTestCallback(3, 6, 4),
		},

		// Monkey 4
		{
			Items:     []int{69, 92},
			Operation: func(old int) int { return old * 17 },
			Test:      buildTestCallback(19, 7, 5),
		},

		// Monkey 5
		{
			Items:     []int{71, 76, 73, 98, 67, 79, 99},
			Operation: func(old int) int { return old + 8 },
			Test:      buildTestCallback(5, 0, 2),
		},

		// Monkey 6
		{
			Items:     []int{82, 76, 69, 69, 57},
			Operation: func(old int) int { return old + 7 },
			Test:      buildTestCallback(11, 7, 4),
		},

		// Monkey 7
		{
			Items:     []int{65, 79, 86},
			Operation: func(old int) int { return old + 5 },
			Test:      buildTestCallback(17, 5, 0),
		},
	}
}
