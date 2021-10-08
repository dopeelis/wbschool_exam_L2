package main

import (
	"bufio"
	"fmt"
	"github.com/shirou/gopsutil/v3/process"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"syscall"
)

// checkErr функция для проверки ошибок
func checkErr(err error) {
	if err != nil {
		log.Fatalln()
	}
}

func main() {
	// бесконечный цикл для чтения команд
	for {
		reader := bufio.NewReader(os.Stdin)
		request, err := reader.ReadString('\n')
		checkErr(err)
		splitRequest := strings.Split(request, " ")
		// распознаем команду
		switch splitRequest[0] {
		// завершаем программу, если exit
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

// возвращаем значение текущего пути
func pwd() string {
	wd, err := os.Getwd()
	checkErr(err)
	return wd
}

// меняем директорию
func cd(location []string) {
	// если нужно вернуться на уровень выше
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
		// если нужно двигаться дальше
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
	// показываем все файлы в текущей директории
	if s[0] == "*\n" || string(s[0][0]) == "*" {
		if len(s[0]) == 1 {
			files, err := ioutil.ReadDir(".")
			checkErr(err)

			for _, file := range files {
				fmt.Println(file.Name(), "| is dir:", file.IsDir())
			}
		} else {
			// если указано конкретное расширение
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
		// повторяем введенную строку
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

// сигнал для остановки конкретного приложения
func kill(s []string) {
	var prs []string
	var signal syscall.Signal
	var strSignal string
	for i, w := range s[0] {
		if i == 0 {
			if string(w) == "S" || string(w) == "-" {
				strSignal = s[0]
				prs = s[1:]
			} else {
				signal = syscall.SIGTERM
				prs = s
			}
			break
		}
		break
	}

	switch strSignal {
	case "SIGKILL":
		signal = syscall.SIGKILL
	case "SIGINT":
		signal = syscall.SIGINT
	case "SIGTERM":
		signal = syscall.SIGTERM
	case "SIGQUIT":
		signal = syscall.SIGQUIT
	case "SIGABRT":
		signal = syscall.SIGABRT
	case "SIGHUP":
		signal = syscall.SIGHUP
	case "-9":
		signal = syscall.SIGKILL
	case "-1":
		signal = syscall.SIGHUP
	}

	processes, err := process.Processes()
	if err != nil {
		log.Println(err)
	}
	for _, p := range processes {
		n, err := p.Name()
		if err != nil {
			log.Println(err)
		}
		for _, pp := range prs {
			if n == pp {
				err := p.SendSignal(signal)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
	log.Println("process not found")
}

// возвращаем все текущие процессы
func prs() []int32 {
	processes, err := process.Processes()
	if err != nil {
		log.Println(err)
	}
	var pids []int32

	for _, p := range processes {
		pids = append(pids, p.Pid)
	}
	return pids
}
