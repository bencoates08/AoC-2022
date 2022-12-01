package day1

import (
	"os"
	"strconv"
)

func Part1() int {
	data, _ := os.ReadFile("day1/input.txt")
	dataLength := len(data)

	c := 0

	var max int
	for c < dataLength-1 {

		var total int
		for c < dataLength-1 {
			if data[c] == 10 {
				break
			}

			var calorie []byte
			for c < dataLength-1 {
				if data[c] == 10 {
					break
				}

				calorie = append(calorie, data[c])
				c++
			}

			calVal, _ := strconv.Atoi(string(calorie))
			total = total + calVal
			c++
		}

		if total > max {
			max = total
		}

		c++
	}

	return max
}

func Part2() int {
	data, _ := os.ReadFile("day1/input.txt")
	dataLength := len(data)

	c := 0

	var top3 [3]int
	for c < dataLength-1 {

		var total int
		for c < dataLength-1 {
			if data[c] == 10 {
				break
			}

			var calorie []byte
			for c < dataLength-1 {
				if data[c] == 10 {
					break
				}

				calorie = append(calorie, data[c])
				c++
			}

			calVal, _ := strconv.Atoi(string(calorie))
			total = total + calVal
			c++
		}

		smallestIndex := getSmallestIndex(top3)

		if total > top3[smallestIndex] {
			top3[smallestIndex] = total
		}

		c++
	}

	var top3Total int
	for _, v := range top3 {
		top3Total = top3Total + v
	}

	return top3Total
}

func getSmallestIndex (arr [3]int) int {
	var smallestIndex int
	for i, v := range arr {
		if i == 0 {
			smallestIndex = i
		}

		if v < arr[smallestIndex] {
			smallestIndex = i
		}
	}

	return smallestIndex
}