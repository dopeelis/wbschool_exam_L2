package main

import (
	"testing"
)

// Equal функция сравнения двух слайсов
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

// тестирование функции Cut
func TestCut(t *testing.T) {
	testTable := []struct {
		f        int
		d        string
		s        bool
		data     []string
		expected []string
	}{
		{
			f: 2,
			d: "\t",
			s: false,
			data: []string{
				"Winter\tcold\t0",
				"Summer\thot\t10",
				"Autumn\tmedium\t6",
				"Spring\tmed\t7",
			},
			expected: []string{
				"cold\n",
				"hot\n",
				"medium\n",
				"med\n",
			},
		},
		{
			f: 3,
			d: "\t",
			s: false,
			data: []string{
				"245:789\t4567\tM:4540\tAdmin\t01:10:1980",
				"535:763\t4987\tM:3476\tSales\t11:04:1978",
			},
			expected: []string{
				"M:4540\n",
				"M:3476\n",
			},
		},
		{
			f: 2,
			d: ":",
			s: false,
			data: []string{
				"Winter:cold:0",
				"Summer:hot:10",
				"Autumn:medium:6",
				"Spring:med:7",
			},
			expected: []string{
				"cold\n",
				"hot\n",
				"medium\n",
				"med\n",
			},
		},
		{
			f: 3,
			d: " ",
			s: false,
			data: []string{
				"Winter cold 0",
				"Summer:hot:10",
				"Autumn:medium:6",
				"Spring med 7",
			},
			expected: []string{
				"0\n",
				"",
				"",
				"7\n",
			},
		},
		{
			f: 2,
			d: "\t",
			s: true,
			data: []string{
				"Winter\tcold\t0",
				"Summer\thot\t10",
				"Autumn\tmedium\t6",
				"Spring\tmed\t7",
			},
			expected: []string{
				"cold\n",
				"hot\n",
				"medium\n",
				"med\n",
			},
		},
		{
			f: 3,
			d: "\t",
			s: true,
			data: []string{
				"Winter\tcold\t0",
				"Summer:hot:10",
				"Autumn:medium:6",
				"Spring\tmed\t7",
			},
			expected: []string{
				"0\n",
				"",
				"",
				"7\n",
			},
		},
	}

	for _, testCase := range testTable {
		var result []string
		for _, str := range testCase.data {
			result = append(result, Cut(str, testCase.f, testCase.d, testCase.s))
		}
		if !Equal(result, testCase.expected) {
			t.Errorf("Incorrect result. Expect: %v, Got %v\n", testCase.expected, result)
		}
	}
}
