package fetcher

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"math/rand"
	"net/http"
	"time"
)

/*
参数：string类型的url
功能：根据url爬取页面
返回值：返回爬取结果OR错误信息
*/

func Fetcher(targetUrl string) ([]byte, error) {
	request := gorequest.New().Get(targetUrl)
	setHeader(request)
	response, body, errs := request.End()
	if errs != nil {
		//log.Print(errs)
		return nil, errs[0]
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status_error: %d", response.StatusCode)
	}
	//fmt.Println(response)
	//fmt.Println()
	//fmt.Println(body)
	return []byte(body), nil
}

func setHeader(r *gorequest.SuperAgent) {
	r.Set("x-devtools-emulate-network-conditions-client-id","5f2fc4da-c727-43c0-aad4-37fce8e3ff39")
	r.Set("upgrade-insecure-requests", "1")
	r.Set("user-agent", getAgent())
	//r.Set("ccept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	r.Set("dnt", "1")
	//r.Set("accept-encoding", "gzip, deflate")
	//r.Set("accept-language", "zh-CN,zh;q=0.8,en;q=0.6")
	r.Set("cookie", "__c=1501326829; lastCity=101020100; __g=-; __l=r=https%3A%2F%2Fwww.google.com.hk%2F&l=%2F; __a=38940428.1501326829..1501326829.20.1.20.20; Hm_lvt_194df3105ad7148dcf2b98a91b5e727a=1501326839; Hm_lpvt_194df3105ad7148dcf2b98a91b5e727a=1502948718; __c=1501326829; lastCity=101020100; __g=-; Hm_lvt_194df3105ad7148dcf2b98a91b5e727a=1501326839; Hm_lpvt_194df3105ad7148dcf2b98a91b5e727a=1502954829; __l=r=https%3A%2F%2Fwww.google.com.hk%2F&l=%2F; __a=38940428.1501326829..1501326829.21.1.21.21")
	r.Set("cache-control", "no-cache")
	//r.Set("postman-token", "76554687-c4df-0c17-7cc0-5bf3845c9831")
}

/**
* 随机返回一个User-Agent
*/
func getAgent() string {
	agent  := [...]string{
		"Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:50.0) Gecko/20100101 Firefox/50.0",
		"Opera/9.80 (Macintosh; Intel Mac OS X 10.6.8; U; en) Presto/2.8.131 Version/11.11",
		"Opera/9.80 (Windows NT 6.1; U; en) Presto/2.8.131 Version/11.11",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; 360SE)",
		"Mozilla/5.0 (Windows NT 6.1; rv:2.0.1) Gecko/20100101 Firefox/4.0.1",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; The World)",
		"User-Agent,Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_8; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
		"User-Agent, Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Maxthon 2.0)",
		"User-Agent,Mozilla/5.0 (Windows; U; Windows NT 6.1; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	len := len(agent)
	return agent[r.Intn(len)]
}