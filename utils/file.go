package utils

import (
	"bufio"
	"io"
)

func ScanFileLineByLine(file io.Reader, cb func(line string)) error {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		cb(scanner.Text())
	}

	return nil
}
