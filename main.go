package main

import (
	"fmt"
	"simpleStats"

	"数据"
)

func main() {
	价值表 := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 17.0}
	fmt.Println(simpleStats.S_平均(价值表))
	fmt.Println(simpleStats.S_方差(价值表))
	fmt.Println(simpleStats.S_标准偏差(价值表))
	Google数据 := 数据.Google数据("GOOG", "2017-7-1", "2017-7-31")
	var GoogleClose数据 []float64
	for _, 行 := range Google数据 {
		GoogleClose数据 = append(GoogleClose数据, 行.Close)
	}
	fmt.Println(simpleStats.S_标准偏差(GoogleClose数据))
}
