package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func startNewGame() (game){
	var playerName, difficulty string
	fmt.Println("Enter your name")
	fmt.Scan(&playerName)
	fmt.Println("Enter difficulty: easy, medium, hard")
	fmt.Scan(&difficulty)
	newGame := game {playerName : playerName, difficulty: difficulty, isComplete: false} 
	newGame.createBoard()
	return newGame
}

func (g *game) makeAMove() {
	var row, col, num int
	fmt.Println("Enter row")
	fmt.Scan(&row)
	fmt.Println("Enter column")
	fmt.Scan(&col)
	fmt.Println("Enter number")
	fmt.Scan(&num)
	g.placeNumber(row, col, num)
}


func (g *game) printBoard() {
    for i := 0; i < len(g.gameBoard); i++ {
        if i%3 == 0 && i != 0 {
            fmt.Println("---+---+---")
        }
        for j := 0; j < len(g.gameBoard); j++ {
            if j%3 == 0 && j != 0 {
                fmt.Print("|")
            }
            if g.gameBoard[i][j] == 0 {
                fmt.Print(" ")
            } else {
                fmt.Print(g.gameBoard[i][j])
            }
        }
        fmt.Println()
    }
}

func saveGame(g game) {
	savedGames = append(savedGames, g);
}

func loadSavedGame() (*game, error) {
	if len(savedGames) == 0 {
		return nil, errors.New("no games have been saved yet")
	}

	for i := 0; i < len(savedGames); i++ {
		fmt.Println(i, "name:", savedGames[i].playerName, "- difficulty:", savedGames[i].difficulty)
	}
	fmt.Println("Select a game to continue:")
	var userInput int
	fmt.Scan(&userInput)

	return &savedGames[userInput], nil
}

func (g *Graph) printMenuItems(currMenuItem int) {
	sb := strings.Builder{}

	childItems := g.GetNeighbours(currMenuItem)
	// Child items are returned in a random order. For the menu to work correctly, they need to be displayed in the correct order, so sorting here is necessary.
	sort.Ints(childItems)

	for i := 0; i < len(childItems); i++ {
		text := g.Vertices[childItems[i]].Text
		sb.WriteString(strconv.Itoa(i+1))
		sb.WriteString(". ")
		sb.WriteString(text)
		if i != len(childItems)-1 {
			sb.WriteString(" -- ")
		}
	}

	fmt.Println(sb.String())
}

