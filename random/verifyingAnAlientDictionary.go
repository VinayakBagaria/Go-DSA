// https://leetcode.com/problems/verifying-an-alien-dictionary/description/
package random

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

			minLength := findMin(len(first), len(second))

			for k := 0; k < minLength; k++ {
				if orderDict[first[k]] < orderDict[second[k]] {
					break
				} else if orderDict[first[k]] > orderDict[second[k]] {
					return false
				} else if k == minLength - 1 && len(first) > len(second) {
                    return false
                }
			}
		}
    }

    return true
}

func findMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func DoIsAlienSorted() {
	fmt.Println(isAlienSorted([]string{"hello","leetcode"}, "hlabcdefgijkmnopqrstuvwxyz"))
	fmt.Println(isAlienSorted([]string{"word","world","row"}, "worldabcefghijkmnpqstuvxyz"))
	fmt.Println(isAlienSorted([]string{"apple", "app"}, "worldabcefghijkmnpqstuvxyz"))
}