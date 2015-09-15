package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"runtime"
)

const isFileNeed = false
const fileName = "url.txt"
const dirName = "./pictures/"
const zhihu = "http://www.zhihu.com"
const agreeNum = 30

var quit chan int
var allQuestions []string

func getUrls(url string) {
	doc, _ := goquery.NewDocument(url)
	fmt.Println(doc.Find("title").Text())
	fmt.Println(url)
	body := doc.Find("body")
	body.Find("div.zm-item").Each(func(i int, s *goquery.Selection) {
		h2 := s.Find("h2")
		question_url, _ := h2.Find("a").Attr("href")
		if strings.HasPrefix(question_url, "/question/") {
			question_url = zhihu + question_url
			if !isUrlInSlice(question_url) {
				getImgUrl(question_url)
				allQuestions = append(allQuestions, question_url)
			}
		}
	})
	quit <- 0
}

func isUrlInSlice(url string) bool {
	for _, s := range allQuestions {
		if s == url {
			return true
		}
	}
	return false
}

func isNotMale(url string) bool {
	url = zhihu + url
	doc, _ := goquery.NewDocument(url)
	body := doc.Find("body")
	div := body.Find("div.zm-profile-header-info")
	span := div.Find("span.item")
	gender, exists := span.Find("i").Attr("class")
	if exists {
		if strings.Contains(gender, "icon-profile-male") {
			return false
		} else {
			return true
		}
	}
	return false
}

//根据问题获取图片url
func getImgUrl(url string) {
	var temp string
	doc, _ := goquery.NewDocument(url)
	dir := doc.Find("title").Text()
	name := strings.Split(dir, "-")[0]
	slice := strings.Split(url, "/")
	length := len(slice)
	id := slice[length-1]
	if isFileNeed {
		writeImgUrl("<br><br><a href=" + url + ">" + name + "</a><br>" + "\n")
	}
	os.Mkdir(dirName+id, 0)
	body := doc.Find("body")
	body.Find("div.zm-item-answer").Each(func(i int, answer *goquery.Selection) {
		answerHead := answer.Find("div.answer-head")

		//获取点赞数
		numString, exists := answerHead.Find("div.zm-item-vote-info").Attr("data-votecount")
		num := 0
		if exists {
			num, _ = strconv.Atoi(numString)
		}
		var numOK bool
		numOK = false
		if num >= agreeNum {
			numOK = true
		}
		//获取性别， 收集匿名用户，女性用户和未填写性别用户的图
		var userOK bool
		userOK = false
		h3 := answerHead.Find("h3")
		if strings.Contains(h3.Text(), "匿名用户") || strings.Contains(h3.Text(), "知乎用户") {
			userOK = true
		} else {
			userUrl, exists := h3.Find("a").Attr("href")
			if exists && strings.Contains(userUrl, "people") {
				userOK = isNotMale(userUrl)
			}
		}

		if numOK && userOK {
			answer.Find("div.zm-item-rich-text").Each(func(i int, s *goquery.Selection) {
				s.Find("img").Each(func(i int, s *goquery.Selection) {
					result, exists := s.Attr("data-original")
					if exists {
						if result != temp {
							temp = result
							getImg("http:"+temp, id)
							imgUrl := "http:" + temp;
							if isFileNeed {
								writeImgUrl("<a href=" + imgUrl + ">" + imgUrl + "</a><br>" + "\n")
							}
						}
					}
				})
			})
		}
	})
}

//获取图片保存到本地
func getImg(url string, dir string) {
	res, _ := http.Get(url)
	slice := strings.Split(url, "/")
	length := len(slice)
	storagePath := dirName + dir + "/" + slice[length-1]
	file, err := os.Create(storagePath)
	if err != nil {
		panic(err)
	}
	io.Copy(file, res.Body)
}

//将图片url写入文件
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

// 检查文件或目录是否存在
// 如果由 filename 指定的文件或目录存在则返回 true，否则返回 false
func existFile(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func main() {
	if existFile(fileName) {
		os.Remove(fileName)
	}
	if !existFile(dirName) {
		os.Mkdir(dirName, 0)
	}

	//	getImgUrl("http://www.zhihu.com/question/20095161")

	NCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(NCPU)
	quit = make(chan int, NCPU)

	urls := []string{
		"http://www.zhihu.com/collection/45762052",
		"http://www.zhihu.com/collection/71963247",
		"http://www.zhihu.com/collection/71964476",
		"http://www.zhihu.com/collection/71977517",
		"http://www.zhihu.com/collection/71964508",
		"http://www.zhihu.com/collection/71578326",
		"http://www.zhihu.com/collection/60771406",
		"http://www.zhihu.com/collection/30256531",
		"http://www.zhihu.com/collection/30822111",
		"http://www.zhihu.com/collection/38624707",
		"http://www.zhihu.com/collection/26347524",
		"http://www.zhihu.com/collection/26348030",
		"http://www.zhihu.com/collection/26815754",
		"http://www.zhihu.com/collection/53719722",
		"http://www.zhihu.com/collection/36731404",
	}
	if isFileNeed {
		writeImgUrl("<html><head><meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\"/><title>www.zhihu.com</title></head>")
	}
	for _, url := range urls {
		go getUrls(url)
	}

	for _, _ = range urls {
		<-quit
	}
	if isFileNeed {
		writeImgUrl("</body></html>")
	}
}
