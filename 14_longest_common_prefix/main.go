package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	prefix := strs[0]
	var flag bool

	for i := 1; i < len(strs); i++ {
		prifixLength := min(strs[i], prefix)
		curr := 0
		flag = false
		for j := 0; j < prifixLength; j++ {
			if prefix[j] == strs[i][j] {
				curr = j
				flag = true
				continue
			}
			// fmt.Println(curr)
			break
		}
		if flag {
			prefix = prefix[:curr+1]
		} else {
			prefix = prefix[:curr]
		}
	}
	if flag {
		return prefix
	}
	return ""
}

func min(x, y string) int {
	xl, yl := len(x), len(y)
	if xl > yl {
		return yl
	}
	return xl
}

func main() {
	x := []string{"c", "acc", "ccc"}
	fmt.Println(longestCommonPrefix(x))
}
