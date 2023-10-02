package random

import "fmt"

const separator = "--------------------------------------"

func DoWork() {
	fmt.Println(separator)
	fmt.Println(wordBreak2("catsanddog", []string{"cat","cats","and","sand","dog"}))
}
