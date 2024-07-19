package main

import (
	"bufio"
	"io"
	"os"
)

func OpenAndReadFile(filename string) ([]string, error) {
	var in io.Reader
	if filename != "" {
		f, err := os.Open(filename)
		if err != nil {
			return nil, err
		}
		defer f.Close()

		in = f
	}
	return ReadFile(bufio.NewScanner(in)), nil
}

func ReadFile(s *bufio.Scanner) []string {
	result := make([]string, 0)
	for s.Scan() {
		result = append(result, s.Text())
	}
	return result
}
