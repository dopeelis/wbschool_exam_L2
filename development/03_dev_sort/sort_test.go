package main

import (
	"fmt"
	"testing"
)

// функция для сравнения двух слайсов
func Equal(a, b []string) bool {
	// сразу проверяем размер
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

// тестируем обычную сортировку
func TestStrSort(t *testing.T) {
	testTable := []struct {
		strs     []string
		k        int
		n, r, u  bool
		expected []string
	}{
		{
			strs:     []string{"e aba", "g aca", "c bab", "d bba", "f aaa", "b aaa"},
			k:        0,
			n:        false,
			r:        false,
			u:        false,
			expected: []string{"b aaa", "c bab", "d bba", "e aba", "f aaa", "g aca"},
		},
		{
			strs:     []string{"aa", "ba", "ab", "bb", "ca"},
			k:        0,
			n:        false,
			r:        false,
			u:        false,
			expected: []string{"aa", "ab", "ba", "bb", "ca"},
		},
		{
			strs:     []string{"aa", "ba", "ab", "bb", "ca"},
			k:        0,
			n:        false,
			r:        true,
			u:        false,
			expected: []string{"ca", "bb", "ba", "ab", "aa"},
		},
		{
			strs:     []string{"e aba", "g aca", "c bab", "d bba", "b aaa"},
			k:        1,
			n:        false,
			r:        true,
			u:        false,
			expected: []string{"d bba", "c bab", "g aca", "e aba", "b aaa"},
		},
		{
			strs:     []string{"aa", "ba", "ab", "bb", "ca", "aa", "ca"},
			k:        0,
			n:        false,
			r:        false,
			u:        true,
			expected: []string{"aa", "ab", "ba", "bb", "ca"},
		},
		{
			strs:     []string{"e aba", "g aca", "c bab", "d bba", "f aaa", "b aaa", "g aca", "d bba", "f aaa", "b aaa"},
			k:        0,
			n:        false,
			r:        true,
			u:        true,
			expected: []string{"g aca", "f aaa", "e aba", "d bba", "c bab", "b aaa"},
		},
		{
			strs:     []string{"3aa", "12ba", "8ab", "46bb", "11ca"},
			k:        0,
			n:        true,
			r:        false,
			u:        false,
			expected: []string{"3aa", "8ab", "11ca", "12ba", "46bb"},
		},
		{
			strs:     []string{"12a", "45b", "33cac", "t"},
			k:        0,
			n:        true,
			r:        false,
			u:        false,
			expected: []string{"t", "12a", "33cac", "45b"},
		},
		{
			strs:     []string{"b a", "a b", "c d", "d c"},
			k:        0,
			n:        false,
			r:        false,
			u:        true,
			expected: []string{"a b", "b a", "c d", "d c"},
		},
		{
			strs:     []string{"b a", "a b", "c d", "d c"},
			k:        1,
			n:        false,
			r:        false,
			u:        false,
			expected: []string{"b a", "a b", "d c", "c d"},
		},
	}

	for _, testCase := range testTable {
		fmt.Println("Data: ", testCase)
		result := strSort(testCase.strs, testCase.k, testCase.n, testCase.r, testCase.u)
		fmt.Println("Result: ", result)
		fmt.Println("Expected:", testCase.expected)
		if !Equal(result, testCase.expected) {
			t.Errorf("Incorrect result. Expect: %v, Got %v\n", testCase.expected, result)
		}
	}
}
