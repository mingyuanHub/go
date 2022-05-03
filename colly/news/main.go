package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/mozillazg/request"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	var spokesman string

	spokesman = "华春莹"

	startDay := "2017-01-01"
	endDay := "2022-05-01"

	Waijiaobu(spokesman, startDay, endDay)
}


type WaijiaobuResponse struct {
	Success bool
	Msg     string
	Data    *WaijiaobuResponseData
}

type WaijiaobuResponseData struct {
	Middle *WaijiaobuResponseDataMiddle
	Pager *WaijiaobuResponseDataPager
}

type WaijiaobuResponseDataMiddle struct {
	List []*WaijiaobuResponseDataMiddleListItem
}

type WaijiaobuResponseDataMiddleListItem struct {
	Time string
	Url  string
}

type WaijiaobuResponseDataPager struct {
	PageNo    int
	PageSize  int
	PageCount int
	Total     int
}

func Waijiaobu(spokesman, startDay, endDay string) {
	url := "https://www.mfa.gov.cn/irs/front/search"
	post := map[string]interface{}{
		"beginDateTime": Int642String(DataToTimeMillSecond(startDay)),
		"endDateTime": Int642String(DataToTimeMillSecond(endDay)),
		"appendixType":"",
		"code":"17e50b77dab",
		"codes":"",
		"configCode":"",
		"dataTypeId":"18",
		"filters":[]string{},
		"granularity": "CUSTOM",
		"historySearchWords": []string{spokesman},
		"isSearchForced": 0,
		"orderBy": "related",
		"pageNo": 1,
		"pageSize": 10,
		"searchBy": "all",
		"searchWord": spokesman,
	}

	c := new(http.Client)
	req := request.NewRequest(c)
	req.Headers = map[string]string{
		"Content-Type":"application/json",
		"Accept-Encoding": "gzip,deflate,sdch",
		"Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
		"User-Agent":"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36",
	}

	fmt.Println(url, Map2Json(post))
	resp, _ := req.PostForm(url, Map2Json(post))
	defer resp.Body.Close()  // Don't forget close the response body

	data, err := ioutil.ReadAll(resp.Body)
	io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var response = &WaijiaobuResponse{}

	if resp.Status == "200 OK" {
		err = json.Unmarshal(data, &response)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	fmt.Println(response)

	if response.Success == true {
		for _, item := range response.Data.Middle.List {
			fmt.Println(item.Url)
			WaijiaobuColly(item.Url, spokesman)
		}
	}
}


//外交部
func WaijiaobuColly(url, spokesman string) {

	c := colly.NewCollector(
		colly.AllowedDomains("www.mfa.gov.cn"),
	)

	// 使用随机user-agent
	extensions.RandomUserAgent(c)

	// HTTP 的配置
	c.WithTransport(&http.Transport{
		Proxy: http.ProxyFromEnvironment, // 使用代理
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second, // 超时时间
			KeepAlive: 30 * time.Second, // keepAlive 超时时间
		}).DialContext,
		MaxIdleConns:          100,              // 最大空闲连接数
		IdleConnTimeout:       90 * time.Second, // 空闲连接超时
		TLSHandshakeTimeout:   10 * time.Second, // TLS 握手超时
		ExpectContinueTimeout: 1 * time.Second,
	})

	// 设置请求信息
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Host", "www.mfa.gov.cn")
		r.Headers.Set("X-Requested-With", "XMLHttpRequest")
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3314.0 Safari/537.36 SE 2.X MetaSr 1.0")

		log.Println("Visiting", r.URL)
	})

	// 请求响应
	c.OnResponse(func(r *colly.Response) {
		fmt.Println(string(r.Body))
		log.Println("response received", r.StatusCode)
	})

	c.OnHTML(".is-news", func(e *colly.HTMLElement) {
		times := e.DOM.Find(".sourceTime").Text()
		fmt.Println(times)
	})

	c.Visit(url)
}

func DataToTimeSecond(day string) int64 {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	tt, _ := time.ParseInLocation("2006-01-02", day, loc)
	return  tt.Unix()
}

func DataToTimeMillSecond(day string) int64 {
	return DataToTimeSecond(day) * 1000
}

func Int642String(i int64) string {
	return fmt.Sprintf("%d", i)
}

func Map2Json(m interface{}) string {
	mjson,_ :=json.Marshal(m)
	return  string(mjson)
}

