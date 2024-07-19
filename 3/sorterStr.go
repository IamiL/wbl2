package main

import (
	"os"
	"sort"
)

type FileString struct {
	FullString string
	MainPart   string
}

type SorterString struct {
	ss []FileString
	c  *FlagsConfig
}

func (s *SorterString) Sort(r bool) {
	sort.Slice(s.ss, func(i, j int) bool {
		return s.ss[i].MainPart < s.ss[j].MainPart != r
	})
}

func (s *SorterString) PrintResult(filename string) error {
	f, err := os.OpenFile(filename, os.O_WRONLY, 0777)
	defer f.Close()
	if err != nil {
		return err
	}
	for _, el := range s.ss {
		_, err = f.WriteString(el.FullString + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *SorterString) Unify(u bool) {
	if u == false {
		return
	}
	newSi := make([]FileString, 1)
	if len(s.ss) == 0 {
		return
	}
	newSi[0] = s.ss[0]
	for i, num := range s.ss {
		if num != s.ss[i] {
			newSi = append(newSi, num)
		}
	}
	s.ss = newSi
}
