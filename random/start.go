package random

import "fmt"

const separator = "--------------------------------------"

func DoWork() {
	fmt.Println(separator)
	DoProductOfArrayExceptSelf()
	fmt.Println(separator)
	DoLisLength()
	fmt.Println(separator)
	DoDeepFilter()
	fmt.Println(separator)
	DoKthLargestElement()
}
