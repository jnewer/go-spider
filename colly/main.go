package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()
	//c.OnHTML(".sidebar-link", func(e *colly.HTMLElement) {
	//	e.Request.Visit(e.Attr("href"))
	//})
	//
	//c.OnRequest(func(r *colly.Request) {
	//	fmt.Println("url:", r.URL)
	//})
	//
	//c.OnRequest(func(r *colly.Request) {
	//	fmt.Println("before request:OnRequest")
	//})
	//
	//c.OnError(func(r *colly.Response, err error) {
	//	fmt.Println("on error:OnError")
	//})
	//
	//c.OnResponse(func(r *colly.Response) {
	//	fmt.Println("after response:OnResponse")
	//})
	//
	//c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	//	fmt.Println("after OnResponse receive html:OnHTML")
	//})
	//
	//c.OnXML("//h1", func(e *colly.XMLElement) {
	//	fmt.Println("after OnResponse receive xml:OnXML")
	//})
	//
	//c.OnScraped(func(r *colly.Response) {
	//	fmt.Println("end", r.Request.URL)
	//})

	c.OnHTML(".sidebar-link", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		if href != "index.html" {
			c.Visit(e.Request.AbsoluteURL(href))
		}
	})

	c.OnHTML(".article-title", func(e *colly.HTMLElement) {
		title := e.Text
		fmt.Printf("title: %v\n", title)
	})

	c.OnHTML(".article", func(e *colly.HTMLElement) {
		content, _ := e.DOM.Html()
		fmt.Printf("content: %v\n", content)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL.String())
	})

	c.Visit("https://gorm.io/zh_CN/docs/")
}
