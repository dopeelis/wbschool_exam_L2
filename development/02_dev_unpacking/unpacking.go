package main

import (
	"errors"
	"strconv"
)

// функция распаковки строки
func unpacking(s string) (string, error) {
	// проверка на пустую строку
	if len(s) == 0 {
		return "", nil
	}

	currentChar := []rune(s)[0]
	res := ""

	// если первым идет не буква
	_, err := strconv.ParseInt(string(currentChar), 10, 32)
	if err == nil {
		return "", errors.New("wrong type of data")
	}

	// проходимя по элементам в строке
	for _, c := range s[1:] {
		// выделяем цифру из строки
		cstr := string(c)
		currentCount, err := strconv.Atoi(cstr)
		if err != nil {
			// если не удалось выделить цифру
			// значит буква одна, добавляем ее к результату
			res += string(currentChar)
			currentChar = c
		}
		// проверяем, вдруг цифра 0
		if currentCount == 0 {
			continue
		}
		// если удалось, то повторяем текущее значение буквы
		// количество раз, равное выделенной цифре
		for j := 0; j < currentCount-1; j++ {
			res += string(currentChar)
		}

	}
	// добавляем последнее значение
	res += string(currentChar)
	return res, nil
}
