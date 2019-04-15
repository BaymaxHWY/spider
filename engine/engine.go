package engine

import (
	"bishe/spider/config"
	"bishe/spider/fetcher"
	"bishe/spider/model"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func Run(tasklist []config.Task) {
	for len(tasklist) > 0 {
		task := tasklist[0]
		tasklist = tasklist[1:]
		data, err := fetcher.Fetcher(task.Url)
		fmt.Printf("fetcher : %s\n", task.Url)
		if err != nil {
			log.Printf("fetcher error :%s", err.Error())
		}
		parseResult, err := task.Parse(data)
		if err != nil {
			log.Printf("parse error : %s", err.Error())
		}
		// zhipin.PrintItem(parseResult.Item)
		// fmt.Println("isStore:", parseResult.IsStore)
		if parseResult.IsStore {
			// 数据清洗
			data, err := cleanData(parseResult.Item)
			if err != nil {
				log.Printf("data clean error: %s", err.Error())
			}
			// 数据合格再进行存储
			printNewData(data)
			//model.Insert(data)
		}
		if parseResult.Tasks == nil {
			continue
		}
		tasklist = append(tasklist, parseResult.Tasks...)
	}
}
// 数据清洗
// 将解析出来的数据进行进一步的整理
// Salary: 18k-35k => {low: 18000, avg: (35000 + 18000) / 2, high: 35000}
// 根据工作年限划分招聘等级：
// 经验不限：0；
// 应届生：1；
// 1年以内：2；
// 1-3年：3；
// 3-5年：4；
// 5-10年：5；
// 10年以上：6；
// 缺失信息则为无效数据
func cleanData(oldData config.ZhiPiItem) (model.PositionInfo, error) {
	var newData = cleanMiss(oldData)
	// Salary清洗
	salaryStr := oldData.Salary
	salaryRe := regexp.MustCompile(`([0-9]+)k-([0-9]+)k`)
	salarymatch := salaryRe.FindStringSubmatch(salaryStr)
	if len(salarymatch) < 2 {
		return newData, errors.New("invalid data")
	}
	low, _ := strconv.Atoi(salarymatch[1])
	high, _ := strconv.Atoi(salarymatch[2])
	salary := config.Salary{
		Low:low*1000,
		Avg:(low + high)*1000/2,
		High:high*1000,
	}
	newData.Salary.Low, newData.Salary.Avg, newData.Salary.High = salary.Low, salary.Avg, salary.High
	// 划分招聘等级
	switch oldData.WorkYear {
	case "经验不限": newData.Level = 0
	case "应届生": newData.Level = 1
	case "1年以内": newData.Level = 2
	case "1-3年": newData.Level = 3
	case "3-5年": newData.Level = 4
	case "5-10年": newData.Level = 5
	case "10年以上": newData.Level = 6
	default:
		return newData, errors.New("level divide failed")
	}
	newData.WorkYear = oldData.WorkYear
	newData.Pid = oldData.Pid

	return newData, nil
}
// 缺失值检查
func cleanMiss(oldData config.ZhiPiItem) model.PositionInfo {
	var newData = model.PositionInfo{}
	isMiss := false
	if oldData.Positon != "" {
		newData.Positon = oldData.Positon
	} else {
		isMiss = true
	}
	if oldData.City != "" {
		newData.City = oldData.City
	} else {
		isMiss = true
	}
	if oldData.Education != "" {
		newData.Education = oldData.Education
	} else {
		isMiss = true
	}
	if oldData.CompanyName != "" {
		newData.CompanyName = oldData.CompanyName
	} else {
		isMiss = true
	}
	if oldData.FinanceStage != "" {
		newData.FinanceStage = oldData.FinanceStage
	} else {
		isMiss = true
	}
	if oldData.CompanySize != "" {
		newData.CompanySize = oldData.CompanySize
	} else {
		isMiss = true
	}
	if oldData.Industry != "" {
		newData.Industry = oldData.Industry
	} else {
		isMiss = true
	}
	if oldData.Location != "" {
		newData.Location = oldData.Location
	} else {
		isMiss = true
	}
	if oldData.Detail != "" {
		newData.Detail = oldData.Detail
	} else {
		isMiss = true
	}
	if isMiss {
		newData.IsMiss = true
	}
	return newData
}

func printNewData(item model.PositionInfo) {
	fmt.Println("{")
	fmt.Printf("Pid: %s,\n", item.Pid)
	fmt.Printf("Positon: %s,\n", item.Positon)
	fmt.Printf("Salary: %v,\n", item.Salary)
	fmt.Printf("City: %s,\n", item.City)
	fmt.Printf("WorkYear: %s,\n", item.WorkYear)
	fmt.Printf("WorkYear: %d,\n", item.Level)
	fmt.Printf("Education: %s,\n", item.Education)
	fmt.Printf("CompanyName: %s,\n", item.CompanyName)
	fmt.Printf("FinanceStage: %s,\n", item.FinanceStage)
	fmt.Printf("CompanySize: %s,\n", item.CompanySize)
	fmt.Printf("Industry: %s,\n", item.Industry)
	fmt.Printf("Detail: %s,\n", item.Detail)
	fmt.Printf("Location: %s,\n", item.Location)
	fmt.Printf("IsMiss: %v,\n", item.IsMiss)
	fmt.Println("}")
}