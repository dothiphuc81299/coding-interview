package main

import "fmt"

func longestOnes(nums []int, k int) int {
	var (
		start          int = 0
		max            int = 0
		mostFreqNumber int = 0
		mp                 = make(map[int]int)
	)

	for end := 0; end < len(nums); end++ {
		if nums[end] == 1 {
			mp[1]++
		}

		if mp[1] > mostFreqNumber {
			mostFreqNumber = mp[1]
		}

		numberToChange := end - start + 1 - mostFreqNumber
		if numberToChange > k {
			if nums[start] == 1 {
				mp[1]--

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
