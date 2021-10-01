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

// описываем общую функцию проверки ошибок
func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	// проверяем, все ли данные введены
	if len(os.Args) == 1 {
		panic("No file to sort")
	}

	// задаем флаги
	k := flag.Int("k", 1, "указание колонки для сортировки")
	n := flag.Bool("n", false, "сортировать по числовому значению")
	r := flag.Bool("r", false, "сортировать в обратном порядке")
	u := flag.Bool("u", false, "не выводить повторяющиеся строки")

	flag.Parse()

	// если не было добавлено флагов, то запуска обычную сортировку
	if !*n && !*r && !*u && *k == 1 {
		f := os.Args[1]
		file, err := ioutil.ReadFile(f) //читаем данные из файла
		checkError(err)

		splitStr := strings.Split(string(file), "\n")

		fmt.Println(strSort(splitStr))
	}

	// если какой-то флаг был добавлен
	if *k != 1 || *n || *r || *u {
		f := os.Args[2]
		file, err := ioutil.ReadFile(f) //читаем данные из файла
		checkError(err)
		splitStr := strings.Split(string(file), "\n")

		// если указана колонка для сортировки
		if *k != 1 {
			fmt.Println(columnSort(splitStr, *k-1))
		}
		// если необходимо сортировать по числовому значению
		if *n {
			fmt.Println(intSort(splitStr))
		}
		// если нужна сортировка в обратном порядке
		if *r {
			fmt.Println(reverseSort(splitStr))
		}
		// если не нужны повторяющиеся строки
		if *u {
			fmt.Println(sortWithoutRepeat(splitStr))
		}
	}

}

// простая сортировка
func strSort(s []string) []string {
	sort.Strings(s)
	return s
}

// обратная сортировка
func reverseSort(s []string) []string {
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	return s
}

// сортировка без повторяющихся строк
func sortWithoutRepeat(s []string) []string {
	strWithoutRepeat := []string{}
	m := make(map[string]struct{})
	exists := struct{}{}
	// проходим по всем строкам и добавляем их в мап
	// повторяющиемся не будут добавлены
	for _, i := range s {
		m[i] = exists
	}
	// проходим по оставшимся в мапе
	// добавляем их к результату
	for i := range m {
		strWithoutRepeat = append(strWithoutRepeat, i)
	}
	// сортируем их
	sort.Strings(strWithoutRepeat)
	return strWithoutRepeat
}

// сортировка по числовому значению
func intSort(s []string) []string {
	m := make(map[int]string)
	str := []string{}
	for _, i := range s {
		// разбиваем по цифре
		// получаем смещение для первой НЕцифры
		offset := strings.IndexFunc(i, func(r rune) bool { return r < '0' || r > '9' })
		if offset == 0 {
			// если пустая строка
			// т.е. нет цифры
			str = append(str, i)
			break
		}
		// превращаем в цисло
		val, err := strconv.Atoi(i[:offset])
		checkError(err)
		// добавляем в мап, где ключом выступает цифра, а значением - строка
		m[int(val)] = i
	}

	values := []int{}
	for k := range m {
		// добавляем ключи из карты в слайс
		values = append(values, k)
	}

	// сортируем этот слайс
	sort.Ints(values)

	// проходим по отсортированному слайсу
	for _, i := range values {
		// добавляем значение по этому ключу из мапы
		str = append(str, m[i])
	}
	return str
}

// сортировка по заданной колонке
func columnSort(s []string, k int) []string {
	m := make(map[string][]string)
	// проходим циклом по строкам
	for _, i := range s {
		// разделяем их по колонкам
		arrS := strings.Split(i, " ")
		// добавляем ключ в виде строки в нужной колонке
		// значение - вся строка
		m[arrS[k]] = arrS
	}
	keyArr := []string{}
	for key := range m {
		// добавляем в слайс ключи (выбранные колонки)
		keyArr = append(keyArr, key)
	}

	// сортируем их
	sort.Strings(keyArr)

	str := []string{}
	for _, key := range keyArr {
		// проходим по отсортированной колонке
		// добавляем значение всей строки
		ss := m[key][0] + " " + m[key][1]
		str = append(str, ss)
	}

	return str
}
