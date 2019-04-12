package engine

import (
	"bishe/spider/config"
	"bishe/spider/fetcher"
	"fmt"
	"learning/spider/model"
	"log"
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
		if parseResult.Tasks == nil {
			continue
		}
		tasklist = append(tasklist, parseResult.Tasks...)
		if parseResult.IsStore {
			// 存储
			//data := dataClean(parseResult.Item)
		}
	}
}
// 数据清洗
// 将解析出来的数据进行进一步的整理
func dataClean(oldData config.ZhiPiItem) model.Position {
	var newData = model.Position{}
	return newData
}