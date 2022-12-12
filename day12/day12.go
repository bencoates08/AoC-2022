package day12

import (
	"bufio"
	"math"
	"os"
	"strings"
)

type Position struct {
	x        int
	y        int
	height   int
	distance int
	visited  bool
}

func loadData() ([][]Position, Position, Position) {
	readFile, _ := os.Open("day12/input.txt")
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var start Position
	var end Position

	grid := [][]Position{}
	i := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		strings.Split(line, "")

		grid = append(grid, []Position{})
		for j, char := range line {
			if char == 'S' {
				start = Position{i, j, int(rune('a') - 97), 0, true}
				grid[i] = append(grid[i], start)
			} else if char == 'E' {
				end = Position{i, j, int(rune('z') - 97), math.MaxInt32, false}
				grid[i] = append(grid[i], end)
			} else {
				grid[i] = append(grid[i], Position{i, j, int(rune(char) - 97), math.MaxInt32, false})
			}
		}
		i++
	}

	return grid, start, end
}

func getNeighbours(p Position, g [][]Position) []Position {
	neighbours := []Position{}

	if p.x > 0 {
		neighbours = append(neighbours, g[p.x-1][p.y])
	}
	if p.x < len(g)-1 {
		neighbours = append(neighbours, g[p.x+1][p.y])
	}
	if p.y > 0 {
		neighbours = append(neighbours, g[p.x][p.y-1])
	}
	if p.y < len(g[0])-1 {
		neighbours = append(neighbours, g[p.x][p.y+1])
	}

	return neighbours
}

func getNextPosition(g [][]Position) Position {
	var pos Position
	for _, row := range g {
		for _, p := range row {
			if !p.visited && (pos.distance == 0 || p.distance < pos.distance) {
				pos = p
			}
		}
	}

	return pos
}

func getShortestDistance(grid [][]Position, start, end Position) int {
	currentPosition := start

	for {
		if currentPosition.x == end.x && currentPosition.y == end.y {
			return currentPosition.distance
		}

		neighbours := getNeighbours(currentPosition, grid)

		for _, neighbour := range neighbours {
			if neighbour.visited || neighbour.height > currentPosition.height+1 {
				continue
			}

			if neighbour.distance > currentPosition.distance + +1 {
				grid[neighbour.x][neighbour.y].distance = currentPosition.distance + 1
			}
		}

		grid[currentPosition.x][currentPosition.y].visited = true
		currentPosition = getNextPosition(grid)
	}
}

func Part1() int {
	grid, start, end := loadData()

	return getShortestDistance(grid, start, end)
}

func getLowestPositions(g [][]Position) []Position {
	var lowest []Position
	for _, row := range g {
		for _, p := range row {
			if p.height == 0 {
				lowest = append(lowest, p)
			}
		}
	}

	return lowest
}

func Part2() int {
	grid, start, _ := loadData()
	grid[start.x][start.y] = Position{start.x, start.y, 0, math.MaxInt32, false}

	startingPositions := getLowestPositions(grid)

	var shortestDistance int
	for _, startingPosition := range startingPositions {
		grid, start, end := loadData()
		grid[start.x][start.y] = Position{start.x, start.y, 0, math.MaxInt32, false}

		startingPosition.distance = 0
		startingPosition.visited = true
		grid[startingPosition.x][startingPosition.y] = startingPosition

		distance := getShortestDistance(grid, startingPosition, end)
		if shortestDistance == 0 || distance < shortestDistance {
			shortestDistance = distance
		}
	}

	return shortestDistance
}
