package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	An := flag.Int("A", 0, "'after' печатать +N строк после совпадения")
	Bn := flag.Int("B", 0, "'before' печатать +N строк до совпадения")
	C := flag.Int("C", 0, "'context' (A+B) печатать ±N строк вокруг совпадения")
	c := flag.Bool("c", false, "'count' (количество строк)")
	i := flag.Bool("i", false, "'ignore-case' (игнорировать регистр)")
	v := flag.Bool("v", false, "'invert' (вместо совпадения, исключать)")
	F := flag.Bool("F", false, "'fixed', точное совпадение со строкой")
	n := flag.Bool("n", false, "'line num', печатать номер строки")

	flag.Parse()

	if len(os.Args) < 3 {
		fmt.Println("No file or expression")
		return
	}

	if *An == 0 && *Bn == 0 && *C == 0 && !*c && !*i && !*v && !*F && !*n {
		f := os.Args[len(os.Args)-1]
		phrase := os.Args[1 : len(os.Args)-1]
		allPhrase := strings.Join(phrase, " ")
		file, err := ioutil.ReadFile(f)
		if err != nil {
			log.Fatalln(err)
		}

		splitStr := strings.Split(string(file), "\n")

		res := simpleSearch(allPhrase, splitStr)
		fmt.Println(res)
		return
	}

	if *An != 0 || *Bn != 0 || *C != 0 || *c || *i || *v || *F || *n {
		if *An != 0 || *Bn != 0 || *C != 0 {
			f := os.Args[len(os.Args)-1]
			phrase := os.Args[3 : len(os.Args)-1]
			allPhrase := strings.Join(phrase, " ")
			file, err := ioutil.ReadFile(f)
			if err != nil {
				log.Fatalln(err)
			}

			splitStr := strings.Split(string(file), "\n")

			if *An != 0 {
				res := ASearch(allPhrase, splitStr, *An)
				fmt.Println(res)
			}
			if *Bn != 0 {
				res := BSearch(allPhrase, splitStr, *Bn)
				fmt.Println(res)
			}
			if *C != 0 {
				res := CSearch(allPhrase, splitStr, *C)
				fmt.Println(res)
			}

		} else {
			f := os.Args[len(os.Args)-1]
			phrase := os.Args[2 : len(os.Args)-1]
			allPhrase := strings.Join(phrase, " ")
			file, err := ioutil.ReadFile(f)
			if err != nil {
				log.Fatalln(err)
			}
			splitStr := strings.Split(string(file), "\n")

			if *c {
				res := countSearch(allPhrase, splitStr)
				fmt.Println(res)
			}
			if *i {
				res := ignoreCaseSearch(allPhrase, splitStr)
				fmt.Println(res)
			}
			if *v {
				res := invertSearch(allPhrase, splitStr)
				fmt.Println(res)
			}
			if *F {
				res := fixedtSearch(allPhrase, splitStr)
				fmt.Println(res)
			}
			if *n {
				res := lineNumSearch(allPhrase, splitStr)
				fmt.Println(res)
			}
		}
	}
}

func simpleSearch(phrase string, text []string) []string {
	res := []string{}
	for _, i := range text {
		if strings.Contains(i, phrase) {
			res = append(res, i+"\n")
		}
	}
	return res
}

func countSearch(phrase string, text []string) int {
	counter := 0
	for _, i := range text {
		if strings.Contains(i, phrase) {
			counter++
		}
	}
	return counter
}

func lineNumSearch(phrase string, text []string) []string {
	res := []string{}
	for i, s := range text {
		if strings.Contains(s, phrase) {
			ss := strconv.Itoa(i+1) + " " + s + "\n"
			res = append(res, ss)
		}
	}
	return res
}

func invertSearch(phrase string, text []string) []string {
	res := []string{}
	for _, i := range text {
		if !strings.Contains(i, phrase) {
			res = append(res, i+"\n")
		}
	}
	return res
}

func fixedtSearch(phrase string, text []string) []string {
	res := []string{}
	for _, i := range text {
		if phrase == i {
			res = append(res, i+"\n")
		}
	}
	return res
}

func ignoreCaseSearch(phrase string, text []string) []string {
	res := []string{}
	phrase = strings.ToLower(phrase)
	for _, i := range text {
		if strings.Contains(strings.ToLower(i), phrase) {
			res = append(res, i+"\n")
		}
	}
	return res
}

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

func ASearch(phrase string, text []string, n int) []string {
	res := []string{}
	for i, s := range text {
		if strings.Contains(s, phrase) {
			if i > len(text)-1-n {
				for j := 0; j < len(text)-i+1; j++ {
					if !Find(res, text[i]+"\n") {
						res = append(res, text[i]+"\n")
						i++
					} else {
						i++
					}
				}
			} else {
				for j := 0; j < n+1; j++ {
					if !Find(res, text[i]+"\n") {
						res = append(res, text[i]+"\n")
						i++
					} else {
						i++
					}

				}
			}
		}
	}
	return res
}

func BSearch(phrase string, text []string, n int) []string {
	res := []string{}
	for i, s := range text {
		if strings.Contains(s, phrase) {
			if i <= n-1 {
				for j := i; j >= 0; j-- {
					if !Find(res, text[i-j]+"\n") {
						res = append(res, text[i-j]+"\n")
					} else {
						continue
					}
				}
			} else {
				for j := n; j >= 0; j-- {
					if !Find(res, text[i-j]+"\n") {
						res = append(res, text[i-j]+"\n")
					} else {
						continue
					}
				}
			}
		}
	}
	return res
}

func CSearch(phrase string, text []string, n int) []string {
	res := BSearch(phrase, text, n)
	aSearch := ASearch(phrase, text, n)

	for _, s := range aSearch {
		if !Find(res, s) {
			res = append(res, s)
		}
	}

	return res
}
