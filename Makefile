include $(GOROOT)/src/Make.inc
TARG=github.com/humanfromearth/gopathfinding
GOFILES=\
	astar.go\
	dijkstra.go
include $(GOROOT)/src/Make.pkg
