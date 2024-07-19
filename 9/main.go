package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var fileNameFlag = flag.String("O", "index.html", "Файл в который будут сохранены полученные данные")

func MainWget() {
	flag.Parse()
	URL := flag.Arg(0)
	fileName := *fileNameFlag
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	//Http request
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)

	defer file.Close()

	fmt.Printf("Downloaded a file %s with size %d", fileName, size)

}
