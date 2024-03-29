package heap

import "fmt"

const separator = "--------------------------------------"

func DoWork() {
	fmt.Println(separator)
	StartInsert()
	fmt.Println(separator)
	StartDelete()
	fmt.Println(separator)
	StartHeapify()
	fmt.Println(separator)
	StartInbuilt()
}
