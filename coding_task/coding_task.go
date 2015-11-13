package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

const userName = "nesto"
const password = "e0efb1d23aa3e98057630a7fb44aa1e759a24294"
const projectName = "not-auto-monkey"
const ownerId = "99239"

const task_title = "猴子真坑"
const commit_title = "自动提交"
const coding = "https://coding.net/api"
const task_url = coding + "/user/" + userName + "/project/" + projectName + "/task"
const login_url = coding + "/account/login"
const captcha_url = coding + "/account/captcha/login"
const file_url = coding + "/user/" + userName + "/project/" + projectName + "/git/edit/master%252FREADME.md"
const merge_url = coding + "/user/" + userName + "/project/" + projectName + "/git/merge"
const ide_url = "https://ide.coding.net/ws/yriinn"

var jar = NewJar()
var client = http.Client{Jar: jar}

var iid int

func main() {
	for !netTest() {
	}
	mainProcess()
	time.Sleep(3 * time.Second)
}

func mainProcess() {
	if !isPushedToday() {
		task()
		commit()
		merge()
		cancelMerge()
	} else {
		fmt.Println("今天已提交过")
	}
	ide()
}

func netTest() bool {
	req, _ := http.NewRequest("GET", captcha_url, nil)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("出现了问题，稍后重试")
		for i := 0; i < 11; i++ {
			fmt.Println(10 - i)
			time.Sleep(time.Second)
		}
		return false
	} else {
		resp.Body.Close()
		return true
	}
}

func ide() {
	login()
	req, _ := http.NewRequest("GET", ide_url, nil)
	resp, _ := client.Do(req)
	b, _ := ioutil.ReadAll(resp.Body)
	js, _ := simplejson.NewJson(b)
	fmt.Println(js)
	resp.Body.Close()
}

func login() {
	//captcha
	req, _ := http.NewRequest("GET", captcha_url, nil)
	resp, _ := client.Do(req)
	resp.Body.Close()

	//login
	resp, _ = client.PostForm(login_url, url.Values{
		"email":       {userName},
		"password":    {password},
		"remember_me": {"true"},
	})
	resp.Body.Close()

}

func task() {
	login()
	//task
	//	fmt.Println(jar.cookies)
	resp, _ := client.PostForm(task_url, url.Values{
		"content":      {task_title},
		"status":       {"1"},
		"user_name":    {userName},
		"project_name": {projectName},
		"owner_id":     {ownerId},
	})
	b, _ := ioutil.ReadAll(resp.Body)
	js, _ := simplejson.NewJson(b)
	id, _ := js.Get("data").Get("id").Int()
	fmt.Println(js)
	resp.Body.Close()

	login()
	delete_url := task_url + "/" + strconv.Itoa(id)
	req, _ := http.NewRequest("DELETE", delete_url, nil)
	resp, _ = client.Do(req)
	resp.Body.Close()
}

func commit() {
	login()

	//commit and push
	fmt.Println(file_url)
	req, _ := http.NewRequest("GET", file_url, nil)
	resp, _ := client.Do(req)
	b, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	js, _ := simplejson.NewJson(b)
	s, _ := js.Get("data").Get("lastCommit").String()
	content, _ := js.Get("data").Get("file").Get("data").String()
	fmt.Println("commitId: " + s)

	login()
	today := time.Now().Format("2006-01-02")
	if !strings.Contains(content, today) {
		content = content + "\n# " + today
	}
	//	fmt.Println("content: " + content)

	resp, err := client.PostForm(file_url, url.Values{
		"content":       {content},
		"message":       {commit_title},
		"lastCommitSha": {s},
	})
	if err != nil {
		panic(err)
		mainProcess()
	} else {
		b, _ = ioutil.ReadAll(resp.Body)
		result, _ := simplejson.NewJson(b)
		resultMessage, _ := result.Get("message").Int()
		if resultMessage == 0 {
			fmt.Println("应该成功了")
		}
		resp.Body.Close()
	}
}

func merge() {
	login()
	fmt.Println(merge_url)
	resp, err := client.PostForm(merge_url, url.Values{
		"srcBranch": {"merge"},
		"desBranch": {"master"},
		"title":     {"猴子你妈逼你结婚了吗"},
		"author":    {userName},
		"content":   {"猴子你妈逼你结婚了吗"},
	})

	if err != nil {
		merge()
	}
	b, _ := ioutil.ReadAll(resp.Body)
	result, _ := simplejson.NewJson(b)
	iid, _ = result.Get("data").Get("merge_request").Get("iid").Int()
	resp.Body.Close()
}

func cancelMerge() {
	login()
	cancel_url := merge_url + "/" + strconv.Itoa(iid) + "/cancel"
	resp, err := client.PostForm(cancel_url, url.Values{})
	if err != nil {
		cancelMerge()
	}
	b, _ := ioutil.ReadAll(resp.Body)
	result, _ := simplejson.NewJson(b)
	fmt.Println(result)
	resp.Body.Close()
}

func isPushedToday() bool {
	login()

	//commit and push
	req, _ := http.NewRequest("GET", file_url, nil)
	resp, _ := client.Do(req)
	b, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	js, _ := simplejson.NewJson(b)
	content, _ := js.Get("data").Get("file").Get("data").String()
	today := time.Now().Format("2006-01-02")
	return strings.Contains(content, today)
}

type Jar struct {
	lk      sync.Mutex
	cookies map[string][]*http.Cookie
}

func NewJar() *Jar {
	jar := new(Jar)
	jar.cookies = make(map[string][]*http.Cookie)
	return jar
}

// SetCookies handles the receipt of the cookies in a reply for the
// given URL.  It may or may not choose to save the cookies, depending
// on the jar's policy and implementation.
func (jar *Jar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	jar.lk.Lock()
	jar.cookies[u.Host] = cookies
	jar.lk.Unlock()
}

// Cookies returns the cookies to send in a request for the given URL.
// It is up to the implementation to honor the standard cookie use
// restrictions such as in RFC 6265.
func (jar *Jar) Cookies(u *url.URL) []*http.Cookie {
	return jar.cookies[u.Host]
}
