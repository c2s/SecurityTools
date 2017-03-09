package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
	"sync"
	"strings"
	"github.com/levigross/grequests"
	"time"
	"strconv"
)

var wg sync.WaitGroup

func main() {

	var keyword string
	var page  int
	    page = 10
	    arg_num := len(os.Args)
	if arg_num >= 3 {
		keyword = os.Args[1]
		//page    = os.Args[2]
		//page,err:=strconv.Atoi(page)
	}

	now := time.Now()
	DownloadDir := "urls/"
	o := 1
	//SiteUrl := "https://www.baidu.com/s?wd="+keyword+"&pn=0&rn=50"
	for o <= page {
		SpiderUrl := "https://www.baidu.com/s?wd=inurl:"+keyword+"&rn=50"
		if o > 1 {
			SpiderUrl = SpiderUrl + "&pn=" + (strconv.Itoa(o*50))
		}

		doc, err := goquery.NewDocument(SpiderUrl)
		if err != nil {
			fmt.Errorf("网页打开错误:%#v", err)
			os.Exit(-1)
		}
		limit := 0

		doc.Find(".container_l .result .f13").Each(func(i int, s *goquery.Selection) {
			url := s.Find("a").Text()
			// 拿到url

			fmt.Println(url)
			exists := false
			if (exists) {
				os.Exit(-1)
			}

			if (exists) {
				wg.Add(1)
				go func(url string) {
					defer wg.Done()
					//下载图片

					for {
						n := 0
						//持续下载
						res, _ := grequests.Get(url, &grequests.RequestOptions{
							Headers:map[string]string{
								"Referer":"http://"+url+"/",
								"User-Agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.95 Safari/537.36"}})

						if res.StatusCode != 200 {
							break
						}

						length := res.Header.Get("Content-Length")
						slen, _ := strconv.Atoi(length)
						if slen < 4100 {
							break
						}

						var dirname string
						dirname = strings.Join([]string{dirname, DownloadDir}, "")
						if _, err := os.Stat(dirname); err != nil {
							//fmt.Printf("创建下载文件夹:%s\n", dirname)
							os.MkdirAll(dirname, 0755)
						}
						filename := strings.Join([]string{dirname, "/", "01.jpg"}, "")
						res.DownloadToFile(filename)
						n++
						if limit >=14 {
							break
						}
					}
				}(url)
			}
		})




		//for limit <= 14 {
		//	doc.Find(".container_l").Each(func(i int, s *goquery.Selection) {
		//		url, exists := s.Find(".cr-offset").Attr("id")
		//		// 拿到url
		//
		//		fmt.Println(exists)
		//		if (exists) {
		//			os.Exit(-1)
		//		}
		//
		//		if (exists) {
		//			wg.Add(1)
		//			go func(url string) {
		//				defer wg.Done()
		//				//下载图片
		//
		//				for {
		//					n := 0
		//					//持续下载
		//					res, _ := grequests.Get(url, &grequests.RequestOptions{
		//						Headers:map[string]string{
		//							"Referer":"http://"+url+"/",
		//							"User-Agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.95 Safari/537.36"}})
		//
		//					if res.StatusCode != 200 {
		//						break
		//					}
		//
		//					length := res.Header.Get("Content-Length")
		//					slen, _ := strconv.Atoi(length)
		//					if slen < 4100 {
		//						break
		//					}
		//
		//					var dirname string
		//					dirname = strings.Join([]string{dirname, DownloadDir}, "")
		//					if _, err := os.Stat(dirname); err != nil {
		//						//fmt.Printf("创建下载文件夹:%s\n", dirname)
		//						os.MkdirAll(dirname, 0755)
		//					}
		//					filename := strings.Join([]string{dirname, "/", "01.jpg"}, "")
		//					res.DownloadToFile(filename)
		//					n++
		//					if limit >=14 {
		//						break
		//					}
		//				}
		//			}(url)
		//		}
		//	})
		//	limit++
		//}
		o++
	}

	wg.Wait()

	fmt.Printf("下载任务完成，耗时:%#v\n", time.Now().Sub(now))
}
