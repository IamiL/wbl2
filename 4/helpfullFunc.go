package main

func Min(l, r int) int {
	if l < r {
		return l
	} else {
		return r
	}
}
func isEqualRuneArr(arr1, arr2 []rune) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	ln := len(arr1)
	for i := 0; i < ln; i++ {
		if !(arr1[i] == arr2[i]) {
			return false
		}
	}
	return true
}

func MinEl(arr []int) int {
	min := arr[0]
	for _, el := range arr {
		if min > el {
			min = el
		}
	}
	return min
}
