package day10

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Part1() int {
	readFile, _ := os.Open("day10/input.txt")
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	signalStrength := []int{}
	r0 := 1
	cycle := 1

	for fileScanner.Scan() {
		words := strings.Split(fileScanner.Text(), " ")

		if (cycle+20)%40 == 0 {
			signalStrength = append(signalStrength, r0)
		}

		if words[0] != "noop" {
			cycle++

			if (cycle+20)%40 == 0 {
				signalStrength = append(signalStrength, r0)
			}

			val, _ := strconv.Atoi(words[1])
			r0 += val

			cycle++

			continue
		}

		cycle++
	}

	sigStrengthTotal := 0
	for i, val := range signalStrength {
		sigStrengthTotal += val * ((i+1)*40 - 20)
	}

	return sigStrengthTotal
}

func addToLine(line string, r0, cycle int) string {

	if cycle-r0 <= 1 && cycle-r0 >= -1 {
		line = line + "#"
	} else {
		line = line + "."
	}

	return line
}

func Part2() string {
	readFile, _ := os.Open("day10/input.txt")
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	r0 := 1
	position := 0
	lines := []string{}
	line := ""

	for fileScanner.Scan() {
		words := strings.Split(fileScanner.Text(), " ")

		if position == 40 {
			lines = append(lines, line)
			line = ""
			position = 0
		}

		line = addToLine(line, r0, position)

		if words[0] != "noop" {
			position++

			if position == 40 {
				lines = append(lines, line)
				line = ""
				position = 0
			}

			line = addToLine(line, r0, position)

			val, _ := strconv.Atoi(words[1])
			r0 += val

			position++

			continue
		}

		position++
	}
	lines = append(lines, line)

	result := ""
	for _, line := range lines {
		result += line + "\n"
	}

	return result
}
