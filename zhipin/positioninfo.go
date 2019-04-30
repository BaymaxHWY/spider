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


	item.CompanyName = handleCompanyName(dom.Find(".job-sec>.name").Text())

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

	printItem(item)
	parseResult.Item = item
	return parseResult, nil
}

func handleCompanyName(companyName string) (resultName string) {
	reCompanyName1 := regexp.MustCompile(`(.*)[技术 股份]?有限责任公司`)
	reCompanyName2 := regexp.MustCompile(`(.*)[技术 股份]?有限公司`)
	submatch1 := reCompanyName1.FindStringSubmatch(companyName)
	submatch2 := reCompanyName2.FindStringSubmatch(companyName)
	// fmt.Println("submatch:", submatch1, submatch2)
	sublen1 := len(submatch1)
	sublen2 := len(submatch2)

	if sublen1 <= 1 && sublen2 <= 1 {
		resultName = companyName
	}else if sublen1 > 1 && sublen2 <= 1{
		resultName = submatch1[1]
	}else if sublen1 <= 1 && sublen2 > 1 {
		resultName = submatch2[1]
	}else {
		if len(submatch1[1]) > len(submatch2[1]) {
			resultName = submatch2[1]
		}else {
			resultName = submatch1[1]
		}
	}
	// fmt.Println("new company: ",resultName )
	return
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
