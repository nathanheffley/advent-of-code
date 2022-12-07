package input

import (
	"bufio"
	"os"
	"strings"

	"github.com/nathanheffley/advent-of-code/helpers"
)

func ReadInputFileToLines(filename string) []string {
	dat, err := os.ReadFile(filename)
	helpers.Check(err)

	var lines []string
	sc := bufio.NewScanner(strings.NewReader(string(dat)))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}
