package main

import "fmt"

// Question: Breadth-First Search (BFS) Traversal of a Graph.
// bfsTraversal visits every component of an undirected graph.
func bfsTraversal(vertices int, adjacencyList [][]int) []int {
	visited := make([]bool, vertices)
	traversal := []int{}

	for start := 0; start < vertices; start++ {
		if visited[start] {
			continue
		}

		visited[start] = true
		queue := []int{start}
		for len(queue) > 0 {
			vertex := queue[0]
			queue = queue[1:]
			traversal = append(traversal, vertex)

			for _, neighbor := range adjacencyList[vertex] {
				if !visited[neighbor] {
					visited[neighbor] = true
					queue = append(queue, neighbor)
				}
			}
		}
	}

	return traversal
}

// Question: Depth-First Search (DFS) Traversal of a Graph.
// dfsTraversal visits every component of an undirected graph recursively.
func dfsTraversal(vertices int, adjacencyList [][]int) []int {
	visited := make([]bool, vertices)
	traversal := []int{}

	var dfs func(int)
	dfs = func(vertex int) {
		visited[vertex] = true
		traversal = append(traversal, vertex)
		for _, neighbor := range adjacencyList[vertex] {
			if !visited[neighbor] {
				dfs(neighbor)
			}
		}
	}

	for vertex := 0; vertex < vertices; vertex++ {
		if !visited[vertex] {
			dfs(vertex)
		}
	}

	return traversal
}

// Question: Number of Provinces.
// findCircleNum counts connected components in an adjacency matrix.
func findCircleNum(isConnected [][]int) int {
	visited := make([]bool, len(isConnected))
	provinces := 0

	var dfs func(int)
	dfs = func(city int) {
		visited[city] = true
		for neighbor, connected := range isConnected[city] {
			if connected == 1 && !visited[neighbor] {
				dfs(neighbor)
			}
		}
	}

	for city := range isConnected {
		if !visited[city] {
			provinces++
			dfs(city)
		}
	}

	return provinces
}

// Question: Flood Fill.
func floodFill(image [][]int, startRow, startCol, color int) [][]int {
	if startRow < 0 || startRow >= len(image) || startCol < 0 || startCol >= len(image[startRow]) {
		return image
	}

	originalColor := image[startRow][startCol]
	if originalColor == color {
		return image
	}

	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	queue := [][2]int{{startRow, startCol}}
	image[startRow][startCol] = color

	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]
		for _, direction := range directions {
			nextRow, nextCol := cell[0]+direction[0], cell[1]+direction[1]
			if nextRow < 0 || nextRow >= len(image) || nextCol < 0 || nextCol >= len(image[nextRow]) || image[nextRow][nextCol] != originalColor {
				continue
			}
			image[nextRow][nextCol] = color
			queue = append(queue, [2]int{nextRow, nextCol})
		}
	}

	return image
}

// GraphNode is an undirected graph node used by cloneGraph.
type GraphNode struct {
	Val       int
	Neighbors []*GraphNode
}

// Question: Clone Graph.
func cloneGraph(node *GraphNode) *GraphNode {
	if node == nil {
		return nil
	}

	clones := map[*GraphNode]*GraphNode{}
	var clone func(*GraphNode) *GraphNode
	clone = func(current *GraphNode) *GraphNode {
		if copied, exists := clones[current]; exists {
			return copied
		}

		copied := &GraphNode{Val: current.Val}
		clones[current] = copied
		for _, neighbor := range current.Neighbors {
			copied.Neighbors = append(copied.Neighbors, clone(neighbor))
		}
		return copied
	}

	return clone(node)
}

