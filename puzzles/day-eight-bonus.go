package puzzles

import "strings"

type printer struct {
	digitSeparatorLength int
	horizontalWidth      int
	verticalHeight       int
}

func (p printer) middle(sb *strings.Builder, digit string, row int) {
	middle := 7
	if row == 3 {
		middle = 1
	} else if row == 6 {
		middle = 4
	}

	sb.WriteByte(' ')
	for i := 0; i < p.horizontalWidth; i++ {
		sb.WriteByte(digit[middle])
	}
	sb.WriteByte(' ')
}

func (p printer) leftRight(sb *strings.Builder, digit string, row int) {
	left := 2
	right := 6

	if row > 3 {
		left = 3
		right = 5
	}

	sb.WriteByte(digit[left])
	for i := 0; i < p.horizontalWidth; i++ {
		sb.WriteByte(' ')
	}
	sb.WriteByte(digit[right])
}

func (p printer) buildLine(sb *strings.Builder, digit string, row int) {
	if row == 0 || row == 3 || row == 6 {
		p.middle(sb, digit, row)
	} else {
		p.leftRight(sb, digit, row)
	}

	// spacing
	for i := 0; i < p.digitSeparatorLength; i++ {
		sb.WriteByte(' ')
	}
}

// (p)GFEDCBA
func (p printer) getDisplay(digits []string) [7]string {
	display := [7]string{}

	for row := 0; row < 7; row++ {
		var sb strings.Builder
		for _, digit := range digits {
			p.buildLine(&sb, digit, row)
		}
		display[row] = sb.String()
		sb.Reset()
	}

	return display
}

// PGFEDCBA:

//0 dot
//1 middle G
//2 top left F
//3 bottom left E
//4 bottom D
//5 bottom right C
//6 top right B
//7 top A
