package puzzles

import (
	"testing"
)

// for example 2 & 3
func init11Test() DayEleven {
	exampleInput := []string{
		"5483143223",
		"2745854711",
		"5264556173",
		"6141336146",
		"6357385478",
		"4167524645",
		"2176841721",
		"6882881134",
		"4846848554",
		"5283751526"}

	d := DayEleven{}
	d.day = day{inputLines: exampleInput}
	d.init()

	return d
}
func TestExample11_1(t *testing.T) {
	exampleInput := []string{"11111", "19991", "19191", "19991", "11111"}
	d := DayEleven{}
	d.day = day{inputLines: exampleInput}
	d.init()
	// fmt.Println(d.grid)
	wantNumFlashes := 9
	d.grid.incrementEnergyLevels(1, false)
	if d.grid.numFlashes != wantNumFlashes {
		t.Fatalf("expected %d flashes but got %d", wantNumFlashes, d.grid.numFlashes)
	}
	// fmt.Println(d.grid)
}

func TestExample11_2(t *testing.T) {
	d := init11Test()
	// fmt.Println(d.grid)
	d.grid.incrementEnergyLevels(100, false)
	// fmt.Println(d.grid)
	wantNumFlashes := 1656
	if d.grid.numFlashes != wantNumFlashes {
		t.Fatalf("exptected %d flashes but got %d", wantNumFlashes, d.grid.numFlashes)
	}
}

func TestExample11_3(t *testing.T) {
	d := init11Test()
	d.grid.incrementEnergyLevels(1000, true)

	wantFirstFullFlashAt := 195
	if d.grid.firstFullFlash != wantFirstFullFlashAt {
		t.Fatalf("expected first full flash at %d, got at %d", wantFirstFullFlashAt, d.grid.firstFullFlash)
	}
}
