package puzzles

import (
	"fmt"
	"strconv"
)

type DayEleven struct {
	day
	grid octoGrid
}

func (DayEleven) GetPuzzleName() string {
	return "Day 11: Dumbo Octopus"
}

type octoGrid struct {
	grid               [][]*int
	acceptedDirections []coordinate
	numFlashes         int
	firstFullFlash     int
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

func (d DayEleven) PartOne() string {
	d.grid.incrementEnergyLevels(100, false)
	return fmt.Sprintf("After 100 steps there have been %d flashes", d.grid.numFlashes)
}

func (d DayEleven) PartTwo() string {
	d.init() // reset pointers
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
	coordinatesToFlash := []coordinate{coord}

	for len(coordinatesToFlash) > 0 {
		flashCandidate := coordinatesToFlash[0]
		coordinatesToFlash = coordinatesToFlash[1:]

		if flashed[flashCandidate] {
			continue
		}

		flashed[flashCandidate] = true
		coordinatesToFlash = append(coordinatesToFlash, g.incrementNeighbours(flashCandidate)...)
	}
}

// incrementNeighbours adds 1 to all neighbouring values from coordinate and returns all neighbours that will flash because of it
func (g *octoGrid) incrementNeighbours(coord coordinate) []coordinate {
	newFlashCandidates := []coordinate{}
	for _, dif := range g.acceptedDirections {
		adjacentCoord := coordinate{coord.x + dif.x, coord.y + dif.y}
		if g.outOfBounds(adjacentCoord) {
			continue
		}

		neighbour := g.grid[adjacentCoord.y][adjacentCoord.x]
		*neighbour++ // increment
		if *neighbour > 9 {
			newFlashCandidates = append(newFlashCandidates, adjacentCoord)
		}
	}
	return newFlashCandidates
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
