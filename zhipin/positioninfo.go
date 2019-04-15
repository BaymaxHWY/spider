package zhipin

import (
	"bishe/spider/config"
	"bishe/spider/util"
	"bytes"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strings"
)

func ParsePosition(data []byte, pid string) (*config.ParseResult, error) {
	var parseResult = &config.ParseResult{
		IsStore: true}
	var item = config.ZhiPiItem{
		Pid: pid}

	dom, err := goquery.NewDocumentFromReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	item.Positon = dom.Find(".info-primary > .name > h1").Text()
	item.Salary = util.Trim(dom.Find(".info-primary > .name > .salary").Text())

	p, _ := dom.Find(".info-primary > p").Html()
	reP := regexp.MustCompile(`([^<]+)<em class="dolt"></em>([^<]+)<em class="dolt"></em>([^>]+)`)
	resP := reP.FindStringSubmatch(p)
	if len(resP) < 4 {
		return nil, errors.New("regexp fail")
	}
	item.City = resP[1]
	item.WorkYear = resP[2]
	item.Education = resP[3]

	item.CompanyName = dom.Find(".job-sec>.name").Text()
	//fmt.Println(dom.Find(".job-sec>.name").Html())
	dom.Find(".sider-company p").Each(func(i int, s *goquery.Selection) {
		switch i {
		case 1: item.FinanceStage = s.Text()
		case 2: item.CompanySize = s.Text()
		case 3: item.Industry = s.Text()
		}
	})
	item.Detail, _ = dom.Find(".detail-content>.job-sec>.text").Html()
	item.Detail = strings.TrimSpace(item.Detail)
	item.Detail = strings.Replace(item.Detail, `<br/>`, "\n", -1)
	item.Location = dom.Find(".job-location > .location-address").Text()

	//printItem(item)
	parseResult.Item = item
	return parseResult, nil
}

func printItem(item config.ZhiPiItem) {
	fmt.Println("{")
	fmt.Printf("Positon: %s,\n", item.Positon)
	fmt.Printf("Salary: %s,\n", item.Salary)
	fmt.Printf("City: %s,\n", item.City)
	fmt.Printf("WorkYear: %s,\n", item.WorkYear)
	fmt.Printf("Education: %s,\n", item.Education)
	fmt.Printf("CompanyName: %s,\n", item.CompanyName)
	fmt.Printf("FinanceStage: %s,\n", item.FinanceStage)
	fmt.Printf("CompanySize: %s,\n", item.CompanySize)
	fmt.Printf("Industry: %s,\n", item.Industry)
	fmt.Printf("Detail: %s,\n", item.Detail)
	fmt.Printf("Location: %s,\n", item.Location)
	fmt.Println("}")
}
