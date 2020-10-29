package top100

import "strings"

// 超时,字典包含重复时性能特别差
/*
"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
["a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"]
*/
func wordBreak1(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}
	for i := range wordDict {
		if strings.HasPrefix(s, wordDict[i]) && wordBreak1(s[len(wordDict[i]):], wordDict) {
			return true
		}
	}
	return false
}

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := range wordDict {
			if len(wordDict[j]) > i && strings.HasSuffix(s[:i], wordDict[j]) && dp[i-len(wordDict[j])] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
