package main

import "flag"

type FlagConfig struct {
	after   int
	before  int
	context int
	count   int
	ignore  *bool
	invert  *bool
	fixed   *bool
	num     *bool
}

func NewFlagConfig() *FlagConfig {
	return &FlagConfig{count: 1 << 62}
}

func SetFlags(c *FlagConfig) *flag.FlagSet {

	return &flag.FlagSet{Usage: func() {
		flag.IntVar(&c.after, "A", 0, "Число строк после найденной")
		flag.IntVar(&c.before, "B", 0, "Число строк до найденной")
		flag.IntVar(&c.context, "C", 0, "Число строк до и после найденной")
		flag.IntVar(&c.count, "c", 1<<62, "Максимальное число строк")
		c.ignore = flag.Bool("i", false, "Игнорировать регистр")
		c.invert = flag.Bool("v", false, "Вместо совпадения, совпадения исключать")
		c.fixed = flag.Bool("F", false, "Точное совпадение со строкой а не паттерн")
		c.num = flag.Bool("n", false, "Печатать номер строки ")
	}}
}
