package main

import "flag"

type FlagsConfig struct {
	numberColumn int
	sortNumber   *bool
	reverse      *bool
	unique       *bool
}

func NewFlagsConfig() *FlagsConfig {
	return &FlagsConfig{}
}

func SetFlags(c *FlagsConfig) *flag.FlagSet {
	return &flag.FlagSet{Usage: func() {
		flag.IntVar(&c.numberColumn, "k", 0, "Номер колонки для сортировки")
		c.sortNumber = flag.Bool("n", false, "Сортировать по численному значению")
		c.reverse = flag.Bool("r", false, "Сортировать в обратном порядке")
		c.unique = flag.Bool("u", false, "Убрать повторяющиеся строки")
	}}
}
