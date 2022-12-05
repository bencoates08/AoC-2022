package day2

import (
	"os"
)

func Part1() int {
	var gameOutcomeTable = map[string]map[string]string{
		// Rock
		"X": {"A": "D", "B": "L", "C": "W"},
		// Paper
		"Y": {"A": "W", "B": "D", "C": "L"},
		// Scissors
		"Z": {"A": "L", "B": "W", "C": "D"},
	}

	var points = map[string]int{
		"W": 6,
		"D": 3,
		"L": 0,

		"A": 1,
		"B": 2,
		"C": 3,
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	data, _ := os.ReadFile("day2/input.txt")

	c := 0
	total := 0

	for c < len(data)-1 {
		gameOutcome := gameOutcomeTable[string(data[c+2])][string(data[c])]
		total += points[gameOutcome] + points[string(data[c+2])]
		c += 4
	}

	return total
}

func Part2() int {
	var reverseGameOutcomeTable = map[string]map[string]string{
		// Lose
		"X": {"A": "C", "B": "A", "C": "B"},
		// Draw
		"Y": {"A": "A", "B": "B", "C": "C"},
		// Win
		"Z": {"A": "B", "B": "C", "C": "A"},
	}

	var points = map[string]int{
		"Z": 6,
		"Y": 3,
		"X": 0,

		"A": 1,
		"B": 2,
		"C": 3,
	}

	data, _ := os.ReadFile("day2/input.txt")

	c := 0
	total := 0

	for c < len(data)-1 {
		playersChoice := reverseGameOutcomeTable[string(data[c+2])][string(data[c])]
		total += points[string(data[c+2])] + points[playersChoice]
		c += 4
	}

	return total
}
