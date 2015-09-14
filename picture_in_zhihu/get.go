package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"os"
	"strings"
)

const isTitleNeed = true
const fileName = "url.txt"

func getUrls(url string) []string {
	doc, _ := goquery.NewDocument(url)
	fmt.Println(doc.Find("title").Text())
	body := doc.Find("body")
	urls := []string{}
	body.Find("div.zm-item").Each(func(i int, s *goquery.Selection) {
		h2 := s.Find("h2")
		question_url, _ := h2.Find("a").Attr("href")
		if strings.HasPrefix(question_url, "/question/") {
			question_url = "http://www.zhihu.com" + question_url
			getImgUrl(question_url)
		}
		urls = append(urls, question_url)
	})
	return urls
}

func getImgUrl(url string) {
	var temp string
	doc, _ := goquery.NewDocument(url)
	dirName := doc.Find("title").Text()
	name := strings.Split(dirName, "-")[0]
	slice := strings.Split(url, "/")
	length := len(slice)
	num := slice[length-1]
	if isTitleNeed {
		writeImgUrl(name + "\n")
	}
	os.Mkdir("./pictures/" + num, 0)
	body := doc.Find("body")
	body.Find("div.zm-editable-content").Each(func(i int, s *goquery.Selection) {
		s.Find("img").Each(func(i int, s *goquery.Selection) {
			result, exists := s.Attr("data-original")
			if exists {
				if result != temp {
					temp = result
					getImg("http:" + temp, num)
					writeImgUrl("http:" + temp + "\n")
				}
			}
		})
	})
}

func getImg(url string, dirName string) {
	res, _ := http.Get(url)
	slice := strings.Split(url, "/")
	length := len(slice)
	storagePath := "./pictures/" + dirName + "/" + slice[length-1]
	file, err := os.OpenFile(fileName, os.O_APPEND, 0)
	if err != nil {
		file, err = os.Create(storagePath)
		if err != nil {
			panic(err)
		}
		io.Copy(file, res.Body)
	}
}

func writeImgUrl(url string) {
	file, err := os.OpenFile(fileName, os.O_APPEND, 0)
	if err != nil {
		monkey, _ := os.Create(fileName)
		monkey.WriteString(url)
	} else {
		file.WriteString(url)
	}
	defer file.Close()
}

func main() {
	os.Remove(fileName)
	getUrls("http://www.zhihu.com/collection/71963247")
	getUrls("http://www.zhihu.com/collection/71964476")
	getUrls("http://www.zhihu.com/collection/71977517")
	getUrls("http://www.zhihu.com/collection/71964508")
	getUrls("http://www.zhihu.com/collection/71578326")
}
