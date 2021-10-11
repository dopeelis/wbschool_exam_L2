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

	args := flag.Args()

	// выбор колонки - обязательное условие
	if *f == 0 {
		log.Fatalln("you must use -f with some value > 0")
	}

	// если файл не добавлен
	if len(args) == 0 {
		for {
			reader := bufio.NewReader(os.Stdin) //читаем с ввода
			text, err := reader.ReadString('\n')
			if err != nil {
				log.Fatalln(err)
			}
			// "вырезаем" по указанной колонке
			res := Cut(text, *f, *d, *s)
			fmt.Println(res)
		}
		return
	}

	// если есть файл, читаем его
	fl := args[len(args)-1]
	file, err := ioutil.ReadFile(fl) // читаем файл
	// если не удалось прочитать, то ожидаем ввод
	if err != nil {
		log.Fatalln(err)
	}

	splitStr := strings.Split(string(file), "\n")
	// проходим по всем строкам и для каждой вызываем метод Cut
	for _, i := range splitStr {
		res := Cut(i, *f, *d, *s)
		fmt.Println(res)
	}

}

func Cut(str string, f int, d string, s bool) string {
	// если добавлено -s
	if s {
		if !strings.Contains(str, d) {
			return ""
		}
	}
	// если все содержит, то делим строку по разделителю
	spl := strings.Split(str, d)
	// если номер колонки меньше, чем всего их в строке
	if f <= len(spl) {
		// выводим нужную колонку
		return spl[f-1] + "\n"
	} else {
		// если нет, выводим пустую строку
		return ""
	}
}
