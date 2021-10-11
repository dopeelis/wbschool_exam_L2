package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	// объявляем флаги
	A := flag.Int("A", 0, "'after' печатать +N строк после совпадения")
	B := flag.Int("B", 0, "'before' печатать +N строк до совпадения")
	C := flag.Int("C", 0, "'context' (A+B) печатать ±N строк вокруг совпадения")
	c := flag.Bool("c", false, "'count' (количество строк)")
	i := flag.Bool("i", false, "'ignore-case' (игнорировать регистр)")
	v := flag.Bool("v", false, "'invert' (вместо совпадения, исключать)")
	F := flag.Bool("F", false, "'fixed', точное совпадение со строкой")
	n := flag.Bool("n", false, "'line num', печатать номер строки")

	flag.Parse()

	args := flag.Args()

	if len(args) < 2 {
		log.Fatalln("usage: [flags] [pattern or string] [file]")
	}

	// выделяем из запроса фразу, которую будем искать
	slicePhrase := args[:len(args)-1]
	phrase := strings.Join(slicePhrase, " ")

	// читаем данные из файла
	file, err := ioutil.ReadFile(args[len(args)-1])
	if err != nil {
		log.Fatalln(err)
	}

	splitString := strings.Split(string(file), "\n")

	// запускаем функцию и выводим на экран
	fmt.Println(Grep(phrase, splitString, *A, *B, *C, *c, *i, *v, *F, *n))
}

// Find функция поиска элемента в слайсе
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

// Grep функция поиска фразы или строки в файле с применением доп.условий
func Grep(phrase string, text []string, A, B, C int, c, i, v, F, n bool) []string {
	// переменная для флага -с
	var stringCount = 0
	var result []string
	// условие сравнения
	var condition bool

	// согласовывем флаги А В и С
	// чтобы работало бОльшее значение
	if C != 0 && C > A && C > B {
		A = C
		B = C
	} else if C != 0 && (C > A || C > B) {
		if C > A {
			A = C
		} else if C > B {
			B = C
		}
	}

	// проходим по каждой строке
	for index, str := range text {
		// если применен -i, убираем регистр
		if i {
			str = strings.ToLower(str)
			phrase = strings.ToLower(phrase)
		}
		// меняем условие в зависимсоти от переданных флагов
		if F {
			// если нужно искать полное совпедаение
			condition = phrase == str
		} else {
			// просто фразу в строке
			condition = strings.Contains(str, phrase)
		}
		// если -v
		if v {
			condition = !condition
		}
		// если условие выполняется
		if condition {
			// увеличиваем счетчик положительных результатов
			stringCount++
			// выполняем, если задан -B
			if B != 0 {
				if index <= B-1 { // если вначале текста
					for j := index; j >= 0; j-- { // выводи столько строк, сколько до начала
						if !Find(result, text[index-j]) { // чтобы не выводить повторы
							// если нужно доабвлять номер строки в выводе
							if n {
								result = append(result, strconv.Itoa(index-j+1)+" "+text[index-j])
							} else {
								result = append(result, text[index-j])
							}
						} else {
							continue
						}
					}
				} else {
					for j := B; j >= 0; j-- { // выводим столько строчек, сколько задано
						if !Find(result, text[index-j]) { // чтобы не выводить повторы
							// если нужно доабвлять номер строки в выводе
							if n {
								result = append(result, strconv.Itoa(index-j+1)+" "+text[index-j])
							} else {
								result = append(result, text[index-j])
							}
						} else {
							continue
						}
					}
				}
			}

			// выполняем , если задан -A
			if A != 0 {
				if index > len(text)-1-A { // если в конце текста
					for j := 0; j < len(text)-index+1; j++ { // выводим столько строк, сколько осталось в конце
						if !Find(result, text[index]) { // чтобы не выводить повторы
							// если нужно доабвлять номер строки в выводе
							if n {
								result = append(result, strconv.Itoa(index+1)+" "+text[index])
							} else {
								result = append(result, text[index])
							}
							index++
						} else {
							index++
						}
					}
				} else {
					for j := 0; j < A+1; j++ { // выводим столько строчек, сколько задано
						if !Find(result, text[index]) { // чтобы не выводить повторы
							// если нужно доабвлять номер строки в выводе
							if n {
								result = append(result, strconv.Itoa(index+1)+" "+text[index])
							} else {
								result = append(result, text[index])
							}
							index++
						} else {
							index++
						}

					}
				}
			}
			// если оба не заданы, то просто добавляем совпадение в результат
			if A == 0 && B == 0 {
				if n {
					result = append(result, strconv.Itoa(index+1)+" "+text[index])
				} else {
					result = append(result, text[index])
				}
			}
		}
	}

	// если нужно вывести соличество совпадений
	if c {
		count := strconv.Itoa(stringCount)
		result = []string{count}
	}

	// финальное предствление ответа
	var finalResult []string
	for _, s := range result {
		// применяем ко всем элементам \n
		finalResult = append(finalResult, s+"\n")
	}
	return finalResult
}
