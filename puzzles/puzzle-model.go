package puzzles

import (
	"time"

	"github.com/osagemo/advent-of-go-21/puzzles/input"
)

type day struct {
	day   int
	input input.Input
	start time.Time
}

type Day interface {
	init()
	GetPuzzleName() string
	GetStart() time.Time
	SetStart(time.Time)
	PartOne() string
	PartTwo() string
}

// improvements:
// --Separate days into packages instead, then we can avoid using methods for everything due to fear of polluting package scope
func NewDay(dayNum int) Day {
	base := day{dayNum, input.GetInput(dayNum), time.Time{}}
	var day Day
	// find some better way to map these
	switch dayNum {
	case 1:
		dayOne := new(DayOne)
		dayOne.day = base
		day = dayOne
	case 2:
		dayTwo := new(DayTwo)
		dayTwo.day = base
		day = dayTwo
	case 3:
		dayThree := new(DayThree)
		dayThree.day = base
		day = dayThree
	case 4:
		dayFour := new(DayFour)
		dayFour.day = base
		day = dayFour
	case 5:
		dayFive := new(DayFive)
		dayFive.day = base
		day = dayFive
	case 6:
		daySix := new(DaySix)
		daySix.day = base
		day = daySix
	case 7:
		daySeven := new(DaySeven)
		daySeven.day = base
		day = daySeven
	case 8:
		dayEight := new(DayEight)
		dayEight.day = base
		day = dayEight
	}

	day.SetStart(time.Now())
	day.init()

	return day
}

func (d *day) GetStart() time.Time {
	return d.start
}

func (d *day) SetStart(time time.Time) {
	d.start = time
}
