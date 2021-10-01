package main

import (
	"testing"
)

// функция проверки равенства слайсов
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

// тестирование простого поиска
func TestSimpleSearch(t *testing.T) {
	testTable := []struct {
		phrase   string
		text     []string
		expected []string
	}{
		{
			phrase: "once",
			text: []string{
				"Somebody once told me the world is gonna roll me",
				"I ain't the sharpest tool in the shed",
				"She was looking kind of dumb with her finger and her thumb",
				"In the shape of an 'L' on her forehead",
				"Well the years start coming and they don't stop coming",
			},
			expected: []string{
				"Somebody once told me the world is gonna roll me\n",
			},
		},
		{
			phrase: "if you",
			text: []string{
				"Somebody once told me the world is gonna roll me",
				"I ain't the sharpest tool in the shed",
				"She was looking kind of dumb with her finger and her thumb",
				"In the shape of an 'L' on her forehead",
				"Well the years start coming and they don't stop coming",
				"Fed to the rules and I hit the ground running",
				"Didn't make sense not to live for fun",
				"Your brain gets smart but your head gets dumb",
				"So much to do, so much to see",
				"So what's wrong with taking the back streets?",
				"You'll never know if you don't go",
				"You'll never shine if you don't glow",
			},
			expected: []string{
				"You'll never know if you don't go\n",
				"You'll never shine if you don't glow\n",
			},
		},
	}

	for _, testCase := range testTable {
		result := simpleSearch(testCase.phrase, testCase.text)
		if !Equal(result, testCase.expected) {
			t.Errorf("Incorrect result. Expect: %v, Got %v\n", testCase.expected, result)
		}

	}
}

// тестирование поиска с выводом количества совпадений
func TestCountSearch(t *testing.T) {
	testTable := []struct {
		phrase   string
		text     []string
		expected int
	}{
		{
			phrase: "once",
			text: []string{
				"Somebody once told me the world is gonna roll me",
				"I ain't the sharpest tool in the shed",
				"She was looking kind of dumb with her finger and her thumb",
				"In the shape of an 'L' on her forehead",
				"Well the years start coming and they don't stop coming",
			},
			expected: 1,
		},
		{
			phrase: "if you",
			text: []string{
				"Somebody once told me the world is gonna roll me",
				"I ain't the sharpest tool in the shed",
				"She was looking kind of dumb with her finger and her thumb",
				"In the shape of an 'L' on her forehead",
				"Well the years start coming and they don't stop coming",
				"Fed to the rules and I hit the ground running",
				"Didn't make sense not to live for fun",
				"Your brain gets smart but your head gets dumb",
				"So much to do, so much to see",
				"So what's wrong with taking the back streets?",
				"You'll never know if you don't go",
				"You'll never shine if you don't glow",
			},
			expected: 2,
		},
	}

	for _, testCase := range testTable {
		result := countSearch(testCase.phrase, testCase.text)
		if result != testCase.expected {
			t.Errorf("Incorrect result. Expect: %v, Got %v\n", testCase.expected, result)
		}

	}

}

// тестирование поиска с доп выводом номера строки
func TestLineNumSearch(t *testing.T) {
	testTable := []struct {
		phrase   string
		text     []string
		expected []string
	}{
		{
			phrase: "once",
			text: []string{
				"Somebody once told me the world is gonna roll me",
				"I ain't the sharpest tool in the shed",
				"She was looking kind of dumb with her finger and her thumb",
				"In the shape of an 'L' on her forehead",
				"Well the years start coming and they don't stop coming",
			},
			expected: []string{
				"1 Somebody once told me the world is gonna roll me\n",
			},
		},
		{
			phrase: "if you",
			text: []string{
				"Somebody once told me the world is gonna roll me",
				"I ain't the sharpest tool in the shed",
				"She was looking kind of dumb with her finger and her thumb",
				"In the shape of an 'L' on her forehead",
				"Well the years start coming and they don't stop coming",
				"Fed to the rules and I hit the ground running",
				"Didn't make sense not to live for fun",
				"Your brain gets smart but your head gets dumb",
				"So much to do, so much to see",
				"So what's wrong with taking the back streets?",
				"You'll never know if you don't go",
				"You'll never shine if you don't glow",
			},
			expected: []string{
				"11 You'll never know if you don't go\n",
				"12 You'll never shine if you don't glow\n",
			},
		},
	}

	for _, testCase := range testTable {
		result := lineNumSearch(testCase.phrase, testCase.text)
		if !Equal(result, testCase.expected) {
			t.Errorf("Incorrect result. Expect: %v, Got %v\n", testCase.expected, result)
		}

	}
}

