//pathfinding package implements pathfinding algorithms such as Dijkstra and A*
package pathfinding

import (
	"fmt"
)

//Defining possible graph elements
const (
	MAX_INT int(^uint(0) >> 1)
	UNKNOWN int = iota - 1
	LAND
	WALL
	START
	STOP
)

type MapDict struct {
	data map[int] map[int] int
}

//A point is just a set of x, y coordinates with a PointType attached
type Node struct {
	x, y int //Using int for efficiency
	parent *Node
	H int
}

//Create a new node
func NewNode(x, y int) *Node {
	max_int := MAX_INT
	node := &Node{
		x: x,
		y: y,
		parent: new(Node),
		H: max_int,
	}
	return node
}

//Return string representation of the node
func (self *Node) String() string {
	return fmt.Sprintf("<Node x:%s y:%s addr:%s>", self.x, self.y, &self)
}

//Start, end nodes and a slice of nodes
type Graph struct {
	start, stop *Node
	nodes []*Node
}

//Return a Graph from a map of coordinates (those that are passible)
func NewGraph (map_data MapDict) *Graph {
	var start, stop *Node
	nodes := make([]*Node, len(map_data.data) + len(map_data.data[0]))
	for i, row := range map_data.data {
		for j, _type := range row {
			if _type == LAND || _type == START || _type == STOP {
				node := NewNode(i, j)
				nodes = append(nodes, node)
				if _type == START {
					start = node
				}
				if _type == STOP {
					stop = node
				}
			}
		}
	}
	g := &Graph{
		nodes: nodes,
		start: start,
		stop: stop,
	}
	return g
}

//Get the nodes near some node
func (self *Graph) GetAdjacentNodes(node *Node) []*Node {
	return self.nodes
}

func retracePath(c) []*Node {
	path.insert(0, c)
		if c.parent == nil {
			return
		}
		retracePath(c.parent)
	}
}

//Return the node with the minimum H
func minCost(nodes []*Node) *Node {
	var result_node *Node
	minH := MAX_INT
	for node := range nodes{
		if node.H <= minH {
			minH = node.H
			result_node = node
		}
	}
	return result_node
}

//A* search algorithm. See http://en.wikipedia.org/wiki/A*_search_algorithm
func Astar(graph *Graph) []*Node {
	var path, openSet, closedSet []*Node
	openSet = append(openSet, graph.start)
	for len(openSet) != 0 {
		//Get the minimum cost node
		current := minCost(openSet)
		if current == graph.end {
			retracePath(current)
			break
		}
		openSet.remove(current)
		closedSet.add(current)
		for tile in graph[current]{
			if tile not in closedSet {
				tile.H = (abs(graph.end.x-tile.x)+abs(graph.end.y-tile.y))*10
				if tile not in openSet {
					openSet.add(tile)
				}
				tile.parent = current
			}
		}
	}
	return path
}
