package main

type out struct {
	red  []int
	blue []int
}

type node struct {
	red   bool
	index int
	len   int
}

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	ret := make([]int, n)
	ret[0] = 0
	for i := 1; i < n; i++ {
		ret[i] = -1
	}

	var edges [100]out
	for _, edge := range redEdges {
		edges[edge[0]].red = append(edges[edge[0]].red, edge[1])
	}

	for _, edge := range blueEdges {
		edges[edge[0]].blue = append(edges[edge[0]].blue, edge[1])
	}

	queue := make([]*node, 0, len(edges[0].red)+len(edges[0].blue))
	scatterSeek(&queue, &edges[0].red, ret, 1, true)
	scatterSeek(&queue, &edges[0].blue, ret, 1, false)

	for len(queue) != 0 {
		nd := queue[0]
		queue = queue[1:]

		if nd.red {
			scatterSeek(&queue, &edges[nd.index].blue, ret, nd.len+1, false)
		} else {
			scatterSeek(&queue, &edges[nd.index].red, ret, nd.len+1, true)
		}
	}

	return ret
}

func scatterSeek(queue *[]*node, neighbors *[]int, ret []int, l int, red bool) {
	for _, next := range *neighbors {
		if ret[next] == -1 {
			ret[next] = l
		}

		*queue = append(*queue, &node{
			red:   red,
			index: next,
			len:   l,
		})
	}

	*neighbors = nil
}

//https://leetcode-cn.com/problems/shortest-path-with-alternating-colors/
