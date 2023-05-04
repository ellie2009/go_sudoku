package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

//var savedGames []game; // a slice of games where you can 'save' games

var savedGames map[uuid.UUID]game
var currGame game;
var menu *Graph

func main() {
    rand.Seed(time.Now().UnixNano()) // use this to be able to create random numbers

	menuLevel := 1
	// Menu levels correspond to the Menu Graph below
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
					//saveGame(currGame)
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

	//check that the menu is built correctly
	/*
	for _, item := range menu.Vertices {
		fmt.Println("Item", item.Id, "with text", item.Text, "is connected to", menu.GetNeighbours(item.Id))
	}
	*/


//this works 100%

/*
	// You need a call to this if you want to create Random numbers. 
	// Not entirely sure why - read up on it! 
	// And is it enough to call it just once for the entire programme? Or whenever you need something new random?
    rand.Seed(time.Now().UnixNano())
	//playerIsPlaying := true

	var menuLevel int;

	for menuLevel != 9 {
		switch menuLevel {
        case 0:
			fmt.Println("MAIN MENU")
			fmt.Println("1. Start a New Game -- 2. Continue Saved Game -- 3. Exit")
        case 1:
            fmt.Println("GAME MENU:")
            fmt.Println("1. Place a number -- 2. Undo last move -- 3. Print board -- 4. Save game -- 5. Return to main menu")
		}

		fmt.Print("Enter selection: ")
		var userInput int
		fmt.Scan(&userInput)

		switch menuLevel {
		case 0:
			if userInput == 1 {
				currGame = startNewGame()
				menuLevel = 1
			} else if userInput == 2 {
				savedGame, err := loadSavedGame()
				if err != nil {
					fmt.Println(err)
				} else {
					currGame = *savedGame
					menuLevel = 1
				}
			}else if userInput == 3 {
				menuLevel = 9
			} else {
				fmt.Println("Invalid selection.")
			}
		case 1:
			if userInput == 1 {
				currGame.makeAMove()
				fmt.Println("===================")
				currGame.printBoard()
				fmt.Println("===================")
			} else if userInput == 2 {
				fmt.Println("Undoing a move here")
				currGame.undoLastMove()
			} else if userInput == 3 {
				currGame.printBoard()
			} else if userInput == 4 {
				saveGame(currGame)
			} else if userInput == 5 {
				menuLevel = 0  //return to main menu
			} else {
				fmt.Println("Invalid selection.")
			}
		}
	}

	*/