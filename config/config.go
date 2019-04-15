package config

const ZHIPIN = "https://www.zhipin.com"

// 城市-url
var CityList  = [...]CityUrl{
	{Name: "北京", Url: "https://www.zhipin.com/c101010100"},
	{Name: "上海", Url: "https://www.zhipin.com/c101020100"},
	{Name: "广州", Url: "https://www.zhipin.com/c101280100"},
	{Name: "深圳", Url: "https://www.zhipin.com/c101280600"},
	{Name: "杭州", Url: "https://www.zhipin.com/c101210100"},
	{Name: "天津", Url: "https://www.zhipin.com/c101030100"},
	{Name: "西安", Url: "https://www.zhipin.com/c101110100"},
	{Name: "苏州", Url: "https://www.zhipin.com/c101190400"},
	{Name: "武汉", Url: "https://www.zhipin.com/c101200100"},
	{Name: "厦门", Url: "https://www.zhipin.com/c101230200"},
	{Name: "长沙", Url: "https://www.zhipin.com/c101250100"},
	{Name: "成都", Url: "https://www.zhipin.com/c101270100"},
}

var PositionList = []string{"Golang", "PHP", "Node.js", "Java", "C++", "C#", "Python", "Ruby"}

var MySQL = DBConfig{
	DBname: "mysql",
	Username: "root",
	Password: "",
	Protocol: "tcp",
	Address: "127.0.0.1:3306",
	Tablename: "zhaopin",
}

