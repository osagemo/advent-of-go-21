package main

import (
	"fmt"
	"time"

	"github.com/osagemo/advent-of-go-21/puzzles"
)

var dailyPuzzles []puzzles.Day

func main() {
	for _, day := range dailyPuzzles {
		printPuzzle(day)
	}
}

func printPuzzle(d puzzles.Day) {
	start := time.Now()
	fmt.Println(d.GetPuzzleName())
	fmt.Printf("Part One: %s\n", d.PartOne())
	fmt.Printf("Part Two: %s\n", d.PartTwo())
	elapsed := time.Since(start)
	fmt.Printf("Time elapsed: %s\n", elapsed)
	fmt.Println()
}

func init() {
	dailyPuzzles = append(dailyPuzzles, puzzles.NewDay(1))
	dailyPuzzles = append(dailyPuzzles, puzzles.NewDay(2))
	dailyPuzzles = append(dailyPuzzles, puzzles.NewDay(3))
}
