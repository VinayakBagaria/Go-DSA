package random

import "fmt"

const separator = "--------------------------------------"

func DoWork() {
	fmt.Println(separator)
	DoKthLargestElement()
	fmt.Println(separator)
	DoMedianFromDataStream()
	fmt.Println(separator)
	DoKthSmallestMatrixRowSum()
	fmt.Println(separator)
	DoSlidingWindowMaximum()
}
