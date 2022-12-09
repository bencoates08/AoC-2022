package day9

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	x int
	y int
}

type Trail map[int]map[int]bool

type Rope map[int]Coord

func mapTailTrail(coord Coord, trail Trail) Trail {
	if trail[coord.x] == nil {
		trail[coord.x] = make(map[int]bool)
	}

	trail[coord.x][coord.y] = true
	return trail
}

func numTrailPoints(trail Trail) int {
	numPoints := 0
	for _, yMap := range trail {
		numPoints += len(yMap)
	}
	return numPoints
}

func NewRope(len int) Rope {
	rope := make(Rope)
	for i := 0; i < len; i++ {
		rope[i] = Coord{x: 0, y: 0}
	}
	return rope
}

func moveHead(rope Rope, dir string) Rope {
	switch dir {
	case "U":
		rope[0] = Coord{x: rope[0].x, y: rope[0].y + 1}
	case "D":
		rope[0] = Coord{x: rope[0].x, y: rope[0].y - 1}
	case "R":
		rope[0] = Coord{x: rope[0].x + 1, y: rope[0].y}
	case "L":
		rope[0] = Coord{x: rope[0].x - 1, y: rope[0].y}
	default:
		panic("Invalid direction")
	}

	return rope
}

func move1Step(rope Rope, dir string) Rope {
	rope = moveHead(rope, dir)

	for i := 1; i < len(rope); i++ {
		xDiff := rope[i-1].x - rope[i].x
		yDiff := rope[i-1].y - rope[i].y

		if abs(xDiff) > 1 || abs(yDiff) > 1 {
			var newX, newY int

			if xDiff == 0 {
				newX = rope[i].x
			} else if xDiff > 0 {
				newX = rope[i].x + 1
			} else {
				newX = rope[i].x - 1
			}

			if yDiff == 0 {
				newY = rope[i].y
			} else if yDiff > 0 {
				newY = rope[i].y + 1
			} else {
				newY = rope[i].y - 1
			}

			rope[i] = Coord{newX, newY}
		}
	}

	return rope
}

func move(rope Rope, trail Trail, dir string, dist int) (Rope, Trail) {
	for i := 0; i < dist; i++ {
		rope = move1Step(rope, dir)
		trail = mapTailTrail(rope[len(rope)-1], trail)
	}

	return rope, trail
}

func Part1() int {
	readFile, _ := os.Open("day9/input.txt")
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	rope := NewRope(2)
	trail := make(Trail)

	for fileScanner.Scan() {
		line := strings.Split(fileScanner.Text(), " ")
		dir := line[0]
		dist, _ := strconv.Atoi(line[1])

		rope, trail = move(rope, trail, dir, dist)
	}

	return numTrailPoints(trail)
}

func Part2() int {
	readFile, _ := os.Open("day9/input.txt")
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	rope := NewRope(10)
	trail := make(Trail)

	for fileScanner.Scan() {
		line := strings.Split(fileScanner.Text(), " ")
		dir := line[0]
		dist, _ := strconv.Atoi(line[1])

		rope, trail = move(rope, trail, dir, dist)
	}

	return numTrailPoints(trail)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Better solution but does not match the way the example moves

// func move1Step(rope Rope, dir string) (Rope) {
// 	ogPos := rope[0]
// 	rope = moveHead(rope, dir)

// 	for i:=1; i < len(rope); i++ {
// 		xDiff := abs(rope[i-1].x - rope[i].x)
// 		yDiff := abs(rope[i-1].y - rope[i].y)

// 		newOGPos := rope[i]

// 		if xDiff > 1 || yDiff > 1 {
// 			x2Diff := abs(rope[i-1].x - rope[i+1].x)
// 			y2Diff := abs(rope[i-1].y - rope[i+1].y)

// 			if x2Diff + y2Diff > 2 {
// 				if x2Diff > 1 {
// 					midXPos := rope[i-1].x + ((rope[i+1].x - rope[i-1].x) / 2)
// 					rope[i] = Coord{x: midXPos, y: rope[i].y}
// 				}
// 				if y2Diff > 1 {
// 					midYPos := rope[i-1].y + ((rope[i+1].y - rope[i-1].y) / 2)
// 					rope[i] = Coord{x: rope[i].x, y: midYPos}
// 				}
// 			} else {
// 				rope[i] = ogPos
// 			}
// 		}

// 		ogPos = newOGPos
// 	}

// 	return rope
// }
