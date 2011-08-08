//pathfinding package implements pathfinding algorithms such as Dijkstra and A*
package pathfinding

import (
	"fmt"
)

//Defining possible graph elements
const (
	UNKNOWN int = iota - 1
	LAND
	WALL
	START
	STOP
)

type MapData [][]int

//Return a new MapData by value given some dimensions
func NewMapData(rows, cols int) *MapData {
	result := make(MapData, rows)
	for i := 0; i < rows; i++ {
		result[i] = make([]int, cols)
	}
	return &result
}

//A point is just a set of x, y coordinates with a PointType attached
type Node struct {
	x, y   int //Using int for efficiency
	parent *Node
	H      int
}

//Create a new node
func NewNode(x, y int) *Node {
	max_int := int(^uint(0) >> 1)
	node := &Node{
		x:      x,
		y:      y,
		parent: nil,
		H:      max_int,
	}
	return node
}

//Return string representation of the node
func (self *Node) String() string {
	return fmt.Sprintf("<Node x:%d y:%d addr:%d>", self.x, self.y, &self)
}

//Start, end nodes and a slice of nodes
type Graph struct {
	start, stop *Node
	nodes       []*Node
	data        *MapData
}

//Return a Graph from a map of coordinates (those that are passible)
func NewGraph(map_data *MapData) *Graph {
	var start, stop *Node
	var nodes []*Node
	for i, row := range *map_data {
		for j, _type := range row {
			if _type == START || _type == STOP {
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
		stop:  stop,
		data:  map_data,
	}
	return g
}

//Get the nodes near some node
func (self *Graph) adjacentNodes(node *Node) []*Node {
	var result []*Node
	map_data := *self.data
	rows := len(map_data)
	cols := len(map_data[0])

	//If the coordinates are passable then create a new node and add it
	append_node := func (x, y int) {
		//Check if node is not already in the graph and append that node
		for _, n := range self.nodes {
			if n.x == x && n.y == y {
				result = append(result, n)
				return
			}
		}
		_type := map_data[x][y]
		if _type == LAND || _type == STOP {
			//Create a new node and add it to the graph
			n := NewNode(x, y)
			result = append(result, n)
			self.nodes = append(self.nodes, n)
		}
	}
	if node.x <= rows && node.y + 1 < cols {
		append_node(node.x, node.y + 1)
	}
	if node.x <= rows && node.y - 1 >= 0 {
		append_node(node.x, node.y - 1)
	}
	if node.y <= cols && node.x + 1 < rows {
		append_node(node.x + 1, node.y)
	}
	if node.y <= cols && node.x - 1 >= 0 {
		append_node(node.x - 1, node.y)
	}
	return result
}


func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func removeNode(nodes []*Node, node *Node) []*Node {
	ith := -1
	for i, n := range nodes {
		if n == node {
			ith = i
			break
		}
	}
	if ith != -1 {
		copy(nodes[ith:], nodes[ith+1:])
		nodes = nodes[:len(nodes)-1]
	}
	return nodes
}

func hasNode(nodes []*Node, node *Node) bool {
	for _, n := range nodes {
		if n == node {
			return true
		}
	}
	return false
}

//Return the node with the minimum H
func minH(nodes []*Node) *Node {
	var result_node *Node
	minH := int(^uint(0) >> 1)
	for _, node := range nodes {
		if node.H <= minH {
			minH = node.H
			result_node = node
		}
	}
	return result_node
}

func retracePath(current_node *Node) []*Node {
	var path []*Node
	path = append(path, current_node)
	for current_node.parent != nil {
		path = append(path, current_node.parent)
		current_node = current_node.parent
	}
	//Reverse path
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}

//A* search algorithm. See http://en.wikipedia.org/wiki/A*_search_algorithm
func Astar(graph *Graph) []*Node {
	var path, openSet, closedSet []*Node

	openSet = append(openSet, graph.start)
	for len(openSet) != 0 {
		//Get the node with the min H
		current := minH(openSet)
		if current == graph.stop {
			return retracePath(current)
		}
		openSet = removeNode(openSet, current)
		closedSet = append(closedSet, current)
		for _, tile := range graph.adjacentNodes(current) {
			if tile != nil && graph.stop != nil && !hasNode(closedSet, tile) {
				tile.H = (abs(graph.stop.x-tile.x) +
					abs(graph.stop.y-tile.y)) * 10
				if !hasNode(openSet, tile) {
					openSet = append(openSet, tile)
				}
				tile.parent = current
			}
		}
	}
	return path
}
