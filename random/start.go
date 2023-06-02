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
	fmt.Println(separator)
	DoMedianFromDataStream()
	fmt.Println(separator)
	DoTopKFrequentWords()
	fmt.Println(separator)
	DoSortCharsByFrequency()
}
