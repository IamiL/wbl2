package main

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

func OpenAndReadFile(filename string, c *FlagsConfig) ([]FileString, error) {
	var in io.Reader
	if filename != "" {
		f, err := os.Open(filename)
		if err != nil {
			return nil, err
		}
		defer f.Close()

		in = f
	}
	return ReadFile(bufio.NewScanner(in), c)
}

func ReadFile(s *bufio.Scanner, c *FlagsConfig) ([]FileString, error) {
	result := make([]FileString, 0)
	for s.Scan() {
		result = append(result, FileString{FullString: s.Text()})
	}
	//Если флаг -k не равен нулю, то считываем не всю строку, а только i конкретный столбец
	if c.numberColumn == 0 {
		for i := range result {
			result[i].MainPart = result[i].FullString
		}
	} else {
		for i, el := range result {
			ss := strings.Split(el.FullString, " ")
			if len(ss) >= c.numberColumn { // в строке есть требуемая колонка
				result[i].MainPart = ss[c.numberColumn-1]
			} else {
				return nil, errors.New("Element is not in line")
			}
		}
	}

	return result, nil
}
