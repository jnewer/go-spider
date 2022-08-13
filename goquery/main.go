package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://gorm.io/zh_CN/docs/"
	//dom, err := goquery.NewDocument(url)

	//client := &http.Client{}
	//req, _ := http.NewRequest("GET", url, nil)
	//resp, err := client.Do(req)
	//dom, err := goquery.NewDocumentFromResponse(resp)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//dom.Find(".sidebar-link").Each(func(i int, s *goquery.Selection) {
	//	href, _ := s.Attr("href")
	//	text := s.Text()
	//	fmt.Println(i, href, text)
	//})

	//html := `<body>
	//            <div id="div1">DIV1</div>
	//            <div class="name">DIV2</div>
	//            <span>SPAN</span>
	//        </body>`
	//
	//dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	//if err != nil {
	//	log.Fatal(err)
	//}
	////元素名称选择器
	//dom.Find("div").Each(func(i int, s *goquery.Selection) {
	//	fmt.Println("i", i, "select text", s.Text())
	//})
	//
	////ID选择器
	//dom.Find("#div1").Each(func(i int, s *goquery.Selection) {
	//	fmt.Println(s.Text()) //DIV1
	//})
	//
	////class类选择器
	//dom.Find(".name").Each(func(i int, s *goquery.Selection) {
	//	fmt.Println(s.Text()) //DIV2
	//})

	d, _ := goquery.NewDocument(url)
	d.Find(".sidebar-link").Each(func(i int, s *goquery.Selection) {

		href, _ := s.Attr("href")
		fmt.Printf("href: %v\n", href)

		detailUrl := url + href
		fmt.Printf("detailUrl: %v\n", detailUrl)
		d, _ = goquery.NewDocument(detailUrl)
		//
		title := d.Find(".article-title").Text()
		fmt.Printf("title: %v\n", title)
		
		content, _ := d.Find(".article").Html()
		fmt.Printf("content: %v\n", content)
	})
}
