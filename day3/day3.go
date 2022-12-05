package day3

import (
	"bufio"
	"os"
)

func getPriority(b byte) int {
	val := rune(b) - rune('A') + 1

	if val < 32 {
		return int(val) + 26
	} else {
		return int(val) - 32
	}
}

func getStringSet(s string) map[string]bool {
	m := map[string]bool{}
	for _, v := range s {
		m[string(v)] = true
	}
	return m
}

func getSetIntersectionByte(a, b map[string]bool) byte {
	for k := range a {
		if _, ok := b[k]; ok {
			return byte(k[0])
		}
	}
	return byte(0)
}

func getSetIntersection(a, b map[string]bool) map[string]bool {
	m := map[string]bool{}
	for k := range a {
		if _, ok := b[k]; ok {
			m[k] = true
		}
	}
	return m
}

func Part1() int {
	readFile, _ := os.Open("day3/input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	totalPriority := 0

	for fileScanner.Scan() {
		bags := fileScanner.Text()
		containerSize := len(bags) / 2

		set1 := getStringSet(bags[:containerSize])
		set2 := getStringSet(bags[containerSize:])
		intersect := getSetIntersectionByte(set1, set2)

		totalPriority += getPriority(intersect)
	}

	readFile.Close()

	return totalPriority
}

func Part2() int {
	readFile, _ := os.Open("day3/input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	totalPriority := 0

	for fileScanner.Scan() {
		bags := map[int]map[string]bool{1: getStringSet(fileScanner.Text())}
		for i := 2; i <= 3; i++ {
			fileScanner.Scan()
			bags[i] = getStringSet(fileScanner.Text())
		}

		setInt1 := getSetIntersection(bags[1], bags[2])
		setInt := getSetIntersection(setInt1, bags[3])

		for k := range setInt {
			totalPriority += getPriority(k[0])
		}
	}

	readFile.Close()

	return totalPriority
}
