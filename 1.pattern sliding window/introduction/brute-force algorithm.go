package introduction

import "fmt"

func AverageOfSubarrayOfSizeK(K int, arr []int) []float64 {
	var result = make([]float64, len(arr)-K+1)
	for i := 0; i <= len(arr)-K; i++ {
		var sum float64 = 0
		for j := i; j < i+K; j++ {
			sum += float64(arr[j])
		}
		result[i] = sum / float64(K)
	}
	return result
}

func main() {
	var (
		ar = []int{1, 3, 2, 6, -1, 4, 1, 8, 2}
		K  = 5
	)

	res := AverageOfSubarrayOfSizeK(K, ar)
	fmt.Println(res)
}