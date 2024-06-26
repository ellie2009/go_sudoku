package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func startNewGame() (game){
	var playerName, difficultyInput string 
	difficulty := "easy"

	fmt.Println("Enter your name")
	fmt.Scan(&playerName)
	fmt.Println("Enter difficulty: easy, medium, hard")
	fmt.Scan(&difficultyInput)
	// Some very basic input validation. If input is incorrect or not medium/hard, an easy game is created by default 
	if difficultyInput == "medium" || difficultyInput == "hard" {
		difficulty = difficultyInput
	}
	newGame := game{id: uuid.New(), playerName : playerName, difficulty: difficulty, isComplete: false} 
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
            fmt.Println("---+---+---") // Println adds a new line to the string
        }
        for j := 0; j < len(g.gameBoard); j++ {
            if j%3 == 0 && j != 0 {
                fmt.Print("|") // Print doesn't add one, so has to be used here
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

func (g *game) saveGame() {
	savedGames[g.id] = *g
}

func loadSavedGame() (game, error) {
	if len(savedGames) == 0 {
		return game{}, errors.New("no games have been saved yet")
	}

	for _, game := range savedGames {
		fmt.Println(game.id, game.playerName, game.difficulty)
	}
	
	var userInput string
	fmt.Println("Select a game to continue by entering its UUID.")
	fmt.Scan(&userInput)

	gameId, err := uuid.Parse(userInput)

	if err != nil {
		fmt.Println("Incorrect game UUID entered.")
		return game{}, err 
	}

	game := savedGames[gameId]
	game.printBoard()

	return game, nil
}

func (g *Graph) printMenuItems(currMenuItem int) {
	sb := strings.Builder{}

	childItems := g.GetAllNeighbourElements(currMenuItem)
	// Child items are returned in a random order. For the menu to work correctly, 
	// they need to be displayed in the correct order, so sorting here is necessary.
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

