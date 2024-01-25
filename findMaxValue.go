package main

import "fmt"

func findMaxValue(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	maxValue := arr[0]
	for _, value := range arr {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}

func main() {
	arr := []int{3, 5, 1, 9, 2}
	fmt.Println(findMaxValue(arr)) // 9
}
