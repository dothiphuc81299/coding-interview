package smallestsubarraywithagivensumeasy

func minSubArrayLen(target int, nums []int) int {
	var (
		minLen      int = 0
		windowSum   int = 0
		windowStart     = 0
		temMin          = len(nums)
	)

	for windowEnd := 0; windowEnd < len(nums); windowEnd++ {
		windowSum += nums[windowEnd]
		for windowSum >= target {
			if windowEnd-windowStart+1 <= temMin {
				minLen = windowEnd - windowStart + 1
				temMin = minLen
			}
			windowSum -= nums[windowStart]
			windowStart++
		}
	}
	return minLen

}

// Time Complexity
// The time complexity of the above algorithm will be O(N). The outer for loop runs for all elements, and the inner while loop processes each element only once; therefore, the time complexity of the algorithm will be O(N+N), which is asymptotically equivalent to O(N).

// Space Complexity
// The algorithm runs in constant space O(1).