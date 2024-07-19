package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

type FileInt struct {
	FullString string
	MainPart   int
}

type SorterInt struct {
	si []FileInt
	c  *FlagsConfig
}

func (s *SorterInt) Sort(r bool) {
	sort.Slice(s.si, func(i, j int) bool {
		return s.si[i].MainPart < s.si[j].MainPart != r
	})
}

func (s *SorterInt) PrintResult(filename string) error {
	f, err := os.OpenFile(filename, os.O_WRONLY, 0777)
	defer f.Close()
	if err != nil {
		return err
	}
	for _, el := range s.si {
		_, err = f.WriteString(fmt.Sprintf("%s\n", el.FullString))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *SorterInt) Unify(u bool) {
	if u == false {
		return
	}
	newSi := make([]FileInt, 1)
	if len(s.si) == 0 {
		return
	}
	newSi[0] = s.si[0]
	for i, num := range s.si {
		if num != s.si[i] {
			newSi = append(newSi, num)
		}
	}
	s.si = newSi
}

//При наличии флага -n конвертируем

func SortConvertor(sStr *SorterString) (*SorterInt, error) {
	arr := make([]FileInt, len(sStr.ss))
	for i, str := range sStr.ss {
		var err error
		arr[i].FullString = sStr.ss[i].FullString
		arr[i].MainPart, err = strconv.Atoi(str.MainPart)
		if err != nil {
			return nil, err
		}
	}
	return &SorterInt{
		si: arr,
		c:  sStr.c,
	}, nil
}
