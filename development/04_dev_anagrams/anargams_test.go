package main

import (
	"testing"
)

// функция сравнения двух слайсов
func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// тестируем функцию получения анаграмм
func TestGetAnagrams(t *testing.T) {
	testTable := []struct {
		data        []string
		anagramsMap map[string][]string
	}{
		{
			data: []string{"Пятак", "пятка", "Тяпка", "листок", "слиток", "столиК", "ЛЕС"},
			anagramsMap: map[string][]string{
				"пятак":  {"пятка", "тяпка"},
				"листок": {"слиток", "столик"},
			},
		},
		{
			data: []string{"горбик", "корнет", "грибок", "сани", "пискун", "супник", "гробик", "кретон", "лосось", "ректон"},
			anagramsMap: map[string][]string{
				"горбик": {"грибок", "гробик"},
				"корнет": {"кретон", "ректон"},
				"пискун": {"супник"},
			},
		},
	}

	for _, testCase := range testTable {
		result := getAnagrams(testCase.data)
		// сразу исключаем несовпадение, чтобы не проходить по циклу, если разное количество ключей
		if len(result) != len(testCase.anagramsMap) {
			t.Errorf("Incorrect result. Expect: %v, Got %v\n", testCase.anagramsMap, result)
		}
		for k, v := range result {
			if !Equal(v, testCase.anagramsMap[k]) {
				t.Errorf("Incorrect result. Expect: %v, Got %v\n", testCase.anagramsMap[k], v)
			}
		}
	}
}
