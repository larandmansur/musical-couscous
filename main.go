package main

import (
	"fmt"
	"math"
)

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = int(math.Max(float64(dp[i]), float64(dp[j]+1)))
			}
		}
	}
	maxLength := 0
	for _, length := range dp {
		if length > maxLength {
			maxLength = length
		}
	}
	return maxLength
}

func main() {
	numbers := []int{10, 9, 2, 5, 3, 7, 101, 18}
	length := lengthOfLIS(numbers)
	fmt.Println("Длина наибольшей возрастающей подпоследовательности для чисел", numbers, ":", length)
}
