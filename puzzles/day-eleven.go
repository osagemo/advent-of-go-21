package puzzles

import (
	"fmt"
	"strconv"
)

type DayEleven struct {
	day
	grid octoGrid
}

type octoGrid struct {
	grid               [][]*int
	acceptedDirections []coordinate
	numFlashes         int
	firstFullFlash     int
}

type octopus struct {
	energyLevel *int
	coord       coordinate
}

func (g octoGrid) String() string {
	printString := ""
	for _, row := range g.grid {
		for i, n := range row {
			printString += fmt.Sprintf("%v", *n)
			if i != len(row)-1 {
				printString += ", "
			}
		}
		printString += "\n"
	}
	return printString
}

func (DayEleven) GetPuzzleName() string {
	return "Day 11: Dumbo Octopus"
}

func (d DayEleven) PartOne() string {
	d.grid.incrementEnergyLevels(100, false)
	return fmt.Sprintf("After 100 steps there have been %d flashes", d.grid.numFlashes)
}

func (d DayEleven) PartTwo() string {
	d.init()
	d.grid.incrementEnergyLevels(1000, true)
	return fmt.Sprintf("after %v steps, all octopi flashed", d.grid.firstFullFlash)
}

func (g *octoGrid) incrementEnergyLevels(numSteps int, stopAfterFull bool) {
	for i := 0; i < numSteps; i++ {
		for _, row := range g.grid {
			for _, n := range row {
				*n++ // increase all by 1 for every step
			}
		}
		numFlashes := g.triggerFlashes()
		g.resetFlashed()

		if numFlashes == len(g.grid)*len(g.grid[0]) && g.firstFullFlash == 0 {
			g.firstFullFlash = i + 1
			if stopAfterFull {
				return
			}
		}
		g.numFlashes += numFlashes
	}
}

func (g *octoGrid) triggerFlashes() int {
	flashed := map[coordinate]bool{}
	for y, row := range g.grid {
		for x, n := range row {
			if *n > 9 {
				coord := coordinate{x, y}
				g.triggerFlash(coord, flashed)
			}
		}
	}
	return len(flashed)
}

// triggerFlash adds 1 to all neighbouring values, and keeps triggering flashes if they go above 9
func (g *octoGrid) triggerFlash(coord coordinate, flashed map[coordinate]bool) {
	stack := []coordinate{coord}

	for len(stack) > 0 {
		// Pop
		n := len(stack) - 1
		octopus := stack[n]
		stack = stack[:n]

		if flashed[octopus] {
			continue
		}

		flashed[octopus] = true
		stack = append(stack, g.incrementNeighbours(octopus)...)
	}
}

// incrementNeighbours adds 1 to all neighbouring values from coordinate and returns all neighbours that will flash because of it
func (g *octoGrid) incrementNeighbours(coord coordinate) []coordinate {
	willFlash := []coordinate{}
	for _, dif := range g.acceptedDirections {
		adjacentCoord := coordinate{coord.x + dif.x, coord.y + dif.y}
		if g.outOfBounds(adjacentCoord) {
			continue
		}
		neighbour := g.grid[adjacentCoord.y][adjacentCoord.x]
		*neighbour++
		if *neighbour > 9 {
			willFlash = append(willFlash, adjacentCoord)
		}
	}
	return willFlash
}

func (g *octoGrid) resetFlashed() {
	for _, row := range g.grid {
		for _, n := range row {
			if *n > 9 {
				*n = 0
			}
		}
	}
}

func (g octoGrid) outOfBounds(coord coordinate) bool {
	return coord.x > len(g.grid[0])-1 || coord.x < 0 || coord.y > len(g.grid)-1 || coord.y < 0
}

func (d *DayEleven) init() {
	// all adjecent directions accepted
	acceptedDirections := []coordinate{
		{-1, 0}, {0, -1}, {1, 0}, {0, 1}, {1, 1}, {-1, -1}, {-1, 1}, {1, -1},
	}

	// init grid
	var grid = [][]*int{}

	for _, line := range d.inputLines {
		row := []*int{}
		for _, char := range line {
			s := string(char)
			n, _ := strconv.Atoi(s)
			row = append(row, &n)
		}
		grid = append(grid, row)
	}
	d.grid = octoGrid{grid, acceptedDirections, 0, 0}
}
