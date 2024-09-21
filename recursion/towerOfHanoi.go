package recursion

import "fmt"

func solve(n int, from, to, aux string) int {
	if n == 0 {
		return 0
	}

	count := solve(n-1, from, aux, to)

	fmt.Printf("moved disk %d from rod %s to rod %s\n", n, from, to)
	count++

	count += solve(n-1, aux, to, from)
	return count
}

func DoTowerOfHanoi() {
	fmt.Println(solve(3, "1", "2", "3"))
}
