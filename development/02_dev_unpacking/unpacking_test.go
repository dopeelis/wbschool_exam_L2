package main

import (
	"errors"
	"testing"
)

func TestUnpacking(t *testing.T) {
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
	}

	for _, testCase := range testTable {
		result, err := unpacking(testCase.str)

		if result != testCase.expectedVal && err != testCase.expectedErr {
			t.Errorf("Incorrect result. Expect val: %s, err:%v. Got %s, %v\n", testCase.expectedVal, testCase.expectedErr, result, err)
		}
	}

}