// тестирование поисках строк, НЕ содержащих фразу
func TestInvertSearch(t *testing.T) {
	testTable := []struct {
		phrase   string
		text     []string
		expected []string
	}{
		{
			phrase: "once",
			text: []string{
				"Somebody once told me the world is gonna roll me",
				"I ain't the sharpest tool in the shed",
				"She was looking kind of dumb with her finger and her thumb",
				"In the shape of an 'L' on her forehead",
				"Well the years start coming and they don't stop coming",
			},
			expected: []string{
				"I ain't the sharpest tool in the shed\n",
				"She was looking kind of dumb with her finger and her thumb\n",
				"In the shape of an 'L' on her forehead\n",
				"Well the years start coming and they don't stop coming\n",
			},
		},
		{
			phrase: "if you",
			text: []string{
				"Somebody once told me the world is gonna roll me",
				"I ain't the sharpest tool in the shed",
				"She was looking kind of dumb with her finger and her thumb",
				"In the shape of an 'L' on her forehead",
				"Well the years start coming and they don't stop coming",
				"Fed to the rules and I hit the ground running",
				"Didn't make sense not to live for fun",
				"Your brain gets smart but your head gets dumb",
				"So much to do, so much to see",
				"So what's wrong with taking the back streets?",
				"You'll never know if you don't go",
				"You'll never shine if you don't glow",
			},
			expected: []string{
				"Somebody once told me the world is gonna roll me\n",
				"I ain't the sharpest tool in the shed\n",
				"She was looking kind of dumb with her finger and her thumb\n",
				"In the shape of an 'L' on her forehead\n",
				"Well the years start coming and they don't stop coming\n",
				"Fed to the rules and I hit the ground running\n",
				"Didn't make sense not to live for fun\n",
				"Your brain gets smart but your head gets dumb\n",
				"So much to do, so much to see\n",
				"So what's wrong with taking the back streets?\n",
			},
		},
	}

	for _, testCase := range testTable {
		result := invertSearch(testCase.phrase, testCase.text)
		if !Equal(result, testCase.expected) {
			t.Errorf("Incorrect result. Expect: %v, Got %v\n", testCase.expected, result)
		}

	}
}

// тестирование поиска строк с полным совпадением
func TestFixedtSearch(t *testing.T) {
	testTable := []struct {
		phrase   string
		text     []string
		expected []string
	}{
		{
			phrase: "I ain't the sharpest tool in the shed",
			text: []string{
				"Somebody once told me the world is gonna roll me",
				"I ain't the sharpest tool in the shed",
				"She was looking kind of dumb with her finger and her thumb",
				"In the shape of an 'L' on her forehead",
				"Well the years start coming and they don't stop coming",
			},
			expected: []string{
				"I ain't the sharpest tool in the shed\n",
			},
		},
		{
			phrase: "if you",
			text: []string{
				"Somebody once told me the world is gonna roll me",
				"I ain't the sharpest tool in the shed",
				"She was looking kind of dumb with her finger and her thumb",
				"In the shape of an 'L' on her forehead",
				"Well the years start coming and they don't stop coming",
				"Fed to the rules and I hit the ground running",
				"Didn't make sense not to live for fun",
				"Your brain gets smart but your head gets dumb",
				"So much to do, so much to see",
				"So what's wrong with taking the back streets?",
				"You'll never know if you don't go",
				"You'll never shine if you don't glow",
			},
			expected: []string{},
		},
	}

	for _, testCase := range testTable {
		result := fixedtSearch(testCase.phrase, testCase.text)
		if !Equal(result, testCase.expected) {
			t.Errorf("Incorrect result. Expect: %v, Got %v\n", testCase.expected, result)
		}

	}
}

