package main

import (
	"data"
	"fmt"
	"simpleStats"
)

func main() {
	价值表 := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 17.0}
	fmt.Println(simpleStats.S_平均(价值表))
	fmt.Println(simpleStats.S_方差(价值表))
	fmt.Println(simpleStats.S_标准偏差(价值表))
	data.GoogleData()
}
