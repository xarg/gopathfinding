package pathfinding

import (
	"testing"
	"strings"
	"fmt"
)

const MAP1 = `............................
...e.........#..............
.............#..............
.............#..............
.............#..............
.............#..............
.............#..............
.............#..............
.............#..............
.......................s....
............................
`

func read_map (map_str string) MapDict {
	var result MapDict
	rows := strings.Split(map_str, "\n")

	for i:= 0; i <= len(rows); i++ {
		for j := 0; j <= len(rows[i]); j++ {
			char := rows[i][j]
			switch char{
			case '.':
				result.data[i][j] = LAND
			case '#':
				result.data[i][j] = WALL
			case 's':
				result.data[i][j] = START
			case 'e':
				result.data[i][j] = STOP
			}
		}
	}
	return result
}

func str_map(data MapDict, nodes []*Node) string {
	result := ""
	return result
}

func test_map1 (t *testing.T) {
	map_dict := read_map(MAP1)//Read map data and create a map_dict
	graph := NewGraph(map_dict) //Create a new graph
	nodes_path := Astar(graph) //Get the shortest path
	fmt.Print(str_map(map_dict, nodes_path))
}
