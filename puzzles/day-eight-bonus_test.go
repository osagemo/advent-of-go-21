package puzzles

import (
	"testing"
)

func TestPrinter(t *testing.T) {
	p := printer{2, 4, 2}

	// testDigits := []string{"  111111", "     11 ", " 1 11 11", " 1  1111", " 11  11 ", " 11 11 1", " 11111 1"} // "     111", " 1111111", " 11 1111", " 111 111", " 11111  ", "  111  1", " 1 1111 ", " 1111  1", " 111   1"}
	// testDigits := []string{"  111111", "     11 ", " 1 11 11", " 1  1111", " 11  11 ", " 11 11 1", " 11111 1", "     111", " 1111111", " 11 1111", " 111 111", " 11111  ", "  111  1", " 1 1111 ", " 1111  1", " 111   1"}
	testDigits := []string{"  111111", "     11 ", " 1 11 11", " 1  1111", " 11  11 ", " 11 11 1", " 11111 1", "     111", " 1111111", " 11 1111"}
	// testDigits := []string{".....111", ".....11."}
	display := p.getDisplay(testDigits)

	for _, row := range display {
		t.Logf("%s\n", row)
		if len(row) < 5 {
			t.Fatalf("not good for %s", row)
		}
	}
}
