package graphs

import (
	"fmt"
	"testing"
)

/*
	{
		"A" : ["B","C"]
		"B" : ["A","C","D"]
		"C" : ["A","B","D"]
		"D" : ["B","C","E"]
		"E" : ["D"]
	}

	The map keys are the vertices and the values in their arrays are the vertices they are connected to.
	The number of elements in the arrays gives you the number of edges connected to this vertex
*/

func TestAddVertex(t *testing.T) {

	graph := NewGraph()

	if !graph.AddVertex("A") {
		t.Errorf("adding vertex to graph failed")
	}

}

func TestRemoveEdge(t *testing.T) {

	graph := NewGraph()

	graph.AddVertex("A")
	graph.AddVertex("B")
	graph.AddVertex("C")
	graph.AddVertex("D")
	graph.AddVertex("E")

	graph.AddEdge("A", "B")
	graph.AddEdge("A", "C")
	graph.AddEdge("B", "C")
	graph.AddEdge("B", "D")
	graph.AddEdge("C", "D")
	graph.AddEdge("D", "E")

	vertex1, vertex2 := "A", "B"

	graph.RemoveEdge(vertex1, vertex2)

	list1, _ := graph.adjList[vertex1]
	list2, _ := graph.adjList[vertex2]

	for _, v := range list1 {
		if v == vertex2 {
			t.Errorf("%s <----> %s should not exist in %v", vertex1, vertex2, list1)
			break
		}
	}

	for _, v := range list2 {
		if v == vertex1 {
			t.Errorf("%s <----> %s should not exist in %v", vertex1, vertex2, list2)
			break
		}
	}
}

func TestRemoveVertex(t *testing.T) {
	graph := NewGraph()

	// Add vertices
	vertices := []string{"A", "B", "C", "D", "E"}
	for _, v := range vertices {
		graph.AddVertex(v)
	}

	// Add edges
	edges := [][2]string{
		{"A", "B"},
		{"A", "C"},
		{"B", "C"},
		{"B", "D"},
		{"C", "D"},
		{"D", "E"},
	}
	for _, edge := range edges {
		graph.AddEdge(edge[0], edge[1])
	}

	vertex := "A"

	connectedVertices, exists := graph.adjList[vertex]

	if !exists {
		t.Fatalf("Vertex %s does not exist in the graph", vertex)
	}

	graph.RemoveVertex(vertex)

	if _, exists := graph.adjList[vertex]; exists {
		t.Errorf("Vertex %s still exists in the graph after removal", vertex)
	}

	for _, v := range connectedVertices {
		othersList, exists := graph.adjList[v]
		if !exists {
			t.Errorf("Vertex %s should still exist, but it was removed", v)
			continue
		}
		for _, neighbor := range othersList {
			if neighbor == vertex {
				t.Errorf("Edge %s <-> %s still exists after vertex %s was deleted", vertex, v, vertex)
			}
		}
	}

	// Optionally: Add subtests for each connected vertex to improve readability
	for _, v := range connectedVertices {
		t.Run(fmt.Sprintf("Verify no edges remain to %s", v), func(t *testing.T) {
			othersList, exists := graph.adjList[v]
			if !exists {
				t.Errorf("Vertex %s should still exist, but it was removed", v)
				return
			}
			for _, neighbor := range othersList {
				if neighbor == vertex {
					t.Errorf("Edge %s <-> %s still exists after vertex %s was deleted", vertex, v, vertex)
				}
			}
		})
	}
}
