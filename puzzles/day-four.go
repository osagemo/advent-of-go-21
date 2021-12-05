package puzzles

import (
	"fmt"
	"strconv"
	"strings"
)

type DayFour struct {
	day
	drawingNumbers         []int
	boards                 [][5][5]int
	numberToBoardPosition  map[int][]boardPosition
	boardIndexToMarked     map[int][5][2]int
	winningScore           int
	lastWinningBoardIndex  int
	lastWinningNumberIndex int
	boardsWon              map[int]bool
}

type boardPosition struct {
	boardIndex    int
	boardPosition coordinate
}

type coordinate struct {
	positionX int
	positionY int
}

func (d *DayFour) GetPuzzleName() string {
	return "Day 4: Giant Squid"
}

func (d *DayFour) PartOne() string {
	d.boardsWon = make(map[int]bool)
	d.boardIndexToMarked = make(map[int][5][2]int)

	for i, number := range d.drawingNumbers {
		lastFound := d.drawNumber(number, i)
		if lastFound {
			break
		}
	}
	return fmt.Sprintf("Solution: %d", d.winningScore)
}

func (d *DayFour) PartTwo() string {
	sumOfLosingBoard := d.calculateSum(d.lastWinningBoardIndex, d.lastWinningNumberIndex)
	scoreOfLosingBoard := d.drawingNumbers[d.lastWinningNumberIndex] * sumOfLosingBoard
	return fmt.Sprintf("Solution: %d", scoreOfLosingBoard)
}

func (d *DayFour) drawNumber(number int, numberIndex int) bool {
	matches, found := d.numberToBoardPosition[number]
	if !found {
		return false
	}
	for _, bp := range matches {
		d.increase(bp, number)
		if d.allMatch(bp) {
			d.boardsWon[bp.boardIndex] = true
			sum := d.calculateSum(bp.boardIndex, numberIndex)
			if d.winningScore == 0 {
				d.winningScore = sum * number
			}
			d.lastWinningBoardIndex = bp.boardIndex
			d.lastWinningNumberIndex = numberIndex
			// exit condition for to determine which board was last
			if len(d.boardsWon) == len(d.boards) {
				return true
			}
		}
	}
	return false
}

func (d *DayFour) calculateSum(boardIndex, numberIndex int) int {
	// sum of all unmarked numbers
	unDrawnNumbers := d.drawingNumbers[numberIndex+1 : len(d.drawingNumbers)]
	board := d.boards[boardIndex]
	unmarkedSum := 0
	for _, row := range board {
		for _, val := range row {
			exists := contains(unDrawnNumbers, val)
			if exists {
				unmarkedSum += val
			}
		}
	}
	return unmarkedSum
}

// temporary passing number for debug
func (d *DayFour) increase(bp boardPosition, number int) {
	// given 1, 4
	board := bp.boardIndex
	x := bp.boardPosition.positionX
	y := bp.boardPosition.positionY
	score := d.boardIndexToMarked[board]
	score[x][0] = score[x][0] + 1
	score[y][1] = score[y][1] + 1
	d.boardIndexToMarked[board] = score
}

func (d *DayFour) allMatch(b boardPosition) bool {
	board := b.boardIndex
	x := b.boardPosition.positionX
	y := b.boardPosition.positionY

	return d.boardIndexToMarked[board][x][0] == 5 || d.boardIndexToMarked[board][y][1] == 5
}

// init
func (d *DayFour) init() {
	d.parseDrawingNumbers()
	d.parseBoards()
}

func (d *DayFour) parseDrawingNumbers() {
	numbers := strings.Split(d.input.Lines[0], ",")
	for _, char := range numbers {
		number, _ := strconv.Atoi(string(char))
		d.drawingNumbers = append(d.drawingNumbers, number)
	}
}

func (d *DayFour) parseBoards() {
	var board = [5][5]int{}
	d.numberToBoardPosition = make(map[int][]boardPosition)

	var lineIndex = 0
	for _, line := range d.input.Lines[1:] {
		for i, char := range strings.Fields(line) {
			s := string(char)
			n, _ := strconv.Atoi(s)
			board[lineIndex][i] = n
			boardPosition := boardPosition{len(d.boards), coordinate{i, lineIndex}}
			d.numberToBoardPosition[n] = append(d.numberToBoardPosition[n], boardPosition)
		}
		lineIndex++
		if lineIndex > 0 && lineIndex%5 == 0 {
			d.boards = append(d.boards, board)
			lineIndex = 0
		}
	}
	lineIndex = 1
}
