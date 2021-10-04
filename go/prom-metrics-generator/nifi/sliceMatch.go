package nifi

import (
        "io"
        "bufio"
        "strings"
)

func fileSliceLineMatch(
	rdr io.Reader,
	match string,
) ([]int,[]string) {
	lineNumber := []int{}
	lines := []string{}
	scanner := bufio.NewScanner(rdr)
	i := 1
	for scanner.Scan() {
			lines = append(lines, scanner.Text())
			if strings.Contains(scanner.Text(), match) {
					lineNumber = append(lineNumber, i)
			}
			i++
	}
	return lineNumber,lines
}
