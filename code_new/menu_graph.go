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
	g := &Graph{Vertices: map[int]*Vertex{}}

	return g
}

func (g *Graph) AddVertex(id int, text string) {
	g.Vertices[id] = &Vertex{Id: id, Text: text, Edges: map[int]*Edge{}}
}

func (g *Graph) AddEdge(vertexId int, destVertexId int) {
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

func (g *Graph) GetNeighbours(currVertex int) []int {
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
	g.AddVertex(1, "Main Menu")
	g.AddVertex(2, "Start New Game")
	g.AddVertex(3, "Resume Saved Game")
	g.AddVertex(4, "Exit")
	g.AddVertex(5, "Game Menu")
	g.AddVertex(6, "Choose Game")
	g.AddVertex(7, "Save Game")
	g.AddVertex(8, "Place Number")
	g.AddVertex(9, "Print Board")
	g.AddVertex(10, "Undo Move")
	
	g.AddEdge(15342,1)

	// add all menu connections
	g.AddEdge(1,2)
	g.AddEdge(1,3)
	g.AddEdge(1,4)
	//g.AddEdge(1,5) You should not be able to get from the main menu straight to the games menu

	g.AddEdge(2,5)

	g.AddEdge(3,4)
	g.AddEdge(3,6)

	// no need to connect Vertex 4 as it exits the programme

	g.AddEdge(5,1)
	g.AddEdge(5,4)
	g.AddEdge(5,7)
	g.AddEdge(5,8)
	g.AddEdge(5,9)
	g.AddEdge(5,10)

	g.AddEdge(6,5)

	return g
}
