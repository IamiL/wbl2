package main

import (
	"flag"
	"fmt"
	"os"
)

func MainUtil() {
	//Задание флагов для утилиты
	c := NewFlagConfig()
	flagSet := SetFlags(c)
	flagSet.Usage()
	flag.Parse()

	//Задание паттерна
	pattern := flag.Arg(1)
	re := SetPattern(pattern, *c.ignore)

	//Открытие и чтение файла
	ss, err := OpenAndReadFile(flag.Arg(0))
	if err != nil {
		fmt.Println("Error with OpenFile:", err)
		os.Exit(1)
	}

	//Анализ строк файла
	sr := NewSearchResources(ss, c, re, pattern)
	sr.SearchString()
}
