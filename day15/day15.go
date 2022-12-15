package day15

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Grid [][]byte

type Location struct {
	x int
	y int
}

type Measurement struct {
	sensor Location
	beacon Location
}

func loadData(input string) []Measurement {
	readFile, _ := os.Open(input)
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	measurements := []Measurement{}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		line = strings.Replace(line, "Sensor at x=", "", 1)
		line = strings.Replace(line, " y=", "", 1)
		line = strings.Replace(line, " closest beacon is at x=", "", 1)
		line = strings.Replace(line, " y=", "", 1)

		data := strings.Split(line, ":")

		sensor := strings.Split(data[0], ",")
		beacon := strings.Split(data[1], ",")

		sensorX, _ := strconv.Atoi(sensor[0])
		sensorY, _ := strconv.Atoi(sensor[1])
		beaconX, _ := strconv.Atoi(beacon[0])
		beaconY, _ := strconv.Atoi(beacon[1])

		measurements = append(measurements, Measurement{Location{x: sensorX, y: sensorY}, Location{x: beaconX, y: beaconY}})
	}

	return measurements
}

func findGridDimensions(measurements []Measurement) (Location, Location) {
	minX := 0
	minY := 0
	maxX := 0
	maxY := 0

	for _, measurement := range measurements {
		distance := manhattanDistance(measurement.sensor, measurement.beacon)

		measurementXMin := measurement.sensor.x - distance
		measurementYMin := measurement.sensor.y - distance
		measurementXMax := measurement.sensor.x + distance
		measurementYmax := measurement.sensor.y + distance

		if measurementXMin < minX {
			minX = measurementXMin
		}
		if measurementYMin < minY {
			minY = measurementYMin
		}
		if measurementXMax > maxX {
			maxX = measurementXMax
		}
		if measurementYmax > maxY {
			maxY = measurementYmax
		}
	}

	return Location{x: minX, y: minY}, Location{x: maxX, y: maxY}
}

func manhattanDistance(a Location, b Location) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func occupiedOnLine(measurements []Measurement, lineStart, lineEnd Location) int {
	count := 0
	lineLength := manhattanDistance(lineStart, lineEnd)

	for i := 0; i < lineLength; i++ {
		location := Location{x: lineStart.x + i, y: lineStart.y}

		val := '.'

		for _, measurement := range measurements {
			sensorToBeacon := manhattanDistance(measurement.sensor, measurement.beacon)
			sensorToPos := manhattanDistance(measurement.sensor, location)

			if location.x == measurement.beacon.x && location.y == measurement.beacon.y {
				val = 'B'
				break
			}

			if sensorToPos <= sensorToBeacon {
				val = '#'
			}
		}

		if val == '#' {
			count++
		}
	}

	return count
}

func Part1(line int) int {
	measurments := loadData("day15/input.txt")

	min, max := findGridDimensions(measurments)
	lineStart := Location{x: min.x, y: line}
	lineEnd := Location{x: max.x, y: line}

	total := occupiedOnLine(measurments, lineStart, lineEnd)

	return total
}

// Part 2

func unoccupiedPos(measurements []Measurement, size int) Location {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			point := Location{x: i, y: j}
			isOccupied := false

			for _, measurement := range measurements {
				sensorToBeacon := manhattanDistance(measurement.sensor, measurement.beacon)
				sensorToPoint := manhattanDistance(measurement.sensor, point)

				if sensorToPoint <= sensorToBeacon {
					isOccupied = true

					pointToSensorHeight := abs(point.x - measurement.sensor.x)

					j = measurement.sensor.y + sensorToBeacon - pointToSensorHeight

					break
				}
			}

			if !isOccupied {
				return point
			}
		}
	}

	return Location{x: -1, y: -1}
}

func Part2() int {
	measurments := loadData("day15/input.txt")

	unoccupiedPos := unoccupiedPos(measurments, 4000000)

	return 4000000*unoccupiedPos.x + unoccupiedPos.y
}

func (g Grid) String() string {
	str := ""
	for _, row := range g {
		for _, col := range row {
			if col == 0 {
				str += ". "
			} else {
				str += string(col) + " "
			}
		}
		str += "\n"
	}

	return str
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}

	return x
}

// Previous code that was too slow

// func newGrid(min Location, max Location) Grid {
// 	grid := make(Grid, max.y-min.y+1)

// 	for i := range grid {
// 		grid[i] = make([]byte, max.x-min.x+1)
// 	}

// 	return grid
// }

// func (g Grid) addMeasurment(m Measurement) {
// 	g[m.sensor.y][m.sensor.x] = 'S'
// 	g[m.beacon.y][m.beacon.x] = 'B'

// 	distance := manhattanDistance(m.sensor, m.beacon)

// 	for i := 1; i <= 2*distance+1; i++ {
// 		y := m.sensor.y - distance + i - 1
// 		if y < 0 || y >= len(g) {
// 			continue
// 		}

// 		length := i
// 		if i > distance {
// 			length = 2*distance - i + 2
// 		}

// 		for j := 0; j < length; j++ {
// 			if j < 0 || j >= len(g[y]) {
// 				continue
// 			}

// 			for k := m.sensor.x - j; k < m.sensor.x+j+1; k++ {
// 				if k < 0 || k >= len(g[y]) {
// 					continue
// 				}

// 				if g[y][k] == 0 {
// 					g[y][k] = '#'
// 				}
// 			}
// 		}
// 	}
// }

// func (g Grid) addMeasurements(measurements []Measurement, zeroOffset Location) {
// 	for _, measurement := range measurements {
// 		measurement.sensor.x += zeroOffset.x
// 		measurement.sensor.y += zeroOffset.y
// 		measurement.beacon.x += zeroOffset.x
// 		measurement.beacon.y += zeroOffset.y

// 		g.addMeasurment(measurement)
// 	}
// }

// func (g Grid) occupiedOnLine(l int, offset Location) int {
// 	l += offset.y
// 	count := 0
// 	for _, col := range g[l] {
// 		if col == '#' {
// 			count++
// 		}
// 	}

// 	return count
// }
