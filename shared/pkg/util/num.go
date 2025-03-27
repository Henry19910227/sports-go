package util

func SubtractWithFloor(a int, b int) int {
	ans := a - b
	if ans < 0 {
		return 0
	}
	return ans
}

func IntArrayToInt32Array(items []int) []int32 {
	arr := make([]int32, 0)
	for _, item := range items {
		arr = append(arr, int32(item))
	}
	return arr
}
