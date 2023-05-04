package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

var savedGames map[uuid.UUID]game
var currGame game;
var menu *Graph

func main() {
    rand.Seed(time.Now().UnixNano()) // use this to be able to create random numbers

	menuLevel := 1
	// Menu levels correspond to the Menu Graph that's build below
	// 1 - Main Menu
	// 5 - Game Menu
	menu = buildMenuGraph()
	savedGames = make(map[uuid.UUID]game)

	for menuLevel != 0 {
		switch menuLevel {
		case 1: // Main Menu
			fmt.Println(strings.ToUpper(menu.Vertices[1].Text)) 
			menu.printMenuItems(1)
		case 5: // Game Menu
			fmt.Println(strings.ToUpper(menu.Vertices[5].Text))
			menu.printMenuItems(5)
		}


		fmt.Print("Enter selection: ")
		var userInput int
		fmt.Scan(&userInput)

		switch menuLevel {
			case 1:
				if userInput == 1 { // Start New Game
					currGame = startNewGame()
					currGame.printBoard()
					menuLevel = 5
				} else if userInput == 2 { // Resume Saved Game
					savedGame, err := loadSavedGame()
					if err != nil {
						fmt.Println(err)
					} else {
					currGame = savedGame
					menuLevel = 5
					} 
				} else if userInput == 3 { // Exit Programme
					menuLevel = 0
				} else {
					fmt.Println("Invalid selection.")
				}
			case 5:
				if userInput == 1 { // Return to Main Menu
					menuLevel = 1
				} else if userInput == 2 { // Exit Programme
					fmt.Println("Exiting...")
					menuLevel = 0
				} else if userInput == 3 { // Save Game
					currGame.saveGame()
					fmt.Println("Game saved")
				} else if userInput == 4 { // Place Number
					currGame.makeAMove()
					fmt.Println("===================")
					currGame.printBoard()
					fmt.Println("===================")
					if currGame.isGameComplete() {
						fmt.Println("You have won!\n Returning to Main Menu")
						currGame.isComplete = true
						menuLevel = 1
					}
				} else if userInput == 5 { // Print Board
					currGame.printBoard()
				} else if userInput == 6 { // Undo Last Move
					currGame.undoLastMove()
					fmt.Println("===================")
					currGame.printBoard()
					fmt.Println("===================")
				} else {
					fmt.Println("Invalid selection.")
				}

		}
	}
}
