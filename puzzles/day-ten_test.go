package puzzles

import (
	"fmt"
	"sort"
	"testing"
)

func TestExample10_1(t *testing.T) {
	wantExpected := ']'
	wantFound := '}'

	d := DayTen{}
	d.init()
	var line = "{([(<{}[<>[]}>{[]{[(<()>"
	isCorrupted, expected, found := d.lineIsCorrupted(line)
	if isCorrupted {
		fmt.Printf("Expected %v but found %v instead\n", string(expected), string(found))
	} else {
		fmt.Printf("Line is not corrupted")
	}
	if expected != wantExpected || found != wantFound {
		t.Fatalf("want expected: %c, found: %c, got expected: %c, found: %c", wantExpected, wantFound, expected, found)
	}
}

func TestExample10_2(t *testing.T) {
	wantExpected := ']'
	wantFound := ')'

	d := DayTen{}
	d.init()
	var line = "[[<[([]))<([[{}[[()]]]"
	isCorrupted, expected, found := d.lineIsCorrupted(line)
	if isCorrupted {
		fmt.Printf("Expected %v but found %v instead\n", string(expected), string(found))
	} else {
		fmt.Printf("Line not corrupted")
	}
	if expected != wantExpected || found != wantFound {
		t.Fatalf("want expected: %c, found: %c, got expected: %c, found: %c", wantExpected, wantFound, expected, found)
	}
}

func TestValidLine(t *testing.T) {
	d := DayTen{}
	d.init()
	var line = "[<>({}){}[([])<>]]"
	isCorrupted, expected, found := d.lineIsCorrupted(line)
	if isCorrupted {
		fmt.Printf("Expected %v but found %v instead\n", string(expected), string(found))
	} else {
		fmt.Printf("Line is not corrupted\n")
	}
	if isCorrupted {
		t.Fatalf("expected non-corrupted line, got corrupted")
	}
}

func TestCompletionStringExample1(t *testing.T) {
	d := DayTen{}
	d.init()
	var line = "[({(<(())[]>[[{[]{<()<>>"
	want := "}}]])})]"
	got, _ := d.getClosingChars(line)

	if want != got {
		t.Fatalf("expected: %s, got: %s", want, got)
	}
}

func Test10_PartOne(t *testing.T) {
	wantScore := 26397
	d := DayTen{}
	d.init()
	lines := []string{
		"[({(<(())[]>[[{[]{<()<>>",
		"[(()[<>])]({[<{<<[]>>(",
		"{([(<{}[<>[]}>{[]{[(<()>",
		"(((({<>}<{<{<>}{[]{[]{}",
		"[[<[([]))<([[{}[[()]]]",
		"[{[{({}]{}}([{[{{{}}([]",
		"{<[[]]>}<{[{[{[]{()[[[]",
		"[<(<(<(<{}))><([]([]()",
		"<{([([[(<>()){}]>(<<{{",
		"<{([{{}}[<[[[<>{}]]]>[]]",
	}
	gotScore := d.getCorruptedScore(lines)
	if wantScore != gotScore {
		t.Fatalf("expected score: %d but got score: %d", wantScore, gotScore)
	}
}

func Test10_PartTwo(t *testing.T) {
	wantScore := 288957
	d := DayTen{}
	d.init()
	lines := []string{
		"[({(<(())[]>[[{[]{<()<>>",
		"[(()[<>])]({[<{<<[]>>(",
		"{([(<{}[<>[]}>{[]{[(<()>",
		"(((({<>}<{<{<>}{[]{[]{}",
		"[[<[([]))<([[{}[[()]]]",
		"[{[{({}]{}}([{[{{{}}([]",
		"{<[[]]>}<{[{[{[]{()[[[]",
		"[<(<(<(<{}))><([]([]()",
		"<{([([[(<>()){}]>(<<{{",
		"<{([{{}}[<[[[<>{}]]]>[]]",
	}
	scores := d.getAutoCompleteScores(lines)

	sort.Ints(scores)
	half := len(scores) / 2
	fmt.Println(half)
	middleIndex := len(scores) / 2
	gotScore := scores[middleIndex]
	if wantScore != gotScore {
		t.Fatalf("expected score: %d but got score: %d", wantScore, gotScore)
	}
}
