package main

func maxSumOfContiguousSubArray(K int, arr []int) int {
	var (
		maxSum      int = 0
		result          = make([]int, len(arr)-K+1)
		windowStart     = 0
		sum         int = 0
	)

	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		sum += arr[windowEnd]
		if windowEnd >= K-1 {
			result[windowStart] = sum
			if result[windowStart] > maxSum {
				maxSum = result[windowStart]
			}
			sum = sum - arr[windowStart]
			windowStart++
		}
	}

	return maxSum
}

// Time Complexity
// The time complexity of the above algorithm will be O(N).

// Space Complexity
// The algorithm runs in constant space O(1).