// Question: Course Schedule.
// canFinish uses Kahn's topological-sort algorithm to detect a cycle.
func canFinish(numCourses int, prerequisites [][]int) bool {
	adjacencyList := make([][]int, numCourses)
	inDegree := make([]int, numCourses)
	for _, prerequisite := range prerequisites {
		course, requiredCourse := prerequisite[0], prerequisite[1]
		adjacencyList[requiredCourse] = append(adjacencyList[requiredCourse], course)
		inDegree[course]++
	}

	queue := []int{}
	for course, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, course)
		}
	}

	completed := 0
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		completed++
		for _, nextCourse := range adjacencyList[course] {
			inDegree[nextCourse]--
			if inDegree[nextCourse] == 0 {
				queue = append(queue, nextCourse)
			}
		}
	}

	return completed == numCourses
}

// Question: Is Graph Bipartite?
func isBipartite(graph [][]int) bool {
	colors := make([]int, len(graph)) // 0 = uncolored, 1 and -1 are the two sets.

	for start := range graph {
		if colors[start] != 0 {
			continue
		}

		colors[start] = 1
		queue := []int{start}
		for len(queue) > 0 {
			vertex := queue[0]
			queue = queue[1:]
			for _, neighbor := range graph[vertex] {
				if colors[neighbor] == colors[vertex] {
					return false
				}
				if colors[neighbor] == 0 {
					colors[neighbor] = -colors[vertex]
					queue = append(queue, neighbor)
				}
			}
		}
	}

	return true
}

// Question: Dijkstra's Shortest Path Algorithm.
// dijkstra returns shortest distances from source for non-negative weighted edges.
// Each edge is [from, to, weight]. Unreachable vertices remain -1.
func dijkstra(vertices int, edges [][]int, source int) []int {
	adjacencyList := make([][][2]int, vertices)
	for _, edge := range edges {
		from, to, weight := edge[0], edge[1], edge[2]
		adjacencyList[from] = append(adjacencyList[from], [2]int{to, weight})
	}

	distances := make([]int, vertices)
	visited := make([]bool, vertices)
	for vertex := range distances {
		distances[vertex] = -1
	}
	distances[source] = 0

	for range vertices {
		current := -1
		for vertex := range vertices {
			if !visited[vertex] && distances[vertex] != -1 && (current == -1 || distances[vertex] < distances[current]) {
				current = vertex
			}
		}
		if current == -1 {
			break
		}

		visited[current] = true
		for _, edge := range adjacencyList[current] {
			next, weight := edge[0], edge[1]
			candidate := distances[current] + weight
			if distances[next] == -1 || candidate < distances[next] {
				distances[next] = candidate
			}
		}
	}

	return distances
}

func main() {
	graph := [][]int{{1, 2}, {0, 3}, {0}, {1}, {5}, {4}}

	// Question: Breadth-First Search (BFS) Traversal of a Graph.
	fmt.Println("BFS traversal:", bfsTraversal(6, graph))

	// Question: Depth-First Search (DFS) Traversal of a Graph.
	fmt.Println("DFS traversal:", dfsTraversal(6, graph))

	// Question: Number of Provinces.
	fmt.Println("Provinces:", findCircleNum([][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}))

	// Question: Flood Fill.
	fmt.Println("Flood fill:", floodFill([][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2))

	// Question: Clone Graph.
	first := &GraphNode{Val: 1}
	second := &GraphNode{Val: 2}
	first.Neighbors = []*GraphNode{second}
	second.Neighbors = []*GraphNode{first}
	clone := cloneGraph(first)
	fmt.Println("Cloned graph root:", clone.Val, "neighbor:", clone.Neighbors[0].Val)

	// Question: Course Schedule.
	fmt.Println("Can finish courses:", canFinish(2, [][]int{{1, 0}}))

	// Question: Is Graph Bipartite?
	fmt.Println("Is bipartite:", isBipartite([][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}))

	// Question: Dijkstra's Shortest Path Algorithm.
	fmt.Println("Shortest distances:", dijkstra(5, [][]int{{0, 1, 4}, {0, 2, 1}, {2, 1, 2}, {1, 3, 1}, {2, 3, 5}}, 0))
}
