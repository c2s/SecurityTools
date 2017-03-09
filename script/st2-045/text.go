package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
	"sync"
	"strconv"
)

var wg sync.WaitGroup

func main() {

	var keyword string
	arg_num := len(os.Args)
	if arg_num >= 3 {
		keyword = os.Args[1]
		//page    = os.Args[2]
		//page,err:=strconv.Atoi(page)
	}

	o := 1
	//SiteUrl := "https://www.baidu.com/s?wd="+keyword+"&pn=0&rn=50"


	SpiderUrl := "https://www.baidu.com/s?wd=inurl:"+keyword+"&rn=50"
	if o == 1 {
		SpiderUrl = SpiderUrl + "&pn=" + (strconv.Itoa(o*50))
	}

	doc, err := goquery.NewDocument(SpiderUrl)
	if err != nil {
		fmt.Errorf("网页打开错误:%#v", err)
		os.Exit(-1)
	}

	doc.Find(".wrapper_l .s_tab").Each(func(i int, s *goquery.Selection) {
		url := doc.Find(".wrapper_l .s_tab b").Text()

		fmt.Printf("Review %d: %s\n", i, url)
		// 拿到url
		exists := false
		if (exists) {
			os.Exit(-1)
		}
	})

	wg.Wait()
}
