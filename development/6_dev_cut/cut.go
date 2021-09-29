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
	f := flag.Int("f", 0, "'fields' - выбрать поля (колонки)")
	d := flag.String("d", "\t", "'delimiter' - использовать другой разделитель")
	s := flag.Bool("s", false, "'separated' - только строки с разделителем")

	flag.Parse()

	if *f == 0 {
		fmt.Println("you must use -f with some value > 0")
		return
	}

	if *d == "\t" && !*s {
		fl := os.Args[len(os.Args)-1]
		file, err := ioutil.ReadFile(fl)
		if err != nil {
			reader := bufio.NewReader(os.Stdin)
			text, err := reader.ReadString('\n')
			if err != nil {
				log.Fatalln(err)
			}
			splitStr := strings.Split(text, "\t")
			res := fields(*f, splitStr)
			fmt.Println(res)
		}

		splitStr := strings.Split(string(file), "\n")
		res := fields(*f, splitStr)
		fmt.Println(res)
	}

	if *d != "\t" || *s {

		if *d != "\t" {
			fl := os.Args[len(os.Args)-1]
			file, err := ioutil.ReadFile(fl)
			if err != nil {
				reader := bufio.NewReader(os.Stdin)
				text, err := reader.ReadString('\n')
				if err != nil {
					log.Fatalln(err)
				}
				splitStr := strings.Split(text, " ")
				res := delimiter(*f, *d, splitStr)
				fmt.Println(res)
			}

			splitStr := strings.Split(string(file), "\n")
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
				res := separated(*f, splitStr)
				fmt.Println(res)
			}

			splitStr := strings.Split(string(file), "\n")
			res := separated(*f, splitStr)
			fmt.Println(res)

		}
	}
}

func fields(f int, text []string) []string {
	res := []string{}
	for _, s := range text {
		spl := strings.Split(s, "\t")
		if len(spl) == 1 {
			res = append(res, s+"\n")
			continue
		}
		if f <= len(spl) {
			res = append(res, (spl[f-1])+"\n")
		} else {
			res = append(res, "\n")
		}
	}
	return res
}

func delimiter(f int, sep string, text []string) []string {
	res := []string{}
	for _, s := range text {
		spl := strings.Split(s, sep)
		if f <= len(spl) {
			res = append(res, spl[f-1]+"\n")
		}
	}
	return res
}

func separated(f int, text []string) []string {
	res := []string{}
	for _, s := range text {
		spl := strings.Split(s, "\t")
		if len(spl) == 1 {
			continue
		} else {
			if f <= len(spl) {
				res = append(res, spl[f-1]+"\n")
			}
		}
	}
	return res
}
