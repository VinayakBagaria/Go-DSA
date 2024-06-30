// https://leetcode.com/problems/remove-max-number-of-edges-to-keep-graph-fully-traversable
package graph

import (
	"fmt"
	"sort"
)

type dsu struct {
	parent     []int
	rank       []int
	components int
}

func initDsu(n int) *dsu {
	parent := make([]int, n+1)
	rank := make([]int, n+1)
	for i := range parent {
		parent[i] = i
		rank[i] = 1
	}

	return &dsu{
		parent:     parent,
		rank:       rank,
		components: n,
	}
}

func (d *dsu) find(u int) int {
	if d.parent[u] == u {
		return u
	}
	d.parent[u] = d.find(d.parent[u])
	return d.parent[u]
}

func (d *dsu) union(u, v int) {
	p1 := d.find(u)
	p2 := d.find(v)

	if p1 == p2 {
		return
	}

	if d.rank[p1] > d.rank[p2] {
		d.parent[p2] = p1
	} else if d.rank[p1] < d.rank[p2] {
		d.parent[p1] = p2
	} else {
		d.rank[p1]++
		d.parent[p2] = p1
	}

	d.components--
}

func (d *dsu) isSingle() bool {
	return d.components == 1
}

func maxNumEdgesToRemove(n int, edges [][]int) int {
	dsuAlice := initDsu(n)
	dsuBob := initDsu(n)

	sort.Slice(edges, func(i, j int) bool {
		return edges[i][0] > edges[j][0]
	})

	edgeCount := 0
	for _, edge := range edges {
		edgeType := edge[0]
		u := edge[1]
		v := edge[2]

		if edgeType == 3 {
			hasEdgeAdded := false

			if dsuAlice.find(u) != dsuAlice.find(v) {
				dsuAlice.union(u, v)
				hasEdgeAdded = true
			}

			if dsuBob.find(u) != dsuBob.find(v) {
				dsuBob.union(u, v)
				hasEdgeAdded = true
			}

			if hasEdgeAdded {
				edgeCount++
			}
		}

		if edgeType == 1 {
			if dsuAlice.find(u) != dsuAlice.find(v) {
				dsuAlice.union(u, v)
				edgeCount++
			}
		}

		if edgeType == 2 {
			if dsuBob.find(u) != dsuBob.find(v) {
				dsuBob.union(u, v)
				edgeCount++
			}
		}
	}

	if dsuAlice.isSingle() && dsuBob.isSingle() {
		return len(edges) - edgeCount
	}
	return -1
}

func DoMaxNumEdgesToRemove() {
	fmt.Println(maxNumEdgesToRemove(4, [][]int{
		{3, 1, 2},
		{3, 2, 3},
		{1, 1, 3},
		{1, 2, 4},
		{1, 1, 2},
		{2, 3, 4},
	}))
}
