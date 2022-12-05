package day4

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type elf struct {
	start int
	end   int
}

func elfRange(s string) elf {
	vals := strings.Split(s, "-")
	start, _ := strconv.Atoi(vals[0])
	end, _ := strconv.Atoi(vals[1])
	return elf{start: start, end: end}
}

func isElfInRange(elf1, elf2 elf) bool {
	return elf1.start <= elf2.start && elf1.end >= elf2.end
}

func isElfOverlapping(elf1, elf2 elf) bool {
	return elf1.start <= elf2.start && elf1.end >= elf2.start ||
		elf1.start <= elf2.end && elf1.end >= elf2.end
}

func Part1() int {
	readFile, _ := os.Open("day4/input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	total := 0

	for fileScanner.Scan() {
		elves := strings.Split(fileScanner.Text(), ",")

		elf1 := elfRange(elves[0])
		elf2 := elfRange(elves[1])

		if isElfInRange(elf1, elf2) || isElfInRange(elf2, elf1) {
			total++
		}
	}

	return total
}

func Part2() int {
	readFile, _ := os.Open("day4/input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	total := 0

	for fileScanner.Scan() {
		elves := strings.Split(fileScanner.Text(), ",")

		elf1 := elfRange(elves[0])
		elf2 := elfRange(elves[1])

		if isElfOverlapping(elf1, elf2) || isElfOverlapping(elf2, elf1) {
			total++
		}
	}

	return total
}
