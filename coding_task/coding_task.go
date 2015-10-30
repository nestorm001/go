package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"time"
	"strings"
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

var jar = NewJar()
var client = http.Client{Jar: jar}

func main() {
	task()
	commit()
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
	fmt.Println(jar.cookies)
	resp, _ := client.PostForm(task_url, url.Values{
		"content":      {task_title},
		"status":       {"1"},
		"user_name":    {userName},
		"project_name": {projectName},
		"owner_id":     {ownerId},
	})
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

	resp, _ = client.PostForm(file_url, url.Values{
		"content":  {content},
		"message":  {commit_title},
		"lastCommitSha": {s},
	})
	b, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println(string(b))

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