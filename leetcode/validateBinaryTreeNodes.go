// https://leetcode.com/problems/validate-binary-tree-nodes/description/
package leetcode

func hasTraversedAll(indegree []int, adjList map[int][]int, queue []int, n int) bool {
	count := 0
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		count++

		for _, v := range adjList[u] {
			indegree[v]--
			if indegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	return count == n
}

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	adjList := make(map[int][]int)
	indegree := make([]int, n)
	for i := 0; i < n; i++ {
		l := leftChild[i]
		r := rightChild[i]

		if l != -1 {
			indegree[l]++
			adjList[i] = append(adjList[i], l)
		}

		if r != -1 {
			indegree[r]++
			adjList[i] = append(adjList[i], r)
		}
	}

	queue := []int{}
	for u, indeg := range indegree {
		if indeg == 0 {
			// If length is 1, no disconnected tree
			// - as multiple nodes won't have indegree 0 for valid tree
			if len(queue) > 0 {
				return false
			}
			queue = append(queue, u)
		} else if indeg > 1 {
			// If any node has more than 1 parent, then it can't be a binary tree
			return false
		}
	}

	return hasTraversedAll(indegree, adjList, queue, n)
}
