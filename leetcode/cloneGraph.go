// https://leetcode.com/problems/clone-graph/description/
package leetcode

func cloneGraph(node *Node) *Node {
	seen := make(map[int]*Node)

	var dfs func(*Node) *Node
	dfs = func(n *Node) *Node {
		if n == nil {
			return nil
		}

		if copied, ok := seen[n.Val]; ok {
			return copied
		}

		newNode := &Node{
			Val:       n.Val,
			Neighbors: []*Node{},
		}
		seen[n.Val] = newNode

		for _, each := range n.Neighbors {
			newNode.Neighbors = append(newNode.Neighbors, dfs(each))
		}

		return newNode
	}

	return dfs(node)
}

func DoCloneGraph() {

}
