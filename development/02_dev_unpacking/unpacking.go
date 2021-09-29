package main

import (
	"errors"
	"strconv"
)

func unpacking(s string) (string, error) {
	if len(s) == 0 {
		return "", nil
	}
	currentChar := []rune(s)[0]
	res := ""
	_, err := strconv.ParseInt(string(currentChar), 10, 32)
	if err == nil {
		return "", errors.New("wrong type of data")
	}

	for _, c := range s {
		if c == currentChar {
			continue
		}
		cstr := string(c)
		currentCount, err := strconv.ParseInt(cstr, 10, 32)
		i := int(currentCount)
		if err != nil {
			res += string(currentChar)
			currentChar = c
		}
		for j := 0; j < i-1; j++ {
			res += string(currentChar)
		}

	}
	res += string(currentChar)
	return res, nil
}
