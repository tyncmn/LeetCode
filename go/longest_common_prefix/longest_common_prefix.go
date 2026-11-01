package main

import "fmt"

func main() {
	fmt.Println(LongestCommonPrefix([]string{"flower", "flow", "flight"}))
}

func LongestCommonPrefix(strs []string) string {
	listLen := len(strs)
	switch listLen {
	case 0:
		return ""
	case 1:
		return strs[0]
	}
	prefix := ""
	matchedStr := strs[0]
	for i := 1; i < listLen; i++ {
		minLen := len(strs[i])
		if len(matchedStr) < minLen {
			minLen = len(matchedStr)
		}
		for j := 0; j < minLen; j++ {
			if matchedStr[j] == strs[i][j] {
				prefix += string(strs[i][j])
			} else {
				break
			}
		}
		fmt.Println("Matched prefix:", prefix)
		if prefix == "" {
			return ""
		} else {
			matchedStr = prefix
		}
		if i != listLen-1 {
			prefix = ""
		}
	}
	return prefix
}
