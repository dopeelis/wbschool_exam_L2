package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// checkError функция для проверки ошибок
func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// HTTPGet функция для получения страницы по url с возможностью передать timeout
func HTTPGet(url string, timeout time.Duration) (content []byte, err error) {
	request, err := http.NewRequest("GET", url, nil)
	checkError(err)

	client := &http.Client{
		Timeout: timeout,
	}

	// делаем запрос
	response, err := client.Do(request)
	checkError(err)

	defer func() {
		err := response.Body.Close()
		checkError(err)

	}()

	// если статус не ОК, возвращаем ошибку
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("Error.Status: %s", response.Status)
	}

	// возвращаем тело запроса
	return ioutil.ReadAll(response.Body)
}

func main() {
	// задаем страндартные фраги
	url := flag.String("url", "", "url")
	timeout := flag.Duration("timeout", 5*time.Second, "timeout")
	outputPath := flag.String("output", "test.html", "output path")

	flag.Parse()

	// получаем информацию с основной страницы
	content, err := HTTPGet(*url, *timeout)
	checkError(err)

	// сохраняем основную страницу
	err = ioutil.WriteFile(*outputPath, content, 0666)
	checkError(err)

	// сохраняем все страницы с сайта
	//WriteFile(*url, LinkScrape(*url))
}

// Если нужны ВСЕ страницы с сайта
// LinkScrape получаем все ссылки с сайта
//func LinkScrape(url string) []string {
//	resp, err := http.Get(url)
//	checkError(err)
//	doc, err := goquery.NewDocumentFromReader(resp.Body)
//	checkError(err)
//
//	var links []string
//
//	doc.Find("body a").Each(func(index int, item *goquery.Selection) {
//		linkTag := item
//		link, _ := linkTag.Attr("href")
//		links = append(links, link)
//	})
//	return links
//}
//
//// WriteFile записываем все в html файлы
//func WriteFile(url string, links []string) {
//	for i, l := range links {
//		resp, err := http.Get(url + l)
//		if err != nil {
//			fmt.Println("failed")
//
//		}
//		defer func() {
//			err := resp.Body.Close()
//			checkError(err)
//		}()
//
//		f, err := os.Create(strconv.Itoa(i) + ".html")
//		if err != nil {
//			fmt.Println("creating file failed")
//		}
//		defer func() {
//			err := f.Close()
//			checkError(err)
//		}()
//
//		_, err = io.Copy(f, resp.Body)
//		checkError(err)
//
//	}
//}
