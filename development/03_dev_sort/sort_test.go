package main

import (
	"fmt"
	"testing"
)

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

func TestStrSort(t *testing.T) {
	testTable := []struct {
		strs     []string
		expected []string
	}{
		{
			strs:     []string{"aa", "ba", "ab", "bb", "ca"},
			expected: []string{"aa", "ab", "ba", "bb", "ca"},
		},
		{
			strs:     []string{"e aba", "g aca", "c bab", "d bba", "f aaa", "b aaa"},
			expected: []string{"b aaa", "c bab", "d bba", "e aba", "f aaa", "g aca"},
		},
	}

	for _, testCase := range testTable {
		result := strSort(testCase.strs)
		if !Equal(result, testCase.expected) {
			t.Errorf("Incorrect result. Expect: %v, Got %v\n", testCase.expected, result)
		}
	}
}

func TestReverseSort(t *testing.T) {
	testTable := []struct {
		strs     []string
		expected []string
	}{
		{
			strs:     []string{"aa", "ba", "ab", "bb", "ca"},
			expected: []string{"ca", "bb", "ba", "ab", "aa"},
		},
		{
			strs:     []string{"e aba", "g aca", "c bab", "d bba", "f aaa", "b aaa"},
			expected: []string{"g aca", "f aaa", "e aba", "d bba", "c bab", "b aaa"},
		},
	}

	for _, testCase := range testTable {
		result := reverseSort(testCase.strs)
		if !Equal(result, testCase.expected) {
			t.Errorf("Incorrect result. Expect: %v, Got %v\n", testCase.expected, result)
		}
	}
}

func TestSortWithoutRepeat(t *testing.T) {
	testTable := []struct {
		strs     []string
		expected []string
	}{
		{
			strs:     []string{"aa", "ba", "ab", "bb", "ca", "aa", "ca"},
			expected: []string{"aa", "ab", "ba", "bb", "ca"},
		},
		{
			strs:     []string{"e aba", "g aca", "c bab", "d bba", "f aaa", "b aaa", "g aca", "d bba", "f aaa", "b aaa"},
			expected: []string{"b aaa", "c bab", "d bba", "e aba", "f aaa", "g aca"},
		},
	}

	for _, testCase := range testTable {
		result := sortWithoutRepeat(testCase.strs)
		if !Equal(result, testCase.expected) {
			t.Errorf("Incorrect result. Expect: %v, Got %v\n", testCase.expected, result)
		}
	}
}

func TestIntSort(t *testing.T) {
	testTable := []struct {
		strs     []string
		expected []string
	}{
		{
			strs:     []string{"3aa", "12ba", "8ab", "46bb", "11ca"},
			expected: []string{"3aa", "8ab", "11ca", "12ba", "46bb"},
		},
		{
			strs:     []string{"12a", "45b", "33cac", "t"},
			expected: []string{"t", "12a", "33cac", "45b"},
		},
	}

	for _, testCase := range testTable {
		result := intSort(testCase.strs)
		if !Equal(result, testCase.expected) {
			t.Errorf("Incorrect result. Expect: %v, Got %v\n", testCase.expected, result)
		}
	}
}

func TestColumnSort(t *testing.T) {
	testTable := []struct {
		strs      []string
		columnNum int
		expected  []string
	}{
		{
			strs:      []string{"b a", "a b", "c d", "d c"},
			columnNum: 1,
			expected:  []string{"a b", "b a", "c d", "d c"},
		},
		{
			strs:      []string{"b a", "a b", "c d", "d c"},
			columnNum: 2,
			expected:  []string{"b a", "a b", "d c", "c d"},
		},
	}

	for _, testCase := range testTable {
		result := columnSort(testCase.strs, testCase.columnNum-1)
		if !Equal(result, testCase.expected) {
			for i, s := range result {
				fmt.Println(i, s)
			}
			for i, s := range testCase.expected {
				fmt.Println(i, s)
			}

			t.Errorf("Incorrect result. Expect: %v, Got %v\n", testCase.expected, result)
		}
	}
}
