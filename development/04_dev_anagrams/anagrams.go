package main

import (
	"sort"
	"strings"
)

// функция для получения анаграм
func getAnagrams(arr []string) map[string][]string {
	// создаем мап, куда будем складывать результат
	res := make(map[string][]string)

	// приводим все к нижнему регистру
	// и добавляем все слова как ключи в мапе
	for _, i := range arr {
		i = strings.ToLower(i)
		res[i] = []string{}
	}

	// проходим по ключам и исключаем те, что являются анаграммами
	for k := range res {
		for _, i := range arr {
			i = strings.ToLower(i)
			sortStr := SortString(i)
			sortKey := SortString(k)
			if i == k {
				break
			}
			if sortStr == sortKey {
				delete(res, k)
			}
		}
	}

	// проходим по словам и добавляем в мап, если являются анаграммой к ключу
	for _, i := range arr {
		i = strings.ToLower(i)
		for k := range res {
			sortStr := SortString(i)
			sortKey := SortString(k)
			if i == k {
				break
			}
			if sortStr == sortKey {
				res[k] = append(res[k], i)
			}
		}
	}

	// убираем ключи без значений
	// т.е. без анаграмм
	for k, v := range res {
		if len(v) == 0 {
			delete(res, k)
		}
	}

	return res
}

// реализация сортировки строки по элементам
// для сравнения с другими словами
type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
	// в итоге получаем отсортированную по рунам строку
	// если две отсортированные строки совпадают,
	// значит они являются анаграммами
}
