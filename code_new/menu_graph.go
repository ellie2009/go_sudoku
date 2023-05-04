package main

import "fmt"

type Graph struct {
	Vertices map[int]*Vertex
}

type Vertex struct {
	Id int
	Text string
	Edges map[int]*Edge
}

type Edge struct {
	Vertex *Vertex
}

func CreateNewGraph() *Graph {
	graph := &Graph{Vertices: map[int]*Vertex{}}

	return graph
}

func (g *Graph) AddNewVertex(id int, text string) {
	g.Vertices[id] = &Vertex{Id: id, Text: text, Edges: map[int]*Edge{}}
}

func (g *Graph) AddNewEdge(vertexId int, destVertexId int) {
	if g.Vertices[vertexId] == nil {
		fmt.Println("Node doesn't exist")
		return
	}

	if g.Vertices[destVertexId] == nil {
		fmt.Println("Destination node doesn't exist")
		return
	}

	g.Vertices[vertexId].Edges[destVertexId] = &Edge{Vertex: g.Vertices[destVertexId]}
} 

func (g *Graph) GetAllNeighbourElements(currVertex int) []int {
	neighbourList := []int{}

	// for each
	// _ indicates that the index is not being used
	for _, edge := range g.Vertices[currVertex].Edges {
		neighbourList = append(neighbourList, edge.Vertex.Id)
	}

	return neighbourList
}

func buildMenuGraph() *Graph {
	g := CreateNewGraph()

	// add all menu items
	g.AddNewVertex(1, "Main Menu")
	g.AddNewVertex(2, "Start New Game")
	g.AddNewVertex(3, "Resume Saved Game")
	g.AddNewVertex(4, "Exit")
	g.AddNewVertex(5, "Game Menu")
	g.AddNewVertex(6, "Choose Game")
	g.AddNewVertex(7, "Save Game")
	g.AddNewVertex(8, "Place Number")
	g.AddNewVertex(9, "Print Board")
	g.AddNewVertex(10, "Undo Move")
	
	// add all menu connections
	g.AddNewEdge(1,2)
	g.AddNewEdge(1,3)
	g.AddNewEdge(1,4)

	g.AddNewEdge(2,5)

	g.AddNewEdge(3,4)
	g.AddNewEdge(3,6)

	// no need to connect Vertex 4 as it exits the programme

	g.AddNewEdge(5,1)
	g.AddNewEdge(5,4)
	g.AddNewEdge(5,7)
	g.AddNewEdge(5,8)
	g.AddNewEdge(5,9)
	g.AddNewEdge(5,10)

	g.AddNewEdge(6,5)

	return g
}
