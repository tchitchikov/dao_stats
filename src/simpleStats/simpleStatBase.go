package simpleStats

import (
	"math"
)

func S_和(值 []float64) (返回值 float64) {
	for _, x := range 值 {
		返回值 = 返回值 + x
	}
	return 返回值
}

func S_平均(值 []float64) (返回值 float64) {
	return S_和(值) / float64(len(值))
}

func S_方差(值 []float64) (返回值 float64) {
	平均 := S_平均(值)
	for _, x := range 值 {
		返回值 = 返回值 + math.Pow((x-平均), 2)
	}
	return 返回值 / float64(len(值)-1)
}

func S_标准偏差(值 []float64) (返回值 float64) {
	return math.Pow(S_方差(值), 0.5)
}
