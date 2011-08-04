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
	rows := strings.Split(map_str, "\n")
	result := make(MapDict, len(rows))
	for i := 0; i <= len(rows); i++ {
		result[i] = make(map[int] int, len(rows))
	}
	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(rows[i]); j++ {
			char := rows[i][j]
			switch char{
			case '.':
				result[i][j] = LAND
			case '#':
				result[i][j] = WALL
			case 's':
				result[i][j] = START
			case 'e':
				result[i][j] = STOP
			}
		}
	}
	return result
}

func str_map(data MapDict, nodes []*Node) string {
	result := fmt.Sprintf("%d", len(nodes))
	for _, node := range nodes{
		fmt.Println(node)
	}
	return result
}

func TestMap1 (t *testing.T) {
	map_dict := read_map(MAP1)//Read map data and create a map_dict
	graph := NewGraph(map_dict) //Create a new graph
	nodes_path := Astar(graph) //Get the shortest path
	fmt.Print(str_map(map_dict, nodes_path))
}
