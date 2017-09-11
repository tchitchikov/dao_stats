package 数据

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

const google网址 = "https://www.google.com/finance/historical?q=%s&startdate=%s&enddate=%s&output=csv"

type Google数据Struct struct {
	Date   string
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume int64
}

// Google数据 获得股票市场数据
func Google数据(股票代号 string, 开始日期 string, 结束日期 string) (返回值 []Google数据Struct) {
	网址 := fmt.Sprintf(google网址, 股票代号, 开始日期, 结束日期)
	产量, 故障 := http.Get(网址)
	if 故障 != nil {
		fmt.Println(故障)
	}
	defer 产量.Body.Close()

	var 数据产量 []Google数据Struct
	读者 := csv.NewReader(产量.Body)
	迭代 := 0
	for {
		var 例 Google数据Struct
		行, 故障 := 读者.Read()
		if 故障 == io.EOF {
			break
		} else if 故障 != nil {
			log.Fatalf(" Failed to read csv ")
		}
		if 迭代 >= 1 {
			例 = 转换数据(行)
			数据产量 = append(数据产量, 例)
		}
		迭代 = 迭代 + 1
	}
	return 数据产量
}

func 转换数据(行 []string) (例 Google数据Struct) {
	例.Date = 行[0]
	例.Open, _ = strconv.ParseFloat(行[1], 64)
	例.High, _ = strconv.ParseFloat(行[2], 64)
	例.Low, _ = strconv.ParseFloat(行[3], 64)
	例.Close, _ = strconv.ParseFloat(行[4], 64)
	例.Volume, _ = strconv.ParseInt(行[5], 10, 64)
	return 例
}
