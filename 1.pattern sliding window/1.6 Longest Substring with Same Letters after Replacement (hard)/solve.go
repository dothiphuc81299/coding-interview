package main

import "fmt"

func characterReplacement(s string, k int) int {
	var (
		mostFreqLetter int = 0
		start          int = 0
		max            int = 0
		mp                 = make(map[string]int)
	)

	for end := 0; end < len(s); end++ {
		mp[string(s[end])]++
		if mp[string(s[end])] > mostFreqLetter {
			mostFreqLetter = mp[string(s[end])]
		}

		lettersToChange := (end - start + 1) - mostFreqLetter
		

		if lettersToChange > k {
			mp[string(s[start])]--
			start++
		}

		if end-start+1 > max {
			max = end - start + 1
		}

	}
	return max
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
