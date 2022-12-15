package main

import "fmt"

func characterReplacement(s string, k int) int {
	// var (
	// 	start  int = 0
	// 	maxLen     = 0
	// 	char       = make(map[string]int)
	// )

	// for end := range s {
	// 	endChar := string(s[end])
	// 	// if _, ok := char[endChar]; !ok {
	// 	// 	char[endChar] = 0
	// 	// }
	// 	char[endChar]++
	// 	fmt.Println("char[endChar]", char[endChar])
	// 	fmt.Println("len", len(char))
	// 	for len(char) > k {
	// 		startChar := string(s[start])
	// 		char[startChar]--
	// 		if char[startChar] == 0 {
	// 			delete(char, startChar)
	// 		}

	// 		start++
	// 	}

	// 	fmt.Println("len", char)

	// 	if maxLen < len(char)+k {
	// 		maxLen = len(char) + k
	// 	}
	// }

	// return maxLen
	return 0
}

func main() {
	var s = "abab"
	var k = 2
	fmt.Println(characterReplacement(s, k))

	var ss = "AABABBA"
	var kk = 1
	fmt.Println(characterReplacement(ss, kk))

	var sss = "AAAA"
	var kkk = 0
	fmt.Println(characterReplacement(sss, kkk))
}
