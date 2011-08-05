gopathfinding
=================

A small package that implements pathfinding algorithms.

Implementions
---------------

 * A* (A star) - implemented
 * Dijkstra - not implemented

Installing
------------

Using goinstall
++++++++++++++++++

::

        $ goinstall github.com/humanfromearth/gopathfinding

Using make
+++++++++++++

::

        $ git clone git://github.com/humanfromearth/gopathfinding
        $ cd gopathfinding
        $ make install

Example
-----------

::

        import (
                "fmt"
                pathfinding "github.com/humanfromearth/gopathfinding"
        )

        func main() {
                //A pathfinding.MapData containing the 
                //coordinates(x, y) of LAND, WALL, START and STOP of the map.
                //If your map is something more than 2d matrix then you might want to modify adjacentNodes

                graph := pathfinding.NewGraph(map_data)

                //Returns a list of nodes from START to STOP avoiding all obstacles if possible
                shortest_path := pathfinding.Astar(graph)
        }

Documentation
---------------

http://gopkgdoc.appspot.com/pkg/github.com/humanfromearth/gopathfinding

Or

::

        $ godoc github.com/humanfromearth/gopathfinding
