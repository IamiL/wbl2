package main

type Sorter interface {
	Sort(r bool)
	Unify(u bool)
	PrintResult(filename string) error
}

func Sort(ss *SorterString, c *FlagsConfig) (Sorter, error) {
	var ans Sorter
	var err error
	if *c.sortNumber {
		ans, err = SortConvertor(ss)
		if err != nil {
			return nil, err
		}
	} else {
		ans = ss
	}
	ans.Sort(*c.reverse)
	ans.Unify(*c.unique)
	return ans, nil
}
