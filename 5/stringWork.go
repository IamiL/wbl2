package main

import (
	"fmt"
	"regexp"
	"strings"
)

type SearchResources struct {
	ss      []string
	c       *FlagConfig
	re      *regexp.Regexp
	pattern string
}

func NewSearchResources(ss []string, c *FlagConfig, re *regexp.Regexp, pattern string) *SearchResources {
	return &SearchResources{ss: ss, c: c, re: re, pattern: pattern}
}

func (sr SearchResources) SearchString() {
	for i, c := 0, 0; i < len(sr.ss) && c < sr.c.count; i++ {
		if sr.ValidateString(i) {
			if *sr.c.num {
				fmt.Println("Строка", i+1)
			}
			sr.BeforeString(i)
			fmt.Println(sr.ss[i])
			sr.AfterString(i)
			fmt.Println()
			sr.ContextString(i)
		}
	}
}

func (sr SearchResources) ValidateString(i int) bool {
	//если флаг -f == true, то просто сравниваем строки,не используя паттерн
	var ans bool
	if *sr.c.fixed {
		//обычное сравнение строк
		if *sr.c.ignore {
			//без учета регистра
			ans = strings.ToLower(sr.ss[i]) == strings.ToLower(sr.pattern)
		} else {
			//с учетом регистра
			ans = sr.ss[i] == sr.pattern
		}
	} else {
		ans = sr.re.MatchString(sr.ss[i])
	}
	return ans != *sr.c.invert
}

func (sr SearchResources) BeforeString(i int) {
	sr.BeforeStringWithParam(i, sr.c.before)
}
func (sr SearchResources) BeforeStringWithParam(i, param int) {
	if param == 0 {
		return
	}
	start := 0
	if i-param > 0 {
		start = i - param
	}
	for j := start; j < i; j++ {
		fmt.Println(sr.ss[j])
	}
}

func (sr SearchResources) AfterString(i int) {
	sr.AfterStringWithParam(i, sr.c.after)
}

func (sr SearchResources) AfterStringWithParam(i, param int) {
	if param == 0 {
		return
	}
	fin := len(sr.ss) - 1
	if i+param < len(sr.ss)-1 {
		fin = i + param
	}
	for j := i + 1; j <= fin; j++ {
		fmt.Println(sr.ss[j])
	}
}
func (sr SearchResources) ContextString(i int) {
	if sr.c.context != 0 {
		sr.BeforeStringWithParam(i, sr.c.context)
		fmt.Println(sr.ss[i])
		sr.AfterStringWithParam(i, sr.c.context)
		fmt.Println()
	}
}
