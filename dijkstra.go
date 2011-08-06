package pathfinding

func minDist(dist map[*Node] int) *Node {
	var result_node *Node
	min := int(^uint(0) >> 1)
	for node, _dist := range dist {
		if min >= _dist {
			result_node = node
			min = _dist
		}
	}
	return result_node
}

func removeFromQ(Q []*Node, n *Node) []*Node {
	var result []*Node
	for _, node := range Q {
		if node != n {
			result = append(result, node)
		}
	}
	return result
}

func dist_between(n1 *Node, n2 *Node) int {
	return 0
}

func Dijkstra(graph *Graph) []*Node {
	MAX_INT := int(^uint(0) >> 1)
	var path []*Node
	var Q []*Node
	dist := make(map[*Node] int, len(graph.nodes))

	for _, node := range graph.nodes {
		dist[node] = MAX_INT
	}
	dist[graph.start] = 0
	copy(graph.nodes, Q)
	for len(Q) != 0 {
		u := minDist(dist)
		if dist[u] == MAX_INT {
			break
		}
		Q = removeFromQ(Q, u)
		for _, v := range graph.adjacentNodes(u) {
			alt := dist[u] + dist_between(u, v)
			if alt < dist[u] {
				dist[v] = alt
				v.parent = u
				//Reorder v in Q
			}
		}
	}
	return path
}
