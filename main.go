package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/osagemo/advent-of-go-21/puzzles"
)

var dailyPuzzles []puzzles.Day

func main() {

	printPuzzles()

	printMemoryUsage()
	fmt.Println("Completed, manual GC")
	runtime.GC()
	printMemoryUsage()
}

func printPuzzles() {
	for _, day := range dailyPuzzles {
		printPuzzle(day)
	}
}

func printPuzzle(d puzzles.Day) {
	startAlloc := getTotalAlloc()
	fmt.Println(d.GetPuzzleName())
	fmt.Printf("Part One: %s\n", d.PartOne())
	fmt.Printf("Part Two: %s\n", d.PartTwo())
	start := d.GetStart()
	elapsed := time.Since(start)
	endAlloc := getTotalAlloc()
	allocDif := endAlloc - startAlloc
	fmt.Printf("\nAdditional alloc (after init): %dkB\n", allocDif)
	fmt.Printf("Time elapsed: %s\n", elapsed)
	// printMemoryUsage()
	fmt.Println()
}

func init() {
	// dailyPuzzles = append(dailyPuzzles, puzzles.NewDay(1))
	// dailyPuzzles = append(dailyPuzzles, puzzles.NewDay(2))
	// dailyPuzzles = append(dailyPuzzles, puzzles.NewDay(3))
	// dailyPuzzles = append(dailyPuzzles, puzzles.NewDay(4))
	// dailyPuzzles = append(dailyPuzzles, puzzles.NewDay(5))
	// dailyPuzzles = append(dailyPuzzles, puzzles.NewDay(6))
	// dailyPuzzles = append(dailyPuzzles, puzzles.NewDay(7))
	// dailyPuzzles = append(dailyPuzzles, puzzles.NewDay(8))
	dailyPuzzles = append(dailyPuzzles, puzzles.NewDay(9))
}

func getTotalAlloc() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	kbAlloc := bToKb(m.TotalAlloc)
	return kbAlloc
}

func printMemoryUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	kbAlloc := bToKb(m.Alloc)
	fmt.Printf("Heap Alloc: %dkB\n", kbAlloc)
}

func bToKb(b uint64) uint64 {
	kb := b / 1024
	return kb
}
