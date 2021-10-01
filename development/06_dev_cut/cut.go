package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	// задаем флаги
	f := flag.Int("f", 0, "'fields' - выбрать поля (колонки)")
	d := flag.String("d", "\t", "'delimiter' - использовать другой разделитель")
	s := flag.Bool("s", false, "'separated' - только строки с разделителем")

	flag.Parse()

	// выбор колонки - обязательное условие
	if *f == 0 {
		fmt.Println("you must use -f with some value > 0")
		return
	}

	// если другие флаги не выбраны
	if *d == "\t" && !*s {
		// последний введенный аргумент - название файла
		fl := os.Args[len(os.Args)-1]
		file, err := ioutil.ReadFile(fl) // читаем файл
		// если не удалось прочитать, то ожидаем ввод
		if err != nil {
			reader := bufio.NewReader(os.Stdin) // или читаем с ввода
			text, err := reader.ReadString('\n')
			if err != nil {
				log.Fatalln(err)
			}
			splitStr := strings.Split(text, "\t")
			// "вырезаем" по указанной колонке
			res := fields(*f, splitStr)
			fmt.Println(res)
		}

		splitStr := strings.Split(string(file), "\n")
		// "вырезаем" по указанной колонке
		res := fields(*f, splitStr)
		fmt.Println(res)
	}

	// если были указаны доп.флаги
	if *d != "\t" || *s {

		if *d != "\t" {
			fl := os.Args[len(os.Args)-1]
			file, err := ioutil.ReadFile(fl) // читаем файл
			if err != nil {
				reader := bufio.NewReader(os.Stdin) // или читаем с ввода
				text, err := reader.ReadString('\n')
				if err != nil {
					log.Fatalln(err)
				}
				splitStr := strings.Split(text, " ")
				// "вырезаем" по указанной колонке или разделителю
				res := delimiter(*f, *d, splitStr)
				fmt.Println(res)
			}

			splitStr := strings.Split(string(file), "\n")
			// "вырезаем" по указанной колонке или разделителю
			res := delimiter(*f, *d, splitStr)
			fmt.Println(res)
		}

		if *s {
			fl := os.Args[len(os.Args)-1]
			file, err := ioutil.ReadFile(fl)
			if err != nil {
				reader := bufio.NewReader(os.Stdin)
				text, err := reader.ReadString('\n')
				if err != nil {
					log.Fatalln(err)
				}
				splitStr := strings.Split(text, " ")
				// "вырезаем" по указанной колонке
				// строки только с разделителем
				res := separated(*f, splitStr)
				fmt.Println(res)
			}

			splitStr := strings.Split(string(file), "\n")
			// "вырезаем" по указанной колонке
			// строки только с разделителем
			res := separated(*f, splitStr)
			fmt.Println(res)

		}
	}
}

// вывод по заданной колонке
func fields(f int, text []string) []string {
	res := []string{}
	// проходим по строкам
	for _, s := range text {
		// разделяем их по tab
		spl := strings.Split(s, "\t")
		// если не удалось разделить
		// выводим всю строку
		if len(spl) == 1 {
			res = append(res, s+"\n")
			continue
		}
		if f <= len(spl) {
			// выводим нужную колонку
			res = append(res, (spl[f-1])+"\n")
		} else {
			// выводим пустую строку, если номер колонки больше, чем их кол-во
			res = append(res, "\n")
		}
	}
	return res
}

// вывод по заданной колонке и разделителю
func delimiter(f int, sep string, text []string) []string {
	res := []string{}
	// проходим по строкам
	for _, s := range text {
		// разделяем каждую по указанному разделителю
		spl := strings.Split(s, sep)
		if f <= len(spl) {
			// выводим нужную колонку
			res = append(res, spl[f-1]+"\n")
		}
	}
	return res
}

// вывод строк только с разделителем
func separated(f int, text []string) []string {
	res := []string{}
	for _, s := range text {
		spl := strings.Split(s, "\t")
		// если не удалось разделить - пропускаем
		if len(spl) == 1 {
			continue
		} else if f <= len(spl) {
			// выводим нужную колонку
			res = append(res, spl[f-1]+"\n")
		}
	}
	return res
}
