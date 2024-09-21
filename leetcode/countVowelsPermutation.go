// https://leetcode.com/problems/count-vowels-permutation/description/
package leetcode

import "fmt"

func countVowelPermutation(n int) int {
	dp := make([][]int, 26)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var solve func(byte, int) int
	solve = func(prev byte, left int) int {
		if left == 0 {
			return 1
		}

		prevInt := prev - 'a'

		if dp[prevInt][left] != -1 {
			return dp[prevInt][left]
		}

		total := 0
		if prev == 'a' {
			total += solve('e', left-1)
		} else if prev == 'e' {
			total += solve('a', left-1)
			total += solve('i', left-1)
		} else if prev == 'i' {
			total += solve('a', left-1)
			total += solve('e', left-1)
			total += solve('o', left-1)
			total += solve('u', left-1)
		} else if prev == 'o' {
			total += solve('i', left-1)
			total += solve('u', left-1)
		} else if prev == 'u' {
			total += solve('a', left-1)
		}

		total = total % MOD
		dp[prevInt][left] = total
		return total
	}

	total := 0
	initial := "aeiou"
	for i := range initial {
		total += solve(initial[i], n-1)
	}
	return total % MOD
}

func DoCountVowelsPermutation() {
	fmt.Println(countVowelPermutation(1))
	fmt.Println(countVowelPermutation(2))
	fmt.Println(countVowelPermutation(5))
}
