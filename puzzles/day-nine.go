package puzzles

import (
	"errors"
	"fmt"
	"sort"
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
	answer := 1

	largestBasinSizes := d.findThreeLargestBasinSizes()
	for _, size := range largestBasinSizes {
		answer *= size
	}

	return fmt.Sprintf("The three largest basin sizes multiplied is, %d", answer)
}

func (d DayNine) findThreeLargestBasinSizes() []int {
	lowPoints := d.findLowPoints()
	basinSizes := []int{}

	for _, point := range lowPoints {
		basinSizes = append(basinSizes, d.getSurroundingBasinSize(point))
	}

	sort.Ints(basinSizes)

	return basinSizes[len(basinSizes)-3:]
}

// getSurroundingBasinSize implements DFS-ish to find the surrounding basin for a point
// the problem definition states that each low point belongs to exactly one basin, separated by heights of 9
func (d DayNine) getSurroundingBasinSize(p point) int {
	stack := []coordinate{p.getCoordinate()}
	visited := map[coordinate]bool{}
	basinSize := 0

	for len(stack) > 0 {
		// Pop
		n := len(stack) - 1
		basinPoint := stack[n]
		stack = stack[:n]

		if visited[basinPoint] {
			continue
		}

		visited[basinPoint] = true
		basinSize++

		stack = append(stack, d.getUnvisitedNeighboursInBasin(basinPoint, visited)...)
	}
	return basinSize
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
	height := d.heightMap[y][x]

	for _, dif := range d.acceptedDirections {
		neighbour, err := d.safeGetCellValue(x+dif.x, y+dif.y)
		if err != nil {
			continue
		}
		if neighbour <= height {
			return false
		}
	}
	return true
}

// getNeighboursInBasin returns all unvisited neighbours of c where: the height is < 9 and the indexes are within bounds
func (d DayNine) getUnvisitedNeighboursInBasin(c coordinate, visited map[coordinate]bool) []coordinate {
	neighbours := []coordinate{}
	for _, dif := range d.acceptedDirections {
		adjacentCoord := coordinate{c.x + dif.x, c.y + dif.y}
		neighbour, err := d.safeGetCellValue(adjacentCoord.x, adjacentCoord.y)
		if err == nil && neighbour < 9 && !visited[adjacentCoord] {
			neighbours = append(neighbours, adjacentCoord)
		}
	}
	return neighbours
}

// safeGetCellValue returns the cell value or error if out-of-bounds
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

	// horizontally or vertically adjacent
	d.acceptedDirections = []coordinate{
		{-1, 0}, {0, -1}, {0, 1}, {1, 0},
	}
}
