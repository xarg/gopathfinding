package pathfinding

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

const MAP0 = `s....
.....
##.##
.....
....e`

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

const MAP3 = `............................
............................
............................
............................
............................
.........#..................
.......s#...................
.........#..................
..........##................
............##...e...##.....
..............##...##.......
................###.........
............................`

func read_map(map_str string) *MapData {
	rows := strings.Split(map_str, "\n")
	if len(rows) == 0 {
		panic("The map needs to have at least 1 row")
	}
	row_count := len(rows)
	col_count := len(rows[0])

	result := *NewMapData(row_count, col_count)
	for i := 0; i < row_count; i++ {
		for j := 0; j < col_count; j++ {
			char := rows[i][j]
			switch char {
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
	return &result
}

func str_map(data *MapData, nodes []*Node) string {
	var result string
	for i, row := range *data {
		for j, cell := range row {
			added := false
			for _, node := range nodes {
				if node.X == i && node.Y == j {
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

//Generate a random MapData given some dimensions
func generate_map(n int) *MapData {
	map_data := *NewMapData(n, n)
	map_data[0][0] = START
	map_data[n-1][n-1] = STOP

	xs := rand.Perm(n - 1)
	ys := rand.Perm(n - 1)

	for i := 1; i < len(xs); i += rand.Intn(4) + 1 {
		for j := 1; j < len(ys); j++ {
			map_data[xs[i]][ys[j]] = WALL
		}
	}
	return &map_data
}

func TestAstar0(t *testing.T) {
	map_data := read_map(MAP0)  //Read map data and create the map_data
	graph := NewGraph(map_data) //Create a new graph
	nodes_path := Astar(graph)  //Get the shortest path
	fmt.Println(str_map(map_data, nodes_path))
	if len(nodes_path) != 9 {
		t.Errorf("Expected 9. Got %d", len(nodes_path))
	}
}

func TestAstar1(t *testing.T) {
	map_data := read_map(MAP1)  //Read map data and create the map_data
	graph := NewGraph(map_data) //Create a new graph
	nodes_path := Astar(graph)  //Get the shortest path
	fmt.Println(str_map(map_data, nodes_path))
	if len(nodes_path) != 24 {
		t.Errorf("Expected 24. Got %d", len(nodes_path))
	}
}

func TestAstar2(t *testing.T) {
	map_data := read_map(MAP2)  //Read map data and create the map_data
	graph := NewGraph(map_data) //Create a new graph
	nodes_path := Astar(graph)  //Get the shortest path
	fmt.Println(str_map(map_data, nodes_path))
	if len(nodes_path) != 0 {
		t.Errorf("Expected 0. Got %d", len(nodes_path))
	}
}

func TestAstar3(t *testing.T) {
	map_data := read_map(MAP3)  //Read map data and create the map_data
	graph := NewGraph(map_data) //Create a new graph
	nodes_path := Astar(graph)  //Get the shortest path
	fmt.Println(str_map(map_data, nodes_path))
	if len(nodes_path) != 18 {
		t.Errorf("Expected 18. Got %d", len(nodes_path))
	}
}
func BenchmarkAstar1(b *testing.B) {
	b.StopTimer()
	fmt.Printf("Benchmarking with a %dx%d map\n", b.N, b.N)
	map_data := generate_map(b.N + 5)
	graph := NewGraph(map_data)
	b.StartTimer()
	Astar(graph)
}
