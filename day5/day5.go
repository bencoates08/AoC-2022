package day5

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Table map[int][]string

func (t Table) moveTo(from, to int) {
	val := t[from][0]
	t[from] = t[from][1:]
	t[to] = append([]string{val}, t[to]...)
}

func (t Table) moveAmmountTo(ammount, from, to int) {
	for i := 0; i < ammount; i++ {
		t.moveTo(from, to)
	}
}

func (t Table) moveStackTo(ammount, from, to int) Table {
	crates := make([]string, len(t[from]))
	copy(crates, t[from])

	t[from] = t[from][ammount:]
	t[to] = append(crates[:ammount], t[to]...)

	return t
}

func extractMoves(line string) (ammount int, from int, to int) {
	words := strings.Split(line, " ")
	ammount, _ = strconv.Atoi(words[1])
	from, _ = strconv.Atoi(words[3])
	to, _ = strconv.Atoi(words[5])
	return
}

func Part1() map[int][]string {
	readFile, _ := os.Open("day5/input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	// Load table
	table := Table{}
	for i := 0; fileScanner.Scan() && i < 8; i++ {
		line := fileScanner.Text()

		for i := 0; i < len(line)/4+1; i++ {
			if string(line[i*4+1]) != " " {
				table[i+1] = append(table[i+1], string(line[i*4+1]))
			}
		}
	}

	// Execute moves
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" || line[0] == byte(' ') {
			continue
		}

		ammount, from, to := extractMoves(line)
		table.moveAmmountTo(ammount, from, to)
	}

	return table
}

func Part2() Table {
	readFile, _ := os.Open("day5/input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	// Load table
	table := Table{}
	for i := 0; fileScanner.Scan() && i < 8; i++ {
		line := fileScanner.Text()

		for i := 0; i < len(line)/4+1; i++ {
			if string(line[i*4+1]) != " " {
				table[i+1] = append(table[i+1], string(line[i*4+1]))
			}
		}
	}

	// Execute moves
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" || line[0] == byte(' ') {
			continue
		}

		ammount, from, to := extractMoves(line)
		table.moveStackTo(ammount, from, to)
	}

	return table
}
