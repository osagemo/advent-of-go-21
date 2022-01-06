package input

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

var dayMap map[int]string

// Reads whole file to memory to simplify processing since expected input of AoC is never large
// Should probably return error instead of panicing, but i wanted to test it
func GetInput(day int) []string {
	path, err := getPath(day)
	check(err)

	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			lines = append(lines, line)
		}
	}

	return lines
}

func getPath(day int) (string, error) {
	path, err := filepath.Abs(fmt.Sprintf("./puzzles/input/day-%d.txt", day))
	check(err)

	return path, nil
}

// Used with initial implementation where Input had a scanner
// func lineCounter(r io.Reader) (int, error) {
// 	buf := make([]byte, 32*1024)
// 	count := 0
// 	lineSep := []byte{'\n'}

// 	for {
// 		c, err := r.Read(buf)
// 		count += bytes.Count(buf[:c], lineSep)

// 		switch {
// 		case err == io.EOF:
// 			return count, nil

// 		case err != nil:
// 			return count, err
// 		}
// 	}
// }

func check(e error) {
	if e != nil {
		panic(e)
	}
}
