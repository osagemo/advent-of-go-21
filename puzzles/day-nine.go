package puzzles

import (
	"errors"
	"fmt"
	"strconv"
)

type DayNine struct {
	day
	heightMap          [][]int
	rows               int
	cols               int
	acceptedDirections []coordinate
}

type point struct {
	x      int
	y      int
	height int
}

func (p point) getCoordinate() coordinate {
	return coordinate{p.x, p.y}
}

func (DayNine) GetPuzzleName() string {
	return "Day 9: Smoke Basin"
}

func (d DayNine) PartOne() string {
	lowPoints := d.findLowPoints()
	sumOfRiskLevels := 0

	for _, point := range lowPoints {
		sumOfRiskLevels += (point.height + 1)
	}
	return fmt.Sprintf("The sum of risk levels for all low points is: %d", sumOfRiskLevels)
}

func (d DayNine) PartTwo() string {
	// lowPoints := d.findLowPoints()
	// basinSizes := []int{}
	// coordToVisited := map[coordinate]bool{}

	// for _, point := range lowPoints {
	// 	if coordToVisited[coordinate{point.x, point.y}] {
	// 		continue
	// 	}
	// 	basinSize := d.getSurroundingBasinSize(point, &coordToVisited)
	// }

	basinSize := d.getSurroundingBasinSize(point{10, 1, 0})

	return fmt.Sprintf("not implemented, %d", basinSize)
}

func (d DayNine) getSurroundingBasinSize(p point) int {
	// look in all directions, keep adding to pointsInBasin until we find edges (next is 9)
	nMap := map[coordinate]bool{p.getCoordinate(): true}
	nMap = d.getNeighboursRecursive(p.getCoordinate(), nMap)
	return len(nMap)
}

// rename
func (d DayNine) getNeighboursRecursive(c coordinate, knownNeighbours map[coordinate]bool) map[coordinate]bool {
	neighbours := d.getNeighboursInBasin(c)

	for _, coord := range neighbours {
		if knownNeighbours[coord] {
			continue
		}
		knownNeighbours[coord] = true
		neighbours = append(neighbours, d.getNeighboursInBasin(coord)...)

		// foundNeighbours := d.getNeighboursInBasin(coord)
		// for _, newNeighbour := range foundNeighbours {
		// 	if !knownNeighbours[newNeighbour] {
		// 		found = true
		// 		neighbours = append(neighbours, newNeighbour)
		// 		knownNeighbours[newNeighbour] = true
		// 	}
		// }
	}
	return knownNeighbours
}

func (d DayNine) findLowPoints() []point {
	lowPoints := []point{}
	for y, row := range d.heightMap {
		for x, height := range row {
			if d.isLowPoint(x, y) {
				lowPoints = append(lowPoints, point{x, y, height})
			}
		}
	}
	return lowPoints
}

func (d DayNine) isLowPoint(x int, y int) bool {
	// adjacent
	value := d.heightMap[y][x]

	for _, dif := range d.acceptedDirections {
		neighbour, err := d.safeGetCellValue(x+dif.x, y+dif.y)
		if err != nil {
			continue
		}
		if neighbour <= value {
			return false
		}
	}
	return true
}

// getNeighboursInBasin returns all neighbours of c where the value is < 9
func (d DayNine) getNeighboursInBasin(c coordinate) []coordinate {
	neighbours := []coordinate{}
	for _, dif := range d.acceptedDirections {
		neighbour, err := d.safeGetCellValue(c.x+dif.x, c.y+dif.y)
		if err == nil && neighbour < 9 {
			neighbours = append(neighbours, coordinate{c.x + dif.x, c.y + dif.y})
		}
	}
	return neighbours
}

// returns -1 for out-of-bounds
func (d DayNine) safeGetCellValue(x int, y int) (int, error) {
	if x > d.cols-1 || x < 0 || y > d.rows-1 || y < 0 {
		return -1, errors.New("out of bounds")
	}
	return d.heightMap[y][x], nil
}

func (d *DayNine) init() {
	for _, line := range d.input.Lines {
		row := []int{}
		for _, char := range line {
			n, _ := strconv.Atoi(string(char))
			row = append(row, n)
		}
		d.heightMap = append(d.heightMap, row)
	}
	d.rows = len(d.heightMap)
	d.cols = len(d.heightMap[0])

	d.acceptedDirections = []coordinate{
		{-1, 0}, {0, -1}, {0, 1}, {1, 0},
	}
}
