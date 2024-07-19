package main

import (
	"fmt"
	"sort"
)

func MainSearchAn() {
	arr := []string{"пятак", "пятак", "пятка", "столик", "листок", "слиток", "тяпка", "столик", "патк", "столик"}
	m := searchAn(arr)
	fmt.Println(m)
}

type stringCounter struct {
	r []rune
	i int
}

func searchAn(arr []string) map[string][]string {
	sortArr := make([]stringCounter, len(arr))
	//создание массива
	for i, s := range arr {
		sortArr[i] = stringCounter{
			r: []rune(s),
			i: i,
		}
	}
	//сортировка каждого из элементов в массиве
	for k := range sortArr {
		sort.Slice(sortArr[k].r, func(i, j int) bool {
			return sortArr[k].r[i] < sortArr[k].r[j]
		})
	}
	//сортировка массива в целом
	sort.Slice(sortArr, func(i, j int) bool {
		ln := Min(len(sortArr[i].r), len(sortArr[j].r))
		for k := 0; k < ln; k++ {
			if sortArr[i].r[k] < sortArr[j].r[k] {
				return true
			}
			if sortArr[i].r[k] > sortArr[j].r[k] {
				return false
			}
		}
		if ln == len(sortArr[i].r) {
			return true
		} else {
			return false
		}
	})
	//на данном этапе отсортирован массив строк в котором каждая и строк отсортирована
	//теперь идем по массиву и сравниваем - если соседние строки равны добавляем их в один слайс
	groupStr := make([][]int, 1)
	groupStr[0] = append(groupStr[0], sortArr[0].i)
	for i := 1; i < len(sortArr); i++ {
		if !isEqualRuneArr(sortArr[i].r, sortArr[i-1].r) {
			groupStr = append(groupStr, []int{})
		}
		//добавление элемента
		groupStr[len(groupStr)-1] = append(groupStr[len(groupStr)-1], sortArr[i].i)
	}
	/// используя groupStr получим финальный результат
	res := make(map[string][]string)
	for i := range groupStr {
		if len(groupStr[i]) == 1 {
			continue
		}
		//ищем элемент с наименьшим номером и устанавливаем как ключ
		numMin := MinEl(groupStr[i])
		res[arr[numMin]] = make([]string, len(groupStr[i]))
		for k, num := range groupStr[i] {
			res[arr[numMin]][k] = arr[num]
		}
	}
	//слова лежат в мапке с правильным ключем, но не убраны повторения слов, слова не отсортированы
	//сначала отсортируем слова
	for s := range res {
		sort.Strings(res[s])
	}
	//идя по массиву отсортированных слов будем убирать повторения
	for s := range res {
		newS := make([]string, 1)
		newS[0] = res[s][0]
		//добавление уникальных элементов
		for _, str := range res[s] {
			if !(str == newS[len(newS)-1]) {
				newS = append(newS, str)
			}
		}
		res[s] = newS
	}
	return res
}
