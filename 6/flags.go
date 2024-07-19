package main

import (
	"flag"
	"strconv"
	"strings"
)

type Flags struct {
	fields    *string
	delimiter *string
	sep       *bool
}

func NewFlag() *Flags {
	return &Flags{}
}
func SetFlags(c *Flags) *flag.FlagSet {
	return &flag.FlagSet{Usage: func() {
		c.sep = flag.Bool("s", false, "Печатать только строки с разделителем")
		c.fields = flag.String("f", "", "Строка полей, которые нужно вывести")
		c.delimiter = flag.String("d", " ", "Использовать другой разделитель ")
	}}
}

type FlagsConfig struct {
	fields    []int
	delimiter string
	sep       bool
}

func NewFlagConfig(f *Flags) (*FlagsConfig, error) {
	ss := strings.Split(*f.fields, ",")
	arr := make([]int, len(ss))
	for i := 0; i < len(ss); i++ {
		var err error
		arr[i], err = strconv.Atoi(ss[i])
		if err != nil {
			return nil, err
		}
	}
	return &FlagsConfig{
		fields:    arr,
		delimiter: *f.delimiter,
		sep:       *f.sep,
	}, nil
}
