package puzzles

import (
	"testing"
)

func TestLine(t *testing.T) {
	d := DayTen{}
	var line = "{([(<{}[<>[]}>{[]{[(<()>"
	d.validateLine(line)
}
