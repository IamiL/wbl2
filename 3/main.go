package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	//Установка флагов утилиты
	c := NewFlagsConfig()
	flagSet := SetFlags(c)
	flagSet.Usage()
	flag.Parse()

	//Открытие и чтение файла
	ss, err := OpenAndReadFile(flag.Arg(0), c)
	if err != nil {
		fmt.Println("Error with OpenFile:", err)
		os.Exit(1)
	}
	//Создание сортировщика строк
	sStr := &SorterString{
		ss: ss,
		c:  c,
	}
	//Сортировка
	sorter, err := Sort(sStr, c)
	if err != nil {
		fmt.Println("Error with Sort elements:", err)
		os.Exit(1)
	}
	//Вывод результатов
	err = sorter.PrintResult(flag.Arg(0))
	if err != nil {
		fmt.Println("Error with Write string in file:", err)
		os.Exit(1)
	}
}
