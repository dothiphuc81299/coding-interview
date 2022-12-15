package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	var (
		maxLen int = 0
		start  int = 0
		char       = make(map[string]int)
	)

	for end := range s {
		endChar := string(s[end])
		if _, ok := char[endChar]; !ok {
			char[endChar] = 0
		}
		char[endChar]++
		for char[endChar] > 1 {
			startChar := string(s[start])
			char[startChar]--
			if char[startChar] == 0 {
				delete(char, startChar)
			}

			start++
		}

		if maxLen < len(char) {
			maxLen = len(char)
		}
	}
	return maxLen
}

func main() {
	var s string = "aabccbb"
	fmt.Println(lengthOfLongestSubstring(s))
	var ss string ="abbbb"
	fmt.Println(lengthOfLongestSubstring(ss))
	var sss string ="abccde"
	fmt.Println(lengthOfLongestSubstring(sss))
}


// Time Complexity
// The time complexity of the above algorithm will be O(N) where ‘N’ is the number of characters in the input string.

// Space Complexity
// The space complexity of the algorithm will be O(K) where K is the number of distinct characters in the input string. This also means K<=N, because in the worst case, the whole string might not have any repeating character so the entire string will be added to the HashMap. Having said that, since we can expect a fixed set of characters in the input string (e.g., 26 for English letters), we can say that the algorithm runs in fixed space O(1); in this case, we can use a fixed-size array instead of the HashMap.