package random

import "fmt"

const separator = "--------------------------------------"

func DoWork() {
	fmt.Println(separator)
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
}
