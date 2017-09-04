package data

const google网址 = "https://www.google.com/finance/historical?q=%s&startdate=%s&enddate=%s"

func Google数据(股票代号 String, 开始日期 String, 结束日期 String) {
	网址 := fmt.Sprintf(googleURL, 股票代号， 开始日期， 结束日期)
}
