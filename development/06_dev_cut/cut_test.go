package main

import "testing"

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

func TestFields(t *testing.T) {
	testTable := []struct {
		field    int
		data     []string
		expected []string
	}{
		{
			field: 2,
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
			field: 3,
			data: []string{
				"245:789\t4567\tM:4540\tAdmin\t01:10:1980",
				"535:763\t4987\tM:3476\tSales\t11:04:1978",
			},
			expected: []string{
				"M:4540\n",
				"M:3476\n",
			},
		},
	}

	for _, testCase := range testTable {
		result := fields(testCase.field, testCase.data)
		if !Equal(result, testCase.expected) {
			t.Errorf("Incorrect result. Expect: %v, Got %v\n", testCase.expected, result)
		}
	}
}

func TestDelimiter(t *testing.T) {
	testTable := []struct {
		field     int
		delimiter string
		data      []string
		expected  []string
	}{
		{
			field:     2,
			delimiter: ":",
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
			field:     3,
			delimiter: " ",
			data: []string{
				"Winter cold 0",
				"Summer:hot:10",
				"Autumn:medium:6",
				"Spring med 7",
			},
			expected: []string{
				"0\n",
				"7\n",
			},
		},
	}

	for _, testCase := range testTable {
		result := delimiter(testCase.field, testCase.delimiter, testCase.data)
		if !Equal(result, testCase.expected) {
			t.Errorf("Incorrect result. Expect: %v, Got %v\n", testCase.expected, result)
		}
	}
}

func TestSeparated(t *testing.T) {
	testTable := []struct {
		field    int
		data     []string
		expected []string
	}{
		{
			field: 2,
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
			field: 3,
			data: []string{
				"Winter\tcold\t0",
				"Summer:hot:10",
				"Autumn:medium:6",
				"Spring\tmed\t7",
			},
			expected: []string{
				"0\n",
				"7\n",
			},
		},
	}

	for _, testCase := range testTable {
		result := separated(testCase.field, testCase.data)
		if !Equal(result, testCase.expected) {
			t.Errorf("Incorrect result. Expect: %v, Got %v\n", testCase.expected, result)
		}
	}
}
