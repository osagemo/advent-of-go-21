package input

import (
	"os"
	"path"
	"regexp"
	"runtime"
	"testing"
)

func TestGetPathByDayNumber(t *testing.T) {
	day := 1
	want := regexp.MustCompile(`.+day-one.txt`)

	path, err := getPath(1)
	if !want.MatchString(path) || err != nil {
		t.Fatalf(`getPath(%v) = %q, %v`, day, path, err)
	}
}

func TestGetDayOneInput(t *testing.T) {
	day := 1
	wantLength := 2000
	wantLast := "6632"
	inp := GetInput(day)
	lastLine := inp.Lines[len(inp.Lines)-1]

	if inp.Length != 2000 {
		t.Fatalf("unxepected input length: %d, expected %d", inp.Length, wantLength)
	}

	if lastLine != wantLast {
		t.Fatalf("unexpected last line: %v, expected: %v", lastLine, wantLast)
	}
}

func init() {
	_, filename, _, _ := runtime.Caller(0)
	// The ".." may change depending on you folder structure
	dir := path.Join(path.Dir(filename), `..\..`)
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}
