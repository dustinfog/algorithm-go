package main

import "testing"

func TestShortestAlternatingPaths(t *testing.T) {
	n := 3
	redEdges := [][]int{{0, 1}, {1, 2}}
	blueEdges := [][]int{}

	ret := shortestAlternatingPaths(n, redEdges, blueEdges)
	t.Logf("%v", ret)

	redEdges = [][]int{{0, 1}}
	blueEdges = [][]int{{2, 1}}

	ret = shortestAlternatingPaths(n, redEdges, blueEdges)
	t.Logf("%v", ret)

	redEdges = [][]int{{1, 0}}
	blueEdges = [][]int{{2, 1}}

	ret = shortestAlternatingPaths(n, redEdges, blueEdges)
	t.Logf("%v", ret)

	redEdges = [][]int{{0, 1}}
	blueEdges = [][]int{{1, 2}}

	ret = shortestAlternatingPaths(n, redEdges, blueEdges)
	t.Logf("%v", ret)

	redEdges = [][]int{{0, 1}, {0, 2}}
	blueEdges = [][]int{{1, 0}}

	ret = shortestAlternatingPaths(n, redEdges, blueEdges)
	t.Logf("%v", ret)

	n = 5
	redEdges = [][]int{{3, 2}, {4, 1}, {1, 4}, {2, 4}}
	blueEdges = [][]int{{2, 3}, {0, 4}, {4, 3}, {4, 4}, {4, 0}, {1, 0}}
	//[0,2,-1,-1,1]
	ret = shortestAlternatingPaths(n, redEdges, blueEdges)
	t.Logf("%v", ret)
}
