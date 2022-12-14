package introduction

func AverageOfSubarrayOfSizeK(K int, arr []int) []float64 {
	var (
		result              = make([]float64, len(arr)-K+1)
		windowSum   float64 = 0
		windowStart int     = 0
		windowEnd   int     = 0
	)

	for i := windowEnd; i < len(arr); i++ {
		windowSum += float64(arr[i])
		if i >= K-1 {
			result[windowStart] = windowSum / float64(K)
			windowSum = windowSum - float64(arr[windowStart])
			windowStart++
		}
	}

	return result
}


