package day13

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Data struct {
	Left  []interface{}
	Right []interface{}
}

func compare(left, right interface{}) int {
	switch l := left.(type) {
	case int:
		switch r := right.(type) {
		case int:
			return compareInts(l, r)
		case []interface{}:
			return compare([]interface{}{l}, r)
		}

	case []interface{}:
		switch r := right.(type) {
		case int:
			return compare(l, []interface{}{r})
		case []interface{}:
			if len(l) == 0 && len(r) == 0 {
				return 0
			}

			if len(l) == 0 {
				return -1
			}

			if len(r) == 0 {
				return 1
			}

			for i := 0; i < len(l) && i < len(r); i++ {
				comp := compare(l[i], r[i])
				if comp != 0 {
					return comp
				}
			}

			return compareInts(len(l), len(r))
		}
	}

	return -1
}

func compareInts(left, right int) int {
	if left < right {
		return -1
	}
	if left > right {
		return 1
	}
	return 0
}

func loadData() []Data {
	readFile, _ := os.Open("day13/input.txt")
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	data := []Data{}

	for fileScanner.Scan() {
		leftStr := fileScanner.Text()
		fileScanner.Scan()
		rightStr := fileScanner.Text()

		left := parseData(leftStr)[0].([]interface{})
		right := parseData(rightStr)[0].([]interface{})

		data = append(data, Data{left, right})

		fileScanner.Scan()
	}

	return data
}

func parseData(s string) []interface{} {
	data := []interface{}{}

	for i := 0; i < len(s); i++ {
		if s[i] == '[' {
			subString := ""
			leftBracketCount := 1

			for j := i + 1; j < len(s); j++ {
				if s[j] == '[' {
					leftBracketCount++
				}

				if s[j] == ']' {
					leftBracketCount--
				}

				if leftBracketCount == 0 {
					break
				}

				subString += string(s[j])
			}

			data = append(data, parseData(subString))

			i += len(subString) + 1

		} else if s[i] == ',' {
			continue

		} else {
			val := 0

			for j := i; j < len(s); j++ {
				if s[j] == ',' {
					break
				}
				val, _ = strconv.Atoi(string(s[i : j+1]))
			}
			data = append(data, val)
		}
	}

	return data
}

func Part1() int {
	data := loadData()

	total := 0
	for i, d := range data {
		comp := compare(d.Left, d.Right)
		if comp == -1 {
			total += i + 1
		}
	}

	return total
}

func loadDataP2() []interface{} {
	readFile, _ := os.Open("day13/input.txt")
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	data := []interface{}{}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			continue
		}

		lineData := parseData(line)[0].([]interface{})

		data = append(data, lineData)
	}

	return data
}

func Part2() int {
	data := loadDataP2()

	divPac1 := []interface{}{[]interface{}{2}}
	divPac2 := []interface{}{[]interface{}{6}}

	data = append(data, divPac1)
	data = append(data, divPac2)

	sort.Slice(data, func(i, j int) bool {
		return compare(data[i], data[j]) == -1
	})

	a := 0
	b := 0
	for i, d := range data {
		if fmt.Sprintf("%v", d) == "[[2]]" {
			a = i + 1
		}
		if fmt.Sprintf("%v", d) == "[[6]]" {
			b = i + 1
		}
	}

	return a * b
}
