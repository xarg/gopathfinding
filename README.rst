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
                pathfinding "github.com/humanfromearth/gopathfinding"
        )

        func main() {
                shortest_path = pathfinding.Astar()
        }

Documentation
---------------

http://gopkgdoc.appspot.com/pkg/github.com/humanfromearth/gopathfinding

Or::
        $ godoc github.com/humanfromearth/gopathfinding
