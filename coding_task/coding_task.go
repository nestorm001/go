package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
)

const title = "猴子真坑"
const userName = "nesto"
const password = "e0efb1d23aa3e98057630a7fb44aa1e759a24294"
const projectName = "monkey"
const ownerId = "99239"
const coding = "https://coding.net/api"
const task_url = "/user/" + userName + "/project/" + projectName + "/task"
const login_url = "/account/login"
const captcha_url = "/account/captcha/login"

func main() {
	//captcha
	jar := NewJar()
	client := http.Client{Jar: jar}

	req, _ := http.NewRequest("GET", coding+captcha_url, nil)
	resp, _ := client.Do(req)
	b, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println(string(b))

	//login
	resp, _ = client.PostForm(coding+login_url, url.Values{
		"email":      {userName},
		"password":       {password},
	})
	b, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println(string(b))

	cookies := resp.Cookies()
	for _, cookie := range cookies {
		fmt.Println("cookie:", cookie.Value)
	}

	//task
	resp, _ = client.PostForm(coding+task_url, url.Values{
		"content":      {title},
		"status":       {"1"},
		"user_name":    {userName},
		"project_name": {projectName},
		"owner_id": {ownerId},
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