// тестирование поиска с игнорированием регистра
func TestIgnoreCaseSearch(t *testing.T) {
	testTable := []struct {
		phrase   string
		text     []string
		expected []string
	}{
		{
			phrase: "somebody",
			text: []string{
				"Somebody once told me the world is gonna roll me",
				"I ain't the sharpest tool in the shed",
				"She was looking kind of dumb with her finger and her thumb",
				"In the shape of an 'L' on her forehead",
				"Well the years start coming and they don't stop coming",
			},
			expected: []string{
				"Somebody once told me the world is gonna roll me\n",
			},
		},
		{
			phrase: "NEVER",
			text: []string{
				"Somebody once told me the world is gonna roll me",
				"I ain't the sharpest tool in the shed",
				"She was looking kind of dumb with her finger and her thumb",
				"In the shape of an 'L' on her forehead",
				"Well the years start coming and they don't stop coming",
				"Fed to the rules and I hit the ground running",
				"Didn't make sense not to live for fun",
				"Your brain gets smart but your head gets dumb",
				"So much to do, so much to see",
				"So what's wrong with taking the back streets?",
				"You'll never know if you don't go",
				"You'll never shine if you don't glow",
			},
			expected: []string{
				"You'll never know if you don't go\n",
				"You'll never shine if you don't glow\n",
			},
		},
	}

	for _, testCase := range testTable {
		result := ignoreCaseSearch(testCase.phrase, testCase.text)
		if !Equal(result, testCase.expected) {
			t.Errorf("Incorrect result. Expect: %v, Got %v\n", testCase.expected, result)
		}

	}
}

// тестирование поиска с выводом n строчек после найденной
func TestASearch(t *testing.T) {
	testTable := []struct {
		phrase   string
		number   int
		text     []string
		expected []string
	}{
		{
			phrase: "once",
			number: 2,
			text: []string{
				"Somebody once told me the world is gonna roll me",
				"I ain't the sharpest tool in the shed",
				"She was looking kind of dumb with her finger and her thumb",
				"In the shape of an 'L' on her forehead",
				"Well the years start coming and they don't stop coming",
			},
			expected: []string{
				"Somebody once told me the world is gonna roll me\n",
				"I ain't the sharpest tool in the shed\n",
				"She was looking kind of dumb with her finger and her thumb\n",
			},
		},
		{
			phrase: "me",
			number: 1,
			text: []string{
				"Somebody once told me the world is gonna roll me",
				"I ain't the sharpest tool in the shed",
				"She was looking kind of dumb with her finger and her thumb",
				"In the shape of an 'L' on her forehead",
				"Well the years start coming and they don't stop coming",
				"Fed to the rules and I hit the ground running",
				"Didn't make sense not to live for fun",
				"Your brain gets smart but your head gets dumb",
				"So much to do, so much to see",
				"So what's wrong with taking the back streets?",
				"You'll never know if you don't go",
				"You'll never shine if you don't glow",
				"Hey now, you're an all-star, get your game on, go play",
				"Hey now, you're a rock star, get the show on, get paid",
			},
			expected: []string{
				"Somebody once told me the world is gonna roll me\n",
				"I ain't the sharpest tool in the shed\n",
				"Hey now, you're an all-star, get your game on, go play\n",
				"Hey now, you're a rock star, get the show on, get paid\n",
			},
		},
	}

	for _, testCase := range testTable {
		result := ASearch(testCase.phrase, testCase.text, testCase.number)
		if !Equal(result, testCase.expected) {
			t.Errorf("Incorrect result. Expect: %v, Got %v\n", testCase.expected, result)
		}

	}
}

