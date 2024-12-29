package graphs

import "fmt"

type Graph struct {
	adjList map[string][]string
}

/*
	Adjacency List
-------------------------------------
	{
		"A" : ["B","C"]
		"B" : ["A","C","D"]
		"C" : ["A","B","D"]
		"D" : ["B","C","E"]
		"E" : ["D"]
	}

	The map keys are the vertices and the values in their arrays are the vertices they are connected to.
	The number of elements in the arrays gives you the number of edges connected to this vertex
	Adding a vertex is the same as adding a new key in the map
*/

func NewGraph() *Graph {
	return &Graph{
		adjList: make(map[string][]string),
	}
}

func (graph *Graph) AddVertex(vertex string) bool {

	// If the vertex already exist in the map it should not be added

	if _, exist := graph.adjList[vertex]; !exist {
		graph.adjList[vertex] = []string{}
		return true
	}
	return false
}

func (graph *Graph) AddEdge(vertex1, vertex2 string) bool {

	list1, exist1 := graph.adjList[vertex1]
	list2, exist2 := graph.adjList[vertex2]

	if !exist1 || !exist2 {
		return false
	}

	list1 = append(list1, vertex2)

	graph.adjList[vertex1] = list1

	list2 = append(list2, vertex1)

	graph.adjList[vertex2] = list2

	return true

}

func (graph *Graph) RemoveEdge(vertex1, vertex2 string) {

	list1, exist1 := graph.adjList[vertex1]
	list2, exist2 := graph.adjList[vertex2]

	if !exist1 || !exist2 {
		return
	}

	for i, vertex := range list1 {
		if vertex == vertex2 {
			list1 = append(list1[:i], list1[i+1:]...)
			break
		}
	}

	for i, vertex := range list2 {
		if vertex == vertex1 {
			list2 = append(list2[:i], list2[i+1:]...)
			break
		}
	}

	fmt.Printf("%v", graph.adjList)
}

func (graph *Graph) RemoveVertex(vertex string) {

	fmt.Printf("BEFORE DELETION %v \n", graph.adjList)

	if otherVertices, exist := graph.adjList[vertex]; exist {
		for _, otherVertex := range otherVertices {

			if otherVertexList, exist := graph.adjList[otherVertex]; exist {
				for i, v := range otherVertexList {
					if v == vertex {
						otherVertexList = append(otherVertexList[:i], otherVertexList[i+1:]...)
						graph.adjList[otherVertex] = otherVertexList
						break
					}
				}
			}
		}

	}

	delete(graph.adjList, vertex)

	fmt.Printf("AFTER DELETION %v \n", graph.adjList)
}
