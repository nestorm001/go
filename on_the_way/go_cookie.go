package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"fmt"
)

var gCurCookies []*http.Cookie
var gCurCookieJar *cookiejar.Jar

//do init before all others
func initAll() {
	gCurCookies = nil
	//var err error;
	gCurCookieJar, _ = cookiejar.New(nil)
}

//get url response html
func getUrlRespHtml(url string) string {
	log.Printf("getUrlRespHtml, url=%s", url)

	var respHtml string = ""

	httpClient := &http.Client{
		CheckRedirect: nil,
		Jar:           gCurCookieJar,
	}

	// httpResp, err := httpClient.Get("http://example.com")

	httpReq, err := http.NewRequest("GET", url, nil)
	//httpReq.Header.Add("If-None-Match", `W/"wyzzy"`)
	httpResp, err := httpClient.Do(httpReq)

	//httpResp, err := http.Get(url)
	//log.Printf("http.Get done")
	if err != nil {
		log.Printf("http get url=%s response error=%s\n", url, err.Error())
	}
	log.Printf("httpResp.Header=%s", httpResp.Header)
	log.Printf("httpResp.Status=%s", httpResp.Status)

	defer httpResp.Body.Close()

	body, errReadAll := ioutil.ReadAll(httpResp.Body)
	if errReadAll != nil {
		log.Printf("get response for url=%s got error=%s\n", url, errReadAll.Error())
	}

	gCurCookies = gCurCookieJar.Cookies(httpReq.URL)
	fmt.Println(gCurCookies)
	respHtml = string(body)
	return respHtml
}

func dbgPrintCurCookies() {
	var cookieNum int = len(gCurCookies)
	log.Printf("cookieNum=%d", cookieNum)
	for i := 0; i < cookieNum; i++ {
		var curCk *http.Cookie = gCurCookies[i]
		//log.Printf("curCk.Raw=%s", curCk.Raw)
		log.Printf("Cookie [%d]", i)
		log.Printf("Name\t=%s", curCk.Name)
		log.Printf("Value\t=%s", curCk.Value)
		log.Printf("Path\t=%s", curCk.Path)
		log.Printf("Domain\t=%s", curCk.Domain)
		log.Printf("Expires\t=%s", curCk.Expires)
		log.Printf("RawExpires=%s", curCk.RawExpires)
		log.Printf("MaxAge\t=%d", curCk.MaxAge)
		log.Printf("Secure\t=%t", curCk.Secure)
		log.Printf("HttpOnly=%t", curCk.HttpOnly)
		log.Printf("Raw\t=%s", curCk.Raw)
		log.Printf("Unparsed=%s", curCk.Unparsed)
	}
}

func main() {
	initAll()

	//step1: access baidu url to get cookie BAIDUID
	log.Printf("======BAIDUID Cookie ======")
	var baiduMainUrl string = "http://www.baidu.com/"
	log.Printf("baiduMainUrl=%s", baiduMainUrl)
	respHtml := getUrlRespHtml(baiduMainUrl)
	log.Printf("respHtml=%s", respHtml)
	dbgPrintCurCookies()

	//check cookie

	//step2: login, pass paras, extract resp cookie
	log.Printf("======login_token ======")
	//https://passport.baidu.com/v2/api/?getapi&class=login&tpl=mn&tangram=true
	var getapiUrl string = "https://passport.baidu.com/v2/api/?getapi&class=login&tpl=mn&tangram=true"
	var getApiRespHtml string = getUrlRespHtml(getapiUrl)
	log.Printf("getApiRespHtml=%s", getApiRespHtml)
	dbgPrintCurCookies()
}