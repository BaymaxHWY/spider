package main

import (
	"bishe/spider/config"
	"bishe/spider/zhipin"
	"strconv"
	"bishe/spider/engine"
)
// https://www.zhipin.com/c101010100/?query=golang&page=3
// https://www.zhipin.com/c101010100/?query=web&page=1
// https://www.zhipin.com/c101010100
// 输入城市url，职位，生成task列表
func generateTask(cityList []config.CityUrl, positionList []string) []config.Task {
	var taskList []config.Task
	for _, city := range cityList{
		for _, position := range positionList{
			for i := 1; i < 2; i++ {
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

func main() {
	taskList := generateTask(config.CityList[1:3], config.PositionList[1:3])
	engine.Run(taskList)
}

/*<p>北京<em class="dolt"></em>3-5年<em class="dolt"></em>本科</p>*/