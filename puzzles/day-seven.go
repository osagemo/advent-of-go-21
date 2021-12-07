package puzzles

import "strings"

type DaySeven struct {
	day
	crabs []int
}

func (DaySeven) GetPuzzleName() string {
	return "Day 7: The Treachery of Whales"
}

func (d *DaySeven) PartOne() string {
	return "Not implemented"
}

func (d *DaySeven) PartTwo() string {
	return "Not implemented"
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
