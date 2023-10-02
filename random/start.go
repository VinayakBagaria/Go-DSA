package random

import "fmt"

const separator = "--------------------------------------"

func DoWork() {
	fmt.Println(separator)
	fmt.Println(wordBreakBottomUp("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreakTopDown("leetcode", []string{"leet", "code"}))
}
