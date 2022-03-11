package main

import (
	"data-structure-algorithm/utils"
	"fmt"
	"time"
)

// SearchString 线性查找String元素
func SearchString(data []string, target string) int {
	for idx, value := range data {
		if value == target {
			return idx

		}
	}
	return -1
}

// SearchInt64 线性查找int元素
func SearchInt64(data []int, target int) int {

	for i := 0; i < len(data); i++ {
		if target == data[i] {
			return i
		}
	}

	return -1
}

// FewNumber 求解约数
//对target来说 i * target / i 是成对出现的
func FewNumber(target int64) []int64 {
	res := make([]int64, 0)
	var i int64 = 1
	for ; i*i < target; i++ {
		if target%i == 0 && i != target/i {
			res = append(res, i, target/i)
		}
	}

	return res
}

func main() {

	//数据规模为: 100000, 花费4.905ms
	//数据规模为: 100000, 花费0.005s
	//数据规模为: 1000000, 花费53.709ms
	//数据规模为: 1000000, 花费0.054s
	//数据规模为: 10000000, 花费537.773ms
	//数据规模为: 10000000, 花费0.538s
	dataSize := []int{100000, 1000000, 10000000}
	for _, n := range dataSize {
		data := utils.GenerateOrderArray(n)
		start := time.Now().UnixNano()
		for i := 0; i < 100; i++ {
			SearchInt64(data, n)
		}
		end := time.Now().UnixNano()
		fmt.Printf("数据规模为: %d, 花费%.3fms\n", n, (float64(end)-float64(start))/1e6)
		fmt.Printf("数据规模为: %d, 花费%.3fs\n", n, (float64(end)-float64(start))/1e9)
	}

}
