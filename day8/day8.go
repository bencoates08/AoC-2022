package day8

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Part 1

func totalTallestFromLeft(trees [][]int, visited [][]bool, total int) ([][]bool, int) {
	for y, treeRow := range trees {
		tallestSoFar := 0
		for x := range treeRow {
			if x == 0 || trees[y][x] > tallestSoFar {
				tallestSoFar = trees[y][x]
				if !visited[y][x] {
					visited[y][x] = true
					total++
				}
			}
		}
	}
	return visited, total
}

func totalTallestFromRight(trees [][]int, visited [][]bool, total int) ([][]bool, int) {
	for y, treeRow := range trees {
		tallestSoFar := 0
		for x := len(treeRow) - 1; x >= 0; x-- {
			if x == len(trees)-1 || trees[y][x] > tallestSoFar {
				tallestSoFar = trees[y][x]
				if !visited[y][x] {
					visited[y][x] = true
					total++
				}
			}
		}
	}
	return visited, total
}

func totalTallestFromTop(trees [][]int, visited [][]bool, total int) ([][]bool, int) {
	for x := range trees[0] {
		tallestSoFar := 0
		for y := range trees {
			if y == 0 || trees[y][x] > tallestSoFar {
				tallestSoFar = trees[y][x]
				if !visited[y][x] {
					visited[y][x] = true
					total++
				}
			}
		}
	}
	return visited, total
}

func totalTallestFromBottom(trees [][]int, visited [][]bool, total int) ([][]bool, int) {
	for x := range trees[0] {
		tallestSoFar := 0
		for y := len(trees) - 1; y >= 0; y-- {
			if y == len(trees)-1 || trees[y][x] > tallestSoFar {
				tallestSoFar = trees[y][x]
				if !visited[y][x] {
					visited[y][x] = true
					total++
				}
			}
		}
	}
	return visited, total
}

func Part1() int {
	readFile, _ := os.Open("day8/input.txt")
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	trees := [][]int{}

	for fileScanner.Scan() {
		treeRow := strings.Split(fileScanner.Text(), "")
		treeRowInt := []int{}

		for _, tree := range treeRow {
			t, _ := strconv.Atoi(tree)
			treeRowInt = append(treeRowInt, t)
		}

		trees = append(trees, treeRowInt)
	}

	total := 0
	visited := make([][]bool, len(trees))
	for i := range visited {
		visited[i] = make([]bool, len(trees[i]))
	}

	visited, total = totalTallestFromLeft(trees, visited, total)
	visited, total = totalTallestFromRight(trees, visited, total)
	visited, total = totalTallestFromTop(trees, visited, total)
	_, total = totalTallestFromBottom(trees, visited, total)

	return total
}

// Part 2

func treesVisivleFromTree(trees [][]int, x, y int) int {
	treesRight := 0
	for i := x + 1; i < len(trees); i++ {
		if trees[y][i] < trees[y][x] {
			treesRight++
		} else {
			treesRight++
			break
		}
	}

	treesLeft := 0
	for i := x - 1; i >= 0; i-- {
		if trees[y][i] < trees[y][x] {
			treesLeft++
		} else {
			treesLeft++
			break
		}
	}

	treesBottom := 0
	for i := y + 1; i < len(trees); i++ {
		if trees[i][x] < trees[y][x] {
			treesBottom++
		} else {
			treesBottom++
			break
		}
	}

	treesTop := 0
	for i := y - 1; i >= 0; i-- {
		if trees[i][x] < trees[y][x] {
			treesTop++
		} else {
			treesTop++
			break
		}
	}

	return treesRight * treesLeft * treesTop * treesBottom
}

func findNicestTrees(trees [][]int) (int, int) {
	max := 0
	var x, y int
	for j, treeRow := range trees {
		for i := range treeRow {
			treesVisible := treesVisivleFromTree(trees, j, i)
			if treesVisible > max {
				max = treesVisible
				x = j
				y = i
			}
		}
	}
	return x, y
}

func Part2() int {
	readFile, _ := os.Open("day8/input.txt")
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	trees := [][]int{}

	for fileScanner.Scan() {
		treeRow := strings.Split(fileScanner.Text(), "")
		treeRowInt := []int{}

		for _, tree := range treeRow {
			t, _ := strconv.Atoi(tree)
			treeRowInt = append(treeRowInt, t)
		}

		trees = append(trees, treeRowInt)
	}

	x, y := findNicestTrees(trees)
	bestScore := treesVisivleFromTree(trees, x, y)

	return bestScore
}
