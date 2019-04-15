package zhipin

import (
	"bishe/spider/config"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"github.com/PuerkitoBio/goquery"
)

// https://www.zhipin.com/job_detail/e8ab00b1079407f71XZ52N26GFA~.html
func ParsePositionList(data []byte) (*config.ParseResult, error) {
	var parseResult = &config.ParseResult{
		IsStore: false,
	}
	dom, err := goquery.NewDocumentFromReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	dom.Find(".job-primary > .info-primary > .name > a").Each(func(i int, s *goquery.Selection) {
		val, ok := s.Attr("href")
		if !ok {
			return
		}
		task := config.Task{
			Url: config.ZHIPIN + val,
			// val = /job_detail/[9f725c45aa83d9751nN83t-8F1E~].html  提取出来作为去重id
			Parse: func(data []byte) (result *config.ParseResult, e error) {
				pid := getMd5(val)
				return ParsePosition(data, pid)
			},
		}
		//fmt.Printf("get url: %s\n", config.ZHIPIN + val)
		parseResult.Tasks = append(parseResult.Tasks, task)
	})
	return parseResult, nil
}

func getMd5(str string) string {
	t := md5.Sum([]byte(str))
	return hex.EncodeToString(t[:])
}