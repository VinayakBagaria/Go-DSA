// https://leetcode.com/problems/accounts-merge/description/
package leetcode

import (
	"fmt"
	"sort"
)

func accountsMerge(accounts [][]string) [][]string {
	adj := make(map[string][]string)
	mailToName := make(map[string]string)

	for _, acc := range accounts {
		name := acc[0]
		u := acc[1]
		for i := 1; i < len(acc); i++ {
			v := acc[i]
			mailToName[v] = name
			adj[u] = append(adj[u], v)
			adj[v] = append(adj[v], u)
		}
	}

	seen := make(map[string]bool)

	var dfs func(string, string) []string
	dfs = func(source, neighbour string) []string {
		if _, ok := seen[neighbour]; ok {
			return []string{}
		}

		seen[neighbour] = true
		similar := []string{neighbour}
		for _, next := range adj[neighbour] {
			similar = append(similar, dfs(source, next)...)
		}
		return similar
	}

	result := [][]string{}

	for source := range adj {
		matchers := dfs(source, source)

		if len(matchers) == 0 {
			continue
		}

		current := []string{mailToName[source]}
		sort.Strings(matchers)
		current = append(current, matchers...)
		result = append(result, current)
	}

	return result
}

func DoAccountsMerge() {
	emails := [][]string{
		{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
		{"John", "johnsmith@mail.com", "john00@mail.com"},
		{"Mary", "mary@mail.com"},
		{"John", "johnnybravo@mail.com"},
	}
	fmt.Println(accountsMerge(emails))
}
