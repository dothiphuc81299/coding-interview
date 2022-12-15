package main

import "fmt"

func fruitsIntoBaskets(fruits []string) int {
	var (
		maxLen int = 0
		char       = make(map[string]int)
		start  int = 0
	)

	for end := range fruits {
		// insert fruits until we have 2 distinct fruits.
		endChar := fruits[end]
		if _, ok := char[endChar]; ok {
			char[endChar] = 0
		}
		char[endChar]++

		// shrink the window until there is no more than 2 distinct fruits.
		for len(char) > 2 {
			startChar := fruits[start]

			// decrement the frequency of the one going out of the window and
			// remove if its frequency is zero.
			char[startChar]--
			if char[startChar] == 0 {
				delete(char, startChar)
			}

			// increase the start index to move the window ahead by one element.
			start++
		}

		if maxLen < end-start+1 {
			maxLen = end - start + 1
		}
	}

	return maxLen
}

func main() {
	var arr = []string{"A", "B", "C", "A", "C"}
	aa := fruitsIntoBaskets(arr)
	fmt.Println(aa)
	var arr1 = []string{"A", "B", "C", "B", "B", "C"}
	aa1 := fruitsIntoBaskets(arr1)
	fmt.Println(aa1)

}


// The time complexity of the above algorithm will be O(N)O(N) where ‘N’ is the number of characters in the input array. The outer for loop runs for all characters and the inner while loop processes each character only once, therefore the time complexity of the algorithm will be O(N+N)O(N+N) which is asymptotically equivalent to O(N)O(N).

// Space Complexity
// The algorithm runs in constant space O(1)O(1) as there can be a maximum of three types of fruits stored in the frequency map.

