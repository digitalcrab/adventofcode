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

func ReadFileIntoBytesMatrix(file io.Reader) ([][]byte, error) {
	matrix := make([][]byte, 0)
	return matrix, ScanFileLineByLine(file, func(line string) {
		matrix = append(matrix, []byte(line))
	})
}