// тестирование поиска с выводом n строчек до найденной
func TestBSearch(t *testing.T) {
	testTable := []struct {
		phrase   string
		number   int
		text     []string
		expected []string
	}{
		{
			phrase: "once",
			number: 2,
			text: []string{
				"Somebody once told me the world is gonna roll me",
				"I ain't the sharpest tool in the shed",
				"She was looking kind of dumb with her finger and her thumb",
				"In the shape of an 'L' on her forehead",
				"Well the years start coming and they don't stop coming",
			},
			expected: []string{
				"Somebody once told me the world is gonna roll me\n",
			},
		},
		{
			phrase: "now",
			number: 1,
			text: []string{
				"Somebody once told me the world is gonna roll me",
				"I ain't the sharpest tool in the shed",
				"She was looking kind of dumb with her finger and her thumb",
				"In the shape of an 'L' on her forehead",
				"Well the years start coming and they don't stop coming",
				"Fed to the rules and I hit the ground running",
				"Didn't make sense not to live for fun",
				"Your brain gets smart but your head gets dumb",
				"So much to do, so much to see",
				"So what's wrong with taking the back streets?",
				"You'll never know if you don't go",
				"You'll never shine if you don't glow",
				"Hey now, you're an all-star, get your game on, go play",
				"Hey now, you're a rock star, get the show on, get paid",
			},
			expected: []string{
				"So what's wrong with taking the back streets?\n",
				"You'll never know if you don't go\n",
				"You'll never shine if you don't glow\n",
				"Hey now, you're an all-star, get your game on, go play\n",
				"Hey now, you're a rock star, get the show on, get paid\n",
			},
		},
	}

	for _, testCase := range testTable {
		result := BSearch(testCase.phrase, testCase.text, testCase.number)
		if !Equal(result, testCase.expected) {
			t.Errorf("Incorrect result. Expect: %v, Got %v\n", testCase.expected, result)
		}

	}
}

// тестирование поиска с выводом n строчек до и после найденной
func TestCSearch(t *testing.T) {
	testTable := []struct {
		phrase   string
		number   int
		text     []string
		expected []string
	}{
		{
			phrase: "once",
			number: 2,
			text: []string{
				"Somebody once told me the world is gonna roll me",
				"I ain't the sharpest tool in the shed",
				"She was looking kind of dumb with her finger and her thumb",
				"In the shape of an 'L' on her forehead",
				"Well the years start coming and they don't stop coming",
			},
			expected: []string{
				"Somebody once told me the world is gonna roll me\n",
				"I ain't the sharpest tool in the shed\n",
				"She was looking kind of dumb with her finger and her thumb\n",
			},
		},
		{
			phrase: "now",
			number: 1,
			text: []string{
				"Somebody once told me the world is gonna roll me",
				"I ain't the sharpest tool in the shed",
				"She was looking kind of dumb with her finger and her thumb",
				"In the shape of an 'L' on her forehead",
				"Well the years start coming and they don't stop coming",
				"Fed to the rules and I hit the ground running",
				"Didn't make sense not to live for fun",
				"Your brain gets smart but your head gets dumb",
				"So much to do, so much to see",
				"So what's wrong with taking the back streets?",
				"You'll never know if you don't go",
				"You'll never shine if you don't glow",
				"Hey now, you're an all-star, get your game on, go play",
				"Hey now, you're a rock star, get the show on, get paid",
			},
			expected: []string{
				"So what's wrong with taking the back streets?\n",
				"You'll never know if you don't go\n",
				"You'll never shine if you don't glow\n",
				"Hey now, you're an all-star, get your game on, go play\n",
				"Hey now, you're a rock star, get the show on, get paid\n",
			},
		},
	}

	for _, testCase := range testTable {
		result := CSearch(testCase.phrase, testCase.text, testCase.number)
		if !Equal(result, testCase.expected) {
			t.Errorf("Incorrect result. Expect: %v, Got %v\n", testCase.expected, result)
		}

	}
}
