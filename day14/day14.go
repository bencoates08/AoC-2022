package day14

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	X int
	Y int
}

type Grid [][]byte

func loadData() Grid {
	readFile, _ := os.Open("day14/input.txt")
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	grid := make(Grid, 1000)
	for i := range grid {
		grid[i] = make([]byte, 1000)
	}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		points := []Coord{}

		for _, s := range strings.Split(line, " -> ") {
			point := strings.Split(s, ",")
			col, _ := strconv.Atoi(point[0])
			row, _ := strconv.Atoi(point[1])

			if row > len(grid) || col > len(grid[0]) {
				panic("out of bounds of original grid")
			}

			points = append(points, Coord{row, col})
		}

		for i := 1; i < len(points); i++ {
			grid = grid.parseData(points[i-1], points[i])
		}
	}

	return grid
}

func (g Grid) parseData(a, b Coord) Grid {
	x := min(a.X, b.X)
	xDist := absDiff(a.X, b.X)
	y := min(a.Y, b.Y)
	yDist := absDiff(a.Y, b.Y)

	for i := x; i <= x+xDist; i++ {
		for j := y; j <= y+yDist; j++ {
			g[i][j] = '#'
		}
	}

	return g
}

func trimGrid(g Grid) (Grid, int) {
	top := 0
	bottom := 0
	left := 0
	right := 0

	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			if g[i][j] == '#' {
				if bottom < i {
					bottom = i
				}

				if left == 0 || left > j {
					left = j
				}

				if right < j {
					right = j
				}
			}
		}
	}

	newGrid := make(Grid, bottom-top+1)
	for i := range newGrid {
		newGrid[i] = make([]byte, right-left+1)
	}

	for i := top; i <= bottom; i++ {
		for j := left; j <= right; j++ {
			newGrid[i-top][j-left] = g[i][j]
		}
	}

	start := 500 - left

	return newGrid, start
}

func addSand(g Grid, start int) (Grid, bool) {
	sandPos := Coord{0, start}

	if g[sandPos.X][sandPos.Y] == '#' || g[sandPos.X][sandPos.Y] == 'o' {
		return g, true
	}

	for {
		if sandPos.X == len(g)-1 {
			return g, true
		}

		below := g[sandPos.X+1][sandPos.Y]

		if below == '#' || below == 'o' {
			if sandPos.Y == 0 {
				return g, true
			}

			leftDiag := g[sandPos.X+1][sandPos.Y-1]

			if leftDiag == '#' || leftDiag == 'o' {
				rightDiag := g[sandPos.X+1][sandPos.Y+1]

				if rightDiag == '#' || rightDiag == 'o' {
					g[sandPos.X][sandPos.Y] = 'o'
					break
				}

				sandPos.Y++
				sandPos.X++
				continue
			}

			sandPos.Y--
			sandPos.X++
			continue
		}

		sandPos.X++
	}

	return g, false
}

func fillSand(g Grid, start int) (Grid, int) {
	total := 0
	var done bool
	for {
		g, done = addSand(g, start)
		total++
		if done {
			total--
			break
		}
	}

	return g, total
}

func Part1() int {
	grid, start := trimGrid(loadData())

	total := 0

	grid, total = fillSand(grid, start)
	grid.print()

	return total
}

// Part 2

func loadDataP2() Grid {
	readFile, _ := os.Open("day14/input.txt")
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	grid := make(Grid, 1000)
	for i := range grid {
		grid[i] = make([]byte, 1000)
	}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		points := []Coord{}

		for _, s := range strings.Split(line, " -> ") {
			point := strings.Split(s, ",")
			col, _ := strconv.Atoi(point[0])
			row, _ := strconv.Atoi(point[1])

			if row > len(grid) || col > len(grid[0]) {
				panic("out of bounds of original grid")
			}

			points = append(points, Coord{row, col})
		}

		for i := 1; i < len(points); i++ {
			grid = grid.parseData(points[i-1], points[i])
		}
	}

	bottom := findBottom(grid) + 2
	grid = grid.parseData(Coord{bottom, 0}, Coord{bottom + 1, len(grid[0]) - 1})

	return grid
}

func findBottom(grid Grid) int {
	for i := len(grid) - 1; i >= 0; i-- {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '#' {
				return i
			}
		}
	}
	return 0
}

func Part2() int {
	grid, start := trimGrid(loadDataP2())

	total := 0

	_, total = fillSand(grid, start)
	// grid.print()

	return total
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func absDiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func (g Grid) print() {
	for _, row := range g {
		for _, col := range row {
			if col == 0 {
				fmt.Print(".")
				continue
			}
			fmt.Print(string(col) + "")
		}
		fmt.Println()
	}
	fmt.Println()
}
