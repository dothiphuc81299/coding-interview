package main

import "fmt"

func longestSubstringWithKDistinctChars(s string, k int) int {
	var (
		char       = make(map[string]int)
		start  int = 0
		maxLen     = 0
	)

	for end := range s {
		endChar := string(s[end])
		if _, ok := char[endChar]; !ok {
			char[endChar] = 0
		}

		char[endChar]++

		// shrink the window until there is no more than k distinct characters.
		for len(char) > k {
			startChart := string(s[start])
			char[startChart]--
			if char[startChart] == 0 {
				delete(char, startChart)
			}
			start++
		}
		if maxLen <= end-start+1 {
			maxLen = end - start + 1
		}
	}
	return maxLen
}

func main() {
	var s string = "araaci"
	var K int = 2
	aa := longestSubstringWithKDistinctChars(s, K)
	fmt.Println(aa)
}
