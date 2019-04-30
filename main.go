package main

import (
	"bishe/spider/config"
	"bishe/spider/zhipin"
	"bishe/spider/engine"
	"strconv"
)

// https://www.zhipin.com/c101010100/?query=golang&page=3
// https://www.zhipin.com/c101010100/?query=web&page=1
// https://www.zhipin.com/c101010100
// 输入城市url，职位，生成task列表
func generateTask(cityList []config.CityUrl, positionList []string, pageStart, pageEnd int) []config.Task {
	var taskList []config.Task
	if pageStart < 1 {
		pageStart = 1
	}
	for _, city := range cityList {
		for _, position := range positionList {
			for i := pageStart; i <= pageEnd; i++ {
				tUrl := city.Url + `/?query=` + position + `&page=` + strconv.Itoa(i)
				//fmt.Println(tUrl)
				task := config.Task{
					Url:   tUrl,
					Parse: zhipin.ParsePositionList,
				}
				taskList = append(taskList, task)
			}
		}

	}
	return taskList
}

// 交叉编译 env GOOS=linux GOARCH=amd64 go build .
func main() {
	taskList := generateTask(config.CityList[:], config.PositionList[:],5,10)
	engine.Run(taskList)
}
