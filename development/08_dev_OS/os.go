package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/mitchellh/go-ps"
)

func checkErr(err error) {
	if err != nil {
		log.Fatalln()
	}
}

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		request, err := reader.ReadString('\n')
		checkErr(err)
		splitRequest := strings.Split(request, " ")
		switch splitRequest[0] {
		case "exit\n":
			return
		case "cd":
			cd(splitRequest[1:])
		case "pwd\n":
			fmt.Println(pwd())
		case "echo":
			echo(splitRequest[1:])
		case "kill":
			kill(splitRequest[1:])
		case "ps\n":
			prs()
			fmt.Println(prs())
		}

	}
}

func pwd() string {
	wd, err := os.Getwd()
	checkErr(err)
	return wd
}

func cd(location []string) {
	if location[0] == "-\n" {
		var currDir string
		splitDir := strings.Split(pwd(), "/")
		for _, i := range splitDir[:len(splitDir)-1] {
			currDir += i + "/"
		}
		err := os.Chdir(currDir)
		checkErr(err)
		fmt.Println(pwd())
	} else {
		var path string
		for _, w := range location {
			path += w
		}

		err := os.Chdir(pwd() + "/" + path[:len(path)-1])
		checkErr(err)
		fmt.Println(pwd())
	}

}

func echo(s []string) {
	if s[0] == "*\n" || string(s[0][0]) == "*" {
		if len(s[0]) == 1 {
			files, err := ioutil.ReadDir(".")
			checkErr(err)

			for _, file := range files {
				fmt.Println(file.Name(), "| is dir:", file.IsDir())
			}
		} else {
			files, err := ioutil.ReadDir(".")
			checkErr(err)
			for _, file := range files {
				if file.IsDir() {
					continue
				} else {
					splName := strings.Split(file.Name(), ".")
					if "."+splName[1]+"\n" == string(s[0][1:]) {
						fmt.Println(file.Name())
					} else {
						continue
					}
				}

			}
		}

	} else {
		var out string
		for _, w := range s {
			if string(w[0]) == "\\" {
				if string(w[1]) == "n" {
					out += "\n"
					out += string(w[2:])
				}
				if string(w[1]) == "t" {
					out += "\n"
					out += "\t"
					out += string(w[2:])
				}
			} else {
				out += w + " "
			}
		}
		fmt.Println(out)
	}
}

func kill(s []string) {

}

func prs() []ps.Process {
	prs, err := ps.Processes()
	checkErr(err)
	return prs
}
