package main

import (
	"errors"
	"testing"
)

// тестируем функцию распаковки строки
func TestUnpacking(t *testing.T) {
	// создаем наборы для тестирования
	// с разными вариантами входящих и исходящих данных
	testTable := []struct {
		str         string
		expectedVal string
		expectedErr error
	}{
		{
			str:         "a4bc2d5e",
			expectedVal: "aaaabccddddde",
			expectedErr: nil,
		},
		{
			str:         "abcd",
			expectedVal: "abcd",
			expectedErr: nil,
		},
		{
			str:         "45",
			expectedVal: "",
			expectedErr: errors.New("wrong type of data"),
		},
		{
			str:         "",
			expectedVal: "",
			expectedErr: nil,
		},
		{
			str:         "a4bv2c0a2",
			expectedVal: "aaaabvvaa",
			expectedErr: nil,
		},
	}

	// проходим по ним и используем функцию распаковки
	for _, testCase := range testTable {
		result, err := unpacking(testCase.str)

		// если полученный результат совпадает с ожидаемым, то все в порядке
		// если нет, то выводим это
		if result != testCase.expectedVal && err != testCase.expectedErr {
			t.Errorf("Incorrect result. Expect val: %s, err:%v. Got %s, %v\n", testCase.expectedVal, testCase.expectedErr, result, err)
		}
	}

}
