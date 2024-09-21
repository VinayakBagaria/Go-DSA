// https://leetcode.com/problems/evaluate-division/description/
package leetcode

import "fmt"

type Element struct {
	v    string
	cost float64
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	adj := make(map[string][]Element)
	for i, eq := range equations {
		u := eq[0]
		v := eq[1]
		adj[u] = append(adj[u], Element{v, values[i]})
		adj[v] = append(adj[v], Element{u, 1.0 / values[i]})
	}

	result := make([]float64, len(queries))
	for i, q := range queries {
		u := q[0]
		v := q[1]
		_, uOk := adj[u]
		_, vOk := adj[v]
		if !uOk || !vOk {
			result[i] = -1.0
			continue
		}
		result[i] = startDfs(u, v, adj)
	}

	return result
}

func startDfs(start, end string, adj map[string][]Element) float64 {
	seen := make(map[string]struct{})

	var dfs func(string) float64
	dfs = func(u string) float64 {
		if _, ok := seen[u]; ok {
			return -1.0
		}
		if u == end {
			return 1.0
		}

		seen[u] = struct{}{}

		ans := -1.0
		for _, next := range adj[u] {
			d := dfs(next.v)
			if d != -1.0 {
				ans = maxFloat(ans, next.cost*d)
			}
		}

		delete(seen, u)
		return ans
	}

	return dfs(start)
}

func maxFloat(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func DoEvaluateDivision() {
	fmt.Println(calcEquation([][]string{
		{"a", "b"},
		{"b", "c"},
	}, []float64{2.0, 3.0}, [][]string{
		{"a", "c"},
		{"b", "a"},
		{"a", "e"},
		{"a", "a"},
		{"x", "x"},
	}))
}
