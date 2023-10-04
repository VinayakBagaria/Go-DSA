package random

import (
	"fmt"
)

func minInteger(num string, k int) string {
    n := len(num)

	for i := 0; i < n && k > 0; i++ {
		pos := i

		for j := i + 1; j < n; j++ {
			if j - i > k {
				break
			}

			if num[j] < num[pos] {
				pos = j
			}
		}

		for pos > i {
			num = swapString(num, pos - 1, pos)
			fmt.Printf("num: %s\n", num)
			pos--
			k--
		}
	}

	return num
}

func swapString(s string, i, j int) string {
	ithChar := s[i]
	jthChar := s[j]

	return s[0:i] + string(jthChar) + s[i+1:j] + string(ithChar) + s[j+1:]
}

func DoMinPossibleIntAfterKSwaps() {
	fmt.Println(minInteger("4321", 4))
}