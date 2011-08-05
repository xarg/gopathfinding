package pathfinding

import (
	"testing"
	"strings"
	"fmt"
)

const MAP1 = `............................
............................
.............#..............
.............#..............
.......e.....#..............
.............#..............
.............#..............
.............#..............
.............#..............
.............#.........s....
............................`

const MAP2 = `............................
............................
.............#..............
.............#..............
.......e.....#..............
.............#..............
.............#..............
.............#..............
.............#.......#######
.............#.......#.s....
.....................#......`

func read_map (map_str string) MapData {
	rows := strings.Split(map_str, "\n")
	result := make(MapData, len(rows))
	for i := 0; i <= len(rows) - 1; i++ {
		result[i] = make([]int, len(rows[1]))
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

func str_map(data MapData, nodes []*Node) string {
	var result string
	for i, row := range data {
		for j, cell := range row {
			added := false
			for _, node := range nodes {
				if node.x == i && node.y == j {
					result += "+"
					added = true
					break
				}
			}
			if !added {
				switch cell {
				case LAND:
					result += "."
				case WALL:
					result += "#"
				case START:
					result += "s"
				case STOP:
					result += "e"
				default: //Unknown
					result += "?"
				}
			}
		}
		result += "\n"
	}
	return result
}

func TestAstar1 (t *testing.T) {
	map_data := read_map(MAP1)//Read map data and create the map_data
	graph := NewGraph(map_data) //Create a new graph
	nodes_path := Astar(graph) //Get the shortest path
	fmt.Println(str_map(map_data, nodes_path))
	if len(nodes_path) != 28 {
		t.Errorf("Expected 28. Got %d", len(nodes_path))
	}
}

func TestAstar2 (t *testing.T) {
	map_data := read_map(MAP2)//Read map data and create the map_data
	graph := NewGraph(map_data) //Create a new graph
	nodes_path := Astar(graph) //Get the shortest path
	fmt.Println(str_map(map_data, nodes_path))
	if len(nodes_path) != 0 {
		t.Errorf("Expected 0. Got %d", len(nodes_path))
	}
}
