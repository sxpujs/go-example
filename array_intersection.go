package main

import "fmt"

// 两个数组计算交集，结果要去重。
func main() {
	a := []int{1, 3, 5, 5, 20, 3, 99}
	b := []int{1, 4, 5, 5, 20, 100, 99}
	result := intersection(a, b)
	fmt.Println(result)
}

func intersection(a, b []int) []int {
	m := make(map[int]struct{})
	var result []int

	for _, e := range a {
		m[e] = struct{}{}
	}
	for _, e := range b {
		if _, ok := m[e]; ok {
			result = append(result, e)
			delete(m, e) // 从map中删除该元素，保证结果列表去重。
		}
	}
	return result
}
