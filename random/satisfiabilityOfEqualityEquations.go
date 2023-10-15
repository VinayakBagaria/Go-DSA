// https://leetcode.com/problems/satisfiability-of-equality-equations/description/
package random

import "fmt"

func equationsPossible(equations []string) bool {
	parent := make([]int, 26)
	for i := 0; i < 26; i++ {
		parent[i] = i
	}

	var find func(int) int
	find = func(u int) int {
		if parent[u] == u {
			return u
		}
		parent[u] = find(parent[u])
		return parent[u]
	}

	var union = func(u, v int) {
		p1 := find(u)
		p2 := find(v)
		parent[p2] = p1
	}

	for _, e := range equations {
		if e[1] == '=' {
			u := getIndex(e[0])
			v := getIndex(e[3])
			union(u, v)
		}
	}

	for _, e := range equations {
		if e[1] == '!' {
			u := getIndex(e[0])
			v := getIndex(e[3])
			if find(u) == find(v) {
				return false
			}
		}
	}

	return true
}

func getIndex(c byte) int {
	return int(c - 'a')
}

func DoSatisfiabilityOfEqualityEquations() {
	fmt.Println(equationsPossible([]string{"a==b", "b!=a"}))
	fmt.Println(equationsPossible([]string{"b==a", "a==b"}))
}
