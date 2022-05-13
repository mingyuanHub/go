package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/mozillazg/request"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	startUrl, err := findStartUrl("圣墟", "辰东")
	if err != nil {
		log.Fatal(err)
	}
	println(startUrl)

	crawl(startUrl, "test.txt")
}

// findStartUrl 根据小说名和作者找到 colly 的 startUrl
func findStartUrl(name string, author string) (string, error) {

	var targetUrl string
	var Err error

	// request请求库常规用法，用来发起请求，获取响应
	c := new(http.Client)
	req := request.NewRequest(c)
	req.Headers = map[string]string{
		"Accept-Encoding": "gzip,deflate,sdch",
		"Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
	}
	req.Data = map[string]string{
		"searchkey": name,
	}
	url := "https://www.xbiquge.la/modules/article/waps.php"
	resp, _ := req.Post(url)
	defer resp.Body.Close()  // Don't forget close the response body

	html, err := resp.Text()
	if err != nil {
		Err = err
		log.Fatal(err)
	}

	// goquery 常规用法
	dom, _ := goquery.NewDocumentFromReader(strings.NewReader(html))
	dom.Find("tr").EachWithBreak(func(i int, s *goquery.Selection) bool {
		_title := s.Find("a").Eq(0).Text()
		_author := s.Find("td").Eq(2).Text()
		// 找到匹配的链接
		if _title == name && _author == author {
			url, exists := s.Find("a").Eq(0).Attr("href")
			if !exists {
				Err = nil
				log.Fatal(exists)
			}
			targetUrl = url
		}
		return true
	})

	// 根据小说的首页，获取第一个章节的页面的链接
	resp2, err2 := req.Get(targetUrl)
	if err2 != nil {
		log.Fatalln("get start url request fail...")
	}
	text, _ := resp2.Text()
	dom2, _ := goquery.NewDocumentFromReader(strings.NewReader(text))
	startUrl, ok:= dom2.Find("#list").Find("a").Eq(0).Attr("href")
	if !ok {
		log.Fatalln("get start url fail...")
	}

	return "https://www.xbiquge.la" + startUrl, Err
}

// 爬取小说内容并写入text文件
func crawl(startUrl string, fileName string) {
	c := colly.NewCollector(
		// allowed domain
		colly.AllowedDomains("www.xbiquge.la", "www.coursera.org"),
		//colly.AllowURLRevisit(),
	)

	// 使用随机user-agent
	extensions.RandomUserAgent(c)

	// HTTP 的配置
	c.WithTransport(&http.Transport{
		Proxy: http.ProxyFromEnvironment, // 使用代理
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second, // 超时时间
			KeepAlive: 30 * time.Second, // keepAlive 超时时间
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,              // 最大空闲连接数
		IdleConnTimeout:       90 * time.Second, // 空闲连接超时
		TLSHandshakeTimeout:   10 * time.Second, // TLS 握手超时
		ExpectContinueTimeout: 1 * time.Second,
	})

	// 设置请求信息
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Host", "www.xbiquge.la")
		r.Headers.Set("X-Requested-With", "XMLHttpRequest")
		r.Headers.Set("Referer", "https://www.xbiquge.la/13/13959/24775996.html")
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3314.0 Safari/537.36 SE 2.X MetaSr 1.0")

		log.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		log.Println("response received", r.StatusCode)
	})

	// 解析章节内容 并 写入文本文件
	c.OnHTML("div[id='content']", func(e *colly.HTMLElement) {
		// title
		title := e.DOM.ParentsUntil("body").Find(".bookname h1").Text() + "\n\n"

		println(title)

		// content
		e.DOM.Find("p").Remove()
		content := e.DOM.Text()
		str := title + content

		writeToText(str, fileName)
	})

	// 翻页
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.DOM.EachWithBreak(func(i int, s *goquery.Selection) bool {
			if "下一章" == s.Text() {
				baseUrl, _ := s.Attr("href")
				url := "https://www.xbiquge.la" + baseUrl

				c.Visit(url)
				return false
			}
			return true
		})
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("error:", r.StatusCode, err)
	})

	c.Visit(startUrl)

}

// 将字符串，写入文本文件
func writeToText(str string, fileName string) {
	wd, _ := os.Getwd()
	filePath := wd + "\\file1\\" + fileName  // 存放小说的TXT文件路径
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)  // 以追加方式写入
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	_, err = fmt.Fprintln(f, str)
	if err != nil {
		log.Fatal(err)
	}
}