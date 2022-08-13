package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
	"xorm.io/xorm"
)

func fetch(url string) string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36")
	req.Header.Add("Cookie", "JSESSIONID=_DvL7c5D0RtcON9mGKvPKzhf9yo_XYUpddb2wtIx; Hm_lvt_0882a29bba355751ac6f4fe522f8d6de=1646879724,1646900491,1646973412,1646979544; Hm_lpvt_0882a29bba355751ac6f4fe522f8d6de=1646981342")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Http get err:", err)
		return ""
	}

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Http status code: ", resp.StatusCode)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read error", err)
		return ""
	}

	return string(body)
}

func parse(html string) {
	html = strings.Replace(html, "\n", "", -1)
	reSidebar := regexp.MustCompile(`<aside id="sidebar" role="navigation">(.*?)</aside>`)
	sidebar := reSidebar.FindString(html)

	reLink := regexp.MustCompile(`href="(.*?)"`)
	links := reLink.FindAllString(sidebar, -1)

	baseUrl := "https://gorm.io/zh_CN/docs/"
	for _, v := range links {
		s := v[6 : len(v)-1]
		url := baseUrl + s
		fmt.Printf("url: %v\n", url)

		body := fetch(url)
		go parse2(body)
	}
}

func parse2(body string) {
	body = strings.Replace(body, "\n", "", -1)

	reContent := regexp.MustCompile(`<div class="article">(.*?)<div>`)
	content := reContent.FindString(body)

	reTitle := regexp.MustCompile(`<h1 class="article-title" itemprop="name">(.*?)</h1>`)

	title := reTitle.FindString(content)
	fmt.Printf("title: %v\n", title)
	title = title[42 : len(title)-5]
	fmt.Printf("title: %v\n", title)

	//save(title, content)
	saveToDB(title, content)
}

func save(title string, content string) {
	err := os.WriteFile("./http/pages/"+title+".html", []byte(content), 0644)
	if err != nil {
		panic(err)
	}
}

var engine *xorm.Engine
var err error

func init() {
	engine, err = xorm.NewEngine("mysql", "root:@/go_db?charset=utf8mb4")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		err2 := engine.Ping()
		if err2 != nil {
			fmt.Printf("err2 :v\n", err2)
		} else {
			print("连接成功")
		}
	}
}

type XormPage struct {
	Id      int64
	Title   string
	Content string    `xorm:"text"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated""`
}

func saveToDB(title string, content string) {
	engine.Sync(new(XormPage))

	page := XormPage{
		Title:   title,
		Content: content,
	}

	affected, err := engine.Insert(&page)

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	fmt.Println("save:" + string(affected))
}
func main() {
	//for i := 0; i < 2; i++ {
	//	url := "https://zzk.cnblogs.com/s/blogpost?Keywords=golang&pageindex=%d"
	//	url = fmt.Sprintf(url, i)
	//	fmt.Println(url)
	//	s := fetch(url)
	//	fmt.Printf("s: %v\n", s)
	//}

	url := "https://gorm.io/zh_CN/docs/"
	s := fetch(url)
	//fmt.Printf("s: %v\n", s)
	parse(s)

}
