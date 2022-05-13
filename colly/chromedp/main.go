package main

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func text(res *string) chromedp.Tasks {
	return chromedp.Tasks{
		// 访问页面
		chromedp.Navigate(`https://movie.douban.com/tag/`),
		// 等待列表渲染
		chromedp.Sleep(5 * time.Second),
		// 获取获取服务列表HTML
		chromedp.OuterHTML("#content", res, chromedp.ByID),
	}
}

func main() {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`http://www.nhc.gov.cn/wjw/mtbd/202203/5a5b1ab9d99d45e1b476156dc2912fa6.shtml`),

		chromedp.Reload(),

		chromedp.Sleep(5 * time.Second),

		chromedp.Text(`#article_content`, &res, chromedp.NodeVisible),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(strings.TrimSpace(res))
}