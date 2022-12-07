package day6

import (
	"bufio"
	"os"
)

const (
	PACKET_SIZE_P1 = 4
	PACKET_SIZE_P2 = 14
)

func uniqueCharacters(str string) bool {
	for i := 0; i < len(str)-1; i++ {
		for j := i + 1; j < len(str); j++ {
			if str[i] == str[j] {
				return false
			}
		}
	}

	return true
}

func Part1(packetSize int) int {
	readFile, _ := os.Open("day6/input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanBytes)

	var group string
	var byteNum int

	for fileScanner.Scan() {
		char := fileScanner.Text()
		byteNum++
		if len(group) < packetSize {
			group += char
		} else {
			group = group[1:] + char
		}

		if uniqueCharacters(group) && len(group) == packetSize {
			return byteNum
		}
	}

	return 0
}

func Part2() int {
	return Part1(PACKET_SIZE_P2)
}
