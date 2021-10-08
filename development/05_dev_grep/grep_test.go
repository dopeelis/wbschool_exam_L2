package main

import (
	"fmt"
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

func TestGrep(t *testing.T) {
	testTable := []struct {
		phrase        string
		text          []string
		A, B, C       int
		c, i, v, F, n bool
		expected      []string
	}{
		// grep without arguments
		{
			phrase: "once",
			text: []string{
				"Somebody once told me the world is gonna roll me",
				"I ain't the sharpest tool in the shed",
				"She was looking kind of dumb with her finger and her thumb",
				"In the shape of an 'L' on her forehead",
				"Well the years start coming and they don't stop coming",
			},
			A: 0,
			B: 0,
			C: 0,
			c: false,
			i: false,
			v: false,
			F: false,
			n: false,
			expected: []string{
				"Somebody once told me the world is gonna roll me\n",
			},
		},
		// grep with -c
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
			A:        0,
			B:        0,
			C:        0,
			c:        true,
			i:        false,
			v:        false,
			F:        false,
			n:        false,
			expected: []string{"2\n"},
		},
		// grep with -n
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
			A: 0,
			B: 0,
			C: 0,
			c: false,
			i: false,
			v: false,
			F: false,
			n: true,
			expected: []string{
				"11 You'll never know if you don't go\n",
				"12 You'll never shine if you don't glow\n",
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
			A:        0,
			B:        0,
			C:        0,
			c:        true,
			i:        false,
			v:        false,
			F:        false,
			n:        true,
			expected: []string{"2\n"},
		},
		{
			phrase: "once",
			text: []string{
				"Somebody once told me the world is gonna roll me",
				"I ain't the sharpest tool in the shed",
				"She was looking kind of dumb with her finger and her thumb",
				"In the shape of an 'L' on her forehead",
				"Well the years start coming and they don't stop coming",
			},
			A: 0,
			B: 0,
			C: 0,
			c: false,
			i: false,
			v: true,
			F: false,
			n: false,
			expected: []string{
				"I ain't the sharpest tool in the shed\n",
				"She was looking kind of dumb with her finger and her thumb\n",
				"In the shape of an 'L' on her forehead\n",
				"Well the years start coming and they don't stop coming\n",
			},
		},
		{
			phrase: "I ain't the sharpest tool in the shed",
			text: []string{
				"Somebody once told me the world is gonna roll me",
				"I ain't the sharpest tool in the shed",
				"She was looking kind of dumb with her finger and her thumb",
				"In the shape of an 'L' on her forehead",
				"Well the years start coming and they don't stop coming",
			},
			A: 0,
			B: 0,
			C: 0,
			c: false,
			i: false,
			v: false,
			F: true,
			n: false,
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
			A:        0,
			B:        0,
			C:        0,
			c:        false,
			i:        false,
			v:        false,
			F:        true,
			n:        false,
			expected: []string{},
		},
		{
			phrase: "somebody",
			text: []string{
				"Somebody once told me the world is gonna roll me",
				"I ain't the sharpest tool in the shed",
				"She was looking kind of dumb with her finger and her thumb",
				"In the shape of an 'L' on her forehead",
				"Well the years start coming and they don't stop coming",
			},
			A: 0,
			B: 0,
			C: 0,
			c: false,
			i: true,
			v: false,
			F: false,
			n: false,
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
			A: 0,
			B: 0,
			C: 0,
			c: false,
			i: true,
			v: false,
			F: false,
			n: true,
			expected: []string{
				"11 You'll never know if you don't go\n",
				"12 You'll never shine if you don't glow\n",
			},
		},
		// ASearch
		{
			phrase: "once",
			text: []string{
				"Somebody once told me the world is gonna roll me",
				"I ain't the sharpest tool in the shed",
				"She was looking kind of dumb with her finger and her thumb",
				"In the shape of an 'L' on her forehead",
				"Well the years start coming and they don't stop coming",
			},
			A: 2,
			B: 0,
			C: 0,
			c: false,
			i: false,
			v: false,
			F: false,
			n: false,
			expected: []string{
				"Somebody once told me the world is gonna roll me\n",
				"I ain't the sharpest tool in the shed\n",
				"She was looking kind of dumb with her finger and her thumb\n",
			},
		},
		// BSearch
		{
			phrase: "once",
			text: []string{
				"Somebody once told me the world is gonna roll me",
				"I ain't the sharpest tool in the shed",
				"She was looking kind of dumb with her finger and her thumb",
				"In the shape of an 'L' on her forehead",
				"Well the years start coming and they don't stop coming",
			},
			A: 0,
			B: 2,
			C: 0,
			c: false,
			i: false,
			v: false,
			F: false,
			n: true,
			expected: []string{
				"1 Somebody once told me the world is gonna roll me\n",
			},
		},
		{
			phrase: "now",
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
			A: 0,
			B: 1,
			C: 0,
			c: false,
			i: false,
			v: false,
			F: false,
			n: false,
			expected: []string{
				"So what's wrong with taking the back streets?\n",
				"You'll never know if you don't go\n",
				"You'll never shine if you don't glow\n",
				"Hey now, you're an all-star, get your game on, go play\n",
				"Hey now, you're a rock star, get the show on, get paid\n",
			},
		},
		//CSearch
		{
			phrase: "once",
			text: []string{
				"Somebody once told me the world is gonna roll me",
				"I ain't the sharpest tool in the shed",
				"She was looking kind of dumb with her finger and her thumb",
				"In the shape of an 'L' on her forehead",
				"Well the years start coming and they don't stop coming",
			},
			A: 0,
			B: 0,
			C: 2,
			c: false,
			i: false,
			v: false,
			F: false,
			n: false,
			expected: []string{
				"Somebody once told me the world is gonna roll me\n",
				"I ain't the sharpest tool in the shed\n",
				"She was looking kind of dumb with her finger and her thumb\n",
			},
		},
		{
			phrase: "noW",
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
			A: 0,
			B: 0,
			C: 1,
			c: false,
			i: true,
			v: false,
			F: false,
			n: false,
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
		result := Grep(testCase.phrase, testCase.text, testCase.A, testCase.B, testCase.C,
			testCase.c, testCase.i, testCase.v, testCase.F, testCase.n)
		if !Equal(result, testCase.expected) {
			fmt.Println("Params: ", testCase.A, testCase.B, testCase.C, testCase.c,
				testCase.i, testCase.v, testCase.F, testCase.n)
			t.Errorf("Incorrect result. Expect: %v, Got %v\n", testCase.expected, result)
		}
	}
}
