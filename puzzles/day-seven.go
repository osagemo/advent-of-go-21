package puzzles

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

type DaySeven struct {
	day
	crabs []int
}

func (DaySeven) GetPuzzleName() string {
	return "Day 7: The Treachery of Whales"
}

func (d *DaySeven) PartOne() string {
	minCost, pos := d.findMiniumumFuelCostAndPosition(absDiffInt)
	return fmt.Sprintf("Desired position: %d, would consume %d fuel", pos, minCost)
}

func (d *DaySeven) PartTwo() string {
	minCost, pos := d.findMiniumumFuelCostAndPosition(crabEngineFuelRate)
	return fmt.Sprintf("Desired position: %d, would consume %d fuel", pos, minCost)
}

func (d DaySeven) findMiniumumFuelCostAndPosition(costFunction func(int, int) int) (int, int) {
	// calculate cost for all positions, return min
	first, last := d.firstAndLastCrabPosition()

	minCost := math.MaxInt
	posOfMinCost := -1
	for pos := first; pos <= last; pos++ {
		cost := d.costForAllCrabsToPosition(pos, costFunction)
		if cost < minCost {
			minCost = cost
			posOfMinCost = pos
		}
	}
	return minCost, posOfMinCost
}

func (d DaySeven) firstAndLastCrabPosition() (first int, last int) {
	sort.Ints(d.crabs)
	return d.crabs[0], d.crabs[len(d.crabs)-1]
}

func (d DaySeven) costForAllCrabsToPosition(destination int, costFunction func(int, int) int) int {
	var cost int
	for _, pos := range d.crabs {
		cost += costFunction(pos, destination)
	}

	return cost
}

func (d *DaySeven) init() {
	crabs := d.input.Lines[0]
	if len(crabs) <= 0 {
		panic("no crabs")
	}

	for _, s := range strings.Split(crabs, ",") {
		d.crabs = append(d.crabs, parseInt(s))
	}
}

func crabEngineFuelRate(firstPos int, secondPos int) int {
	dif := absDiffInt(firstPos, secondPos)

	return dif * (dif + 1) / 2 // sum of first n (dif) integers
}
