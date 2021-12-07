package puzzles

import (
	"fmt"
	"strings"
)

type DaySix struct {
	day
	initialFishes []*fish
}

type fish struct {
	timer int
}

func (f *fish) age() bool {
	spawned := false
	if f.timer == 0 {
		f.timer = 6
		spawned = true
	} else {
		f.timer--
	}

	return spawned
}

func (DaySix) newFish() fish {
	fish := fish{8}
	return fish
}

func (DaySix) GetPuzzleName() string {
	return "Day 6: Lanternfish"
}

func (d *DaySix) PartOne() string {
	fishes := d.initialFishes
	for day := 0; day < 80; day++ {
		for _, fish := range fishes {
			spawned := fish.age()
			if spawned {
				newFish := d.newFish()
				fishes = append(fishes, &newFish)
			}
		}
	}

	return fmt.Sprintf("After 80 days there are %d lanternfishies", len(fishes))
}

func (d *DaySix) PartTwo() string {
	d.init()
	ageCount := make(map[int]int)
	// init
	for _, fish := range d.initialFishes {
		ageCount[fish.timer]++
	}

	for i := 0; i < 256; i++ {
		ageCount = d.increaseCounts(ageCount)
	}

	fishSum := 0
	for _, count := range ageCount {
		fishSum += count
	}

	return fmt.Sprintf("After 256 days there are %d fishies", fishSum)
}

func (DaySix) increaseCounts(ageCount map[int]int) map[int]int {
	newCount := make(map[int]int)
	for age, count := range ageCount {
		if age == 0 {
			newCount[8] += count
			newCount[6] += count
		} else {
			newCount[age-1] += count
		}
	}

	return newCount
}

func (d *DaySix) init() {
	fishes := d.input.Lines[0]
	var initialFishes []*fish
	if len(fishes) <= 0 {
		panic("no fishies")
	}
	for _, s := range strings.Split(fishes, ",") {
		initialFishes = append(initialFishes, &fish{parseInt(s)})
	}

	d.initialFishes = initialFishes
}
