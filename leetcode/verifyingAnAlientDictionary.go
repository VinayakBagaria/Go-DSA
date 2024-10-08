// https://leetcode.com/problems/verifying-an-alien-dictionary/description/
package leetcode

import "fmt"

func isAlienSorted(words []string, order string) bool {
	orderDict := make(map[byte]int)
	for i := range order {
		orderDict[order[i]] = i
	}

	for i := range words {
		for j := i + 1; j < len(words); j++ {
			first := words[i]
			second := words[j]

			minLength := min(len(first), len(second))

			for k := 0; k < minLength; k++ {
				if orderDict[first[k]] < orderDict[second[k]] {
					break
				} else if orderDict[first[k]] > orderDict[second[k]] {
					return false
				} else if k == minLength-1 && len(first) > len(second) {
					return false
				}
			}
		}
	}

	return true
}

func DoIsAlienSorted() {
	fmt.Println(isAlienSorted([]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz"))
	fmt.Println(isAlienSorted([]string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz"))
	fmt.Println(isAlienSorted([]string{"apple", "app"}, "worldabcefghijkmnpqstuvxyz"))
}
