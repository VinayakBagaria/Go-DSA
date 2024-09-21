// https://leetcode.com/problems/path-with-maximum-probability/description/
package leetcode

import (
	"fmt"
	"math"
)

type adjElement struct {
	to       int
	succProb float64
}

func makeAdjacencyList(edges [][]int, succProb []float64, adjList map[int][]adjElement) {
	for i, edge := range edges {
		adjList[edge[0]] = append(adjList[edge[0]], adjElement{
			to:       edge[1],
			succProb: succProb[i],
		})
		adjList[edge[1]] = append(adjList[edge[1]], adjElement{
			to:       edge[0],
			succProb: succProb[i],
		})
	}
}

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	adjList := make(map[int][]adjElement)
	makeAdjacencyList(edges, succProb, adjList)

	queue := adjList[start]
	// Store max probability of reaching each node from starting node
	distances := make([]float64, n)

	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			element := queue[i]

			distances[element.to] = math.Max(distances[element.to], element.succProb)

			for _, nextElement := range adjList[element.to] {
				// Calculate new probability
				// If its higher, use this neighbour to go over its neighbour once again
				newSuccessProb := element.succProb * nextElement.succProb
				if newSuccessProb > distances[nextElement.to] {
					queue = append(queue, adjElement{
						to:       nextElement.to,
						succProb: newSuccessProb,
					})
				}
			}
		}

		queue = queue[size:]
	}

	return distances[end]
}

func DoPathWithMaximumProbability() {
	edges1 := [][]int{{0, 1}, {1, 2}, {0, 2}}
	fmt.Println(maxProbability(3, edges1, []float64{0.5, 0.5, 0.2}, 0, 2))
	fmt.Println(maxProbability(3, edges1, []float64{0.5, 0.5, 0.2}, 0, 2))
	edges2 := [][]int{{2, 3}, {1, 2}, {3, 4}, {1, 3}, {1, 4}, {0, 1}, {2, 4}, {0, 4}, {0, 2}}
	fmt.Println(maxProbability(5, edges2, []float64{0.06, 0.26, 0.49, 0.25, 0.2, 0.64, 0.23, 0.21, 0.77}, 0, 3))
}
