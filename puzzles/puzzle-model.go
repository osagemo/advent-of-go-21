package puzzles

import "github.com/osagemo/advent-of-go-21/puzzles/input"

type day struct {
	day   int
	input input.Input
}

type Day interface {
	init()
	GetPuzzleName() string
	PartOne() string
	PartTwo() string
}

func NewDay(dayNum int) Day {
	base := day{dayNum, input.GetInput(dayNum)}
	var day Day
	// find some better way to map these
	switch dayNum {
	case 1:
		dayOne := new(DayOne)
		dayOne.base = base
		dayOne.init()
		day = dayOne
	case 2:
		dayTwo := new(DayTwo)
		dayTwo.base = base
		day = dayTwo
	case 3:
		dayThree := new(DayThree)
		dayThree.base = base
		dayThree.init()
		day = dayThree
	}
	return day
}
