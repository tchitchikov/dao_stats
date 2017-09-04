package 数据

import "fmt"

const google网址 = "https://www.google.com/finance/historical?q=%s&startdate=%s&enddate=%s"

// Google数据 获得股票市场数据
func Google数据(股票代号 string, 开始日期 string, 结束日期 string) {
	网址 := fmt.Sprintf(google网址, 股票代号, 开始日期, 结束日期)
	fmt.Println(网址)
}
