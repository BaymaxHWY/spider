package config

type CityUrl struct{
	Name string
	Url string
}

type Salary struct {
	Low int
	Avg int
	High int
}

type DBConfig struct {
	DBname string
	Username string
	Password string
	Protocol string
	Address string
	Tablename string
}

/*
招聘信息
Pid 去重唯一id
Postion 职位
Salary 薪资
City 城市
WorkYear 工作年限
Education 学历要求
CompanyName 公司名称
Industry 所在行业
FinanceStage 融资阶段
CompanySize 公司规模
Detail 招聘描述
Location 工作地点
*/
type ZhiPiItem struct{
	Pid string
	Positon string
	Salary string
	City string
	WorkYear string
	Education string
	CompanyName string
	Industry string
	FinanceStage string
	CompanySize string
	Detail string
	Location string
}
// 解析器
type ParseFunc func([]byte) (*ParseResult, error)

/*
抓取任务
*/
type Task struct{
	Url string
	Parse ParseFunc
}
/*
解析结果
*/
type ParseResult struct {
	IsStore bool
	Item ZhiPiItem
	Tasks []Task
}

func NilParse(data []byte) (*ParseResult, error) {
	return nil, nil
}