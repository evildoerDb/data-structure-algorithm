package utils

// GenerateOrderArray 生成一个长度为n的有序数组
func GenerateOrderArray(n int) []int {
	res := make([]int, 0)
	for i := 0; i < n; i++ {
		res = append(res, i)
	}

	return res
}
