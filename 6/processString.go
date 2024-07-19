package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type StringProcessor struct {
	c    *FlagsConfig
	text string
}

func (pr *StringProcessor) Process() ([]string, error) {
	res := make([]string, 0)
	//разделяем на строки по \n
	ss := strings.Split(pr.text, "\n")
	//для каждой строки ss смотрим
	for _, s := range ss {
		words := strings.Split(s, pr.c.delimiter)
		if len(words) == 1 && pr.c.sep {
			continue
		}
		//итерируемся по всем полям и достаем элементы
		res = append(res, "")
		for i := 0; i < len(pr.c.fields); i++ {
			if len(words) > pr.c.fields[i]-1 && pr.c.fields[i]-1 >= 0 {
				res[len(res)-1] += fmt.Sprint(words[pr.c.fields[i]-1], " ")
			} else {
				return nil, errors.New("Incorrect number in fields")
			}
		}
	}
	return res, nil
}

func (pr *StringProcessor) ReadString() {
	s := make([]string, 0)
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		txt := sc.Text()
		s = append(s, txt)
	}
	pr.text = strings.Join(s, "\n")
}
func (pr StringProcessor) WriteResult(ss []string) {
	fmt.Println()
	for _, str := range ss {
		fmt.Println(str)
	}
}
