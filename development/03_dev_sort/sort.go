package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

// checkError общая функцию проверки ошибок
func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// Find функция для нахождения элемента в слайсе
func Find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		} else {
			continue
		}
	}
	return false
}

func main() {
	// задаем флаги
	k := flag.Int("k", 1, "указание колонки для сортировки")
	n := flag.Bool("n", false, "сортировать по числовому значению")
	r := flag.Bool("r", false, "сортировать в обратном порядке")
	u := flag.Bool("u", false, "не выводить повторяющиеся строки")

	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		log.Fatalln("usage: [flags] [file]")
	}

	file, err := ioutil.ReadFile(args[0])
	checkError(err)

	splitStr := strings.Split(string(file), "\n")

	fmt.Println(strSort(splitStr, *k-1, *n, *r, *u))
}

func strSort(arr []string, k int, n, r, u bool) []string {
	var trueKey string
	// промежуточная мапа
	interimMap := make(map[string]string)
	// ключи для промежуточной мапы
	var keyArr []string
	var intKeyArr []int
	// конечный результат
	var allStr []string

	// проходим циклом по строкам
	for _, i := range arr {
		// разделяем каждую
		arrS := strings.Split(i, " ")

		// проверяем, не задана ли колонка больше, чем всего их
		if k > len(arrS) {
			// если да, то сортировать будем по первому элементу
			trueKey = arrS[0]
		} else {
			// если нет, то по заданному
			trueKey = arrS[k]
		}

		// если нужно сортировать по числам
		if n {
			// находим смещение НЕцифры у элемента с индексом "заданный номер колонки -1"
			offset := strings.IndexFunc(trueKey, func(r rune) bool { return r < '0' || r > '9' })
			// задаем индекс для строк без чисел (выводятся вперед)
			notIntIndex := 0
			// если числа нет
			if offset == 0 {
				// уменьшвем индекс
				notIntIndex--
				// переводим его в строку
				integer := strconv.Itoa(notIntIndex)
				// ставим ключом в мапе, вся строка является значением
				interimMap[integer] = i
				break
			}
			// если была цифра, то делаем ее в виде строки ключом
			integer := trueKey[:offset]
			// всю строку делаем значением
			interimMap[integer] = i
		} else {
			//	если флаг n не задан, то добвялем выбранную колонку ключом, всю строку значением
			interimMap[trueKey] = i
		}
	}

	//теперь мы имеем промежуточную мапу, где ключи - элемент, по которому надо сортировать
	// а значение - сама строка

	// проходим по ключам в мапе
	for key := range interimMap {
		if n {
			intKey, err := strconv.Atoi(key)
			checkError(err)
			intKeyArr = append(intKeyArr, intKey)
		} else {
			// и добавляем их в слайс ключей для дальнейшей сортировки
			keyArr = append(keyArr, key)
		}
	}

	// если нужно сортировать реверсивно
	if r {
		// выполняем сортировку слайса ключей в обратном порядке

		if n {
			sort.Sort(sort.Reverse(sort.IntSlice(intKeyArr)))
			for _, i := range intKeyArr {
				keyArr = append(keyArr, strconv.Itoa(i))
			}
		} else {
			sort.Sort(sort.Reverse(sort.StringSlice(keyArr)))
		}
	} else {
		//	если обычная сортировка
		// то просто сортируем ключи

		if n {
			sort.Ints(intKeyArr)
			for _, i := range intKeyArr {
				keyArr = append(keyArr, strconv.Itoa(i))
			}
		} else {
			sort.Strings(keyArr)
		}

	}

	// проходим по каждому элементу в отсортированном слайсе ключей
	for _, key := range keyArr {
		str := interimMap[key]
		// проверяем, нужно ли выводить строки с повторениями
		// если нет
		if u {
			// проверяем, если ли строка уже в конечном ответе
			if Find(allStr, str) {
				// если да, то пропускаем
				continue
			} else {
				// если нет,то добавляем
				allStr = append(allStr, str)
			}
		} else {
			// если нужны все строки
			allStr = append(allStr, str)
		}
	}
	return allStr
}
