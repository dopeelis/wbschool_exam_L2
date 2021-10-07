package main

import (
	"sort"
	"strings"
)

// функция для получения анаграм
func getAnagrams(arr []string) map[string][]string {
	// создаем мап, куда будем складывать промежуточный результат
	interimRes := make(map[string][]string)

	// создаем мап, куда будем складывать конечный результат
	finalRes := make(map[string][]string)

	// проходим циклом по всему списку
	for _, i := range arr {
		// приводим к нижнему регистру
		lowerI := strings.ToLower(i)
		// сортируем строку
		sortI := SortString(lowerI)
		// если остортированное слово уже является ключом,
		// добавляем само слово в значение к этому ключу
		if _, ok := interimRes[sortI]; ok {
			interimRes[sortI] = append(interimRes[sortI], lowerI)
			//	если нет, то объявляем сортированное слово ключом и добавляем обычное слоово в значение
		} else {
			interimRes[sortI] = []string{lowerI}
		}
	}

	// проходим по промежуточной мапе
	for _, v := range interimRes {
		// если состоит из 1 слова, то пропускаем
		if len(v) <= 1 {
			continue
		}
		// иначе ставим первое слово ключом, остальные слова - значением
		finalRes[v[0]] = v[1:]
	}

	return finalRes
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
