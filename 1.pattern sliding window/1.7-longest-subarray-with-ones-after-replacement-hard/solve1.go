package main

import "fmt"

func longestOnes(nums []int, k int) int {
	var (
		max       int = 0
		zeroCount int = 0
		start     int = 0
	)

	for end := 0; end < len(nums); end++ {
		if nums[end] == 0 {
			zeroCount++
		}

		for zeroCount > k {
			if nums[start] == 0 {
				zeroCount--
			}

			start++
		}

		if max < end-start+1 {
			max = end - start + 1
		}
	}

	return max

}

func main() {
	var nums = []int{
		0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1,
	}
	k := 3
	fmt.Println(longestOnes(nums, k))

	var numss = []int{
		1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0,
	}
	kk := 2
	fmt.Println(longestOnes(numss, kk))

}
