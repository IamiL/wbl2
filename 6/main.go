package main

import (
	"flag"
	"fmt"
	"os"
)

func MainCut() {
	//Получение флагов
	f := NewFlag()
	flags := SetFlags(f)
	flags.Usage()
	flag.Parse()

	//Создание обработчика
	fConf, err := NewFlagConfig(f)
	if err != nil {
		fmt.Println("Error reading fields -f:", err)
		os.Exit(1)
	}
	pr := StringProcessor{c: fConf}
	//Чтение Str
	pr.ReadString()
	//Обработка и вывод результата
	res, err := pr.Process()
	if err != nil {
		fmt.Println("Such number in text are not:", err)
		os.Exit(1)
	}
	pr.WriteResult(res)
}
