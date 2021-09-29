package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	if len(os.Args) == 1 {
		panic("No file to sort")
	}

	k := flag.Int("k", 1, "указание колонки для сортировки")
	n := flag.Bool("n", false, "сортировать по числовому значению")
	r := flag.Bool("r", false, "сортировать в обратном порядке")
	u := flag.Bool("u", false, "не выводить повторяющиеся строки")

	flag.Parse()

	if !*n && !*r && !*u && *k == 1 {
		f := os.Args[1]
		file, err := ioutil.ReadFile(f)
		checkError(err)

		splitStr := strings.Split(string(file), "\n")

		fmt.Println(strSort(splitStr))
	}

	if *k != 1 || *n || *r || *u {
		f := os.Args[2]
		file, err := ioutil.ReadFile(f)
		checkError(err)
		splitStr := strings.Split(string(file), "\n")

		if *k != 1 {
			fmt.Println(columnSort(splitStr, *k-1))
		}
		if *n {
			fmt.Println(intSort(splitStr))
		}
		if *r {
			fmt.Println(reverseSort(splitStr))
		}
		if *u {
			fmt.Println(sortWithoutRepeat(splitStr))
		}
	}

}

func strSort(s []string) []string {
	sort.Strings(s)
	return s
}

func reverseSort(s []string) []string {
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	return s
}

func sortWithoutRepeat(s []string) []string {
	strWithoutRepeat := []string{}
	m := make(map[string]struct{})
	exists := struct{}{}
	for _, i := range s {
		m[i] = exists
	}
	for i := range m {
		strWithoutRepeat = append(strWithoutRepeat, i)
	}
	sort.Strings(strWithoutRepeat)
	return strWithoutRepeat
}

func intSort(s []string) []string {
	m := make(map[int]string)
	str := []string{}
	for _, i := range s {
		offset := strings.IndexFunc(i, func(r rune) bool { return r < '0' || r > '9' })
		if offset == 0 {
			str = append(str, i)
			break
		}
		val, err := strconv.Atoi(i[:offset])
		checkError(err)
		m[int(val)] = i
	}

	values := []int{}
	for k := range m {
		values = append(values, k)
	}

	sort.Ints(values)

	for _, i := range values {
		str = append(str, m[i])
	}
	return str
}

func columnSort(s []string, k int) []string {
	m := make(map[string][]string)
	for _, i := range s {
		arrS := strings.Split(i, " ")
		m[arrS[k]] = arrS
	}
	keyArr := []string{}
	for key := range m {
		keyArr = append(keyArr, key)
	}

	sort.Strings(keyArr)

	str := []string{}
	for _, key := range keyArr {
		ss := m[key][0] + " " + m[key][1]
		str = append(str, ss)
	}

	return str
}
