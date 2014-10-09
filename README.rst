.. image:: https://secure.travis-ci.org/xarg/gopathfinding.png?branch=master

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

        $ go get github.com/xarg/gopathfinding

Using make
+++++++++++++

::

        $ git clone git://github.com/xarg/gopathfinding
        $ cd gopathfinding
        $ make install

Example
-----------

::

        import (
                "fmt"
                pathfinding "github.com/xarg/gopathfinding"
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

http://gopkgdoc.appspot.com/pkg/github.com/xarg/gopathfinding

Or

::

        $ go doc github.com/xarg/gopathfinding
