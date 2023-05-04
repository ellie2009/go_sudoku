package main

import (
	"fmt"
	"math/rand"

	"github.com/google/uuid"
)

type game struct {
	id uuid.UUID
	playerName string
	isComplete bool
	difficulty string
	gameBoard [9][9] int
	myMoves MoveStack // [row, col, num]
}

func (g *game) createBoard() {
	nums := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    rand.Shuffle(len(nums), func(i, j int) { nums[i], nums[j] = nums[j], nums[i] })

	g.gameBoard[0] = nums 

	// can start from 1 because first row is already filled in as per above.
	g.solveBoard(1,0)

   	g.removeCells()
}

func (g *game) createBoardWithPrefilledDiagonalBoxes() {
	// fill in the three diagonal boxes. No validation is required because they are independent from each other.

	for i := 0; i < 9; i += 3 {
		nums := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		rand.Shuffle(len(nums), func(i, j int) { nums[i], nums[j] = nums[j], nums[i] })
		numsIndex := 0
		for row := i; row < i+3; row++ {
			for col := i; col < i+3; col ++{
				g.gameBoard[row][col] = nums[numsIndex]
				numsIndex++
			}
		}	
	}

	g.solveBoard(0,0)

	g.removeCells()
}

func (g *game) solveBoard(row int, col int) bool {
	if row == 9 {
		return true
	}
	if g.gameBoard[row][col] != 0 {
		return g.solveBoard(nextCell(row, col))
	} else {
		for num := 1; num <= 9; num++ {
			if g.isValidMove(row, col, num) {
				g.gameBoard[row][col] = num
				if g.solveBoard(nextCell(row, col)){
					return true
				}
				g.gameBoard[row][col] = 0
			}
		}
		return false
	}
}

func (g *game) placeNumber(row int, col int, num int) {
	if g.isValidMove(row, col, num) {
		g.gameBoard[row][col] = num
		m := [3]int{row, col, num}
		myMove := Move {move: m}
		g.myMoves = append(g.myMoves, myMove)
	} else {
		fmt.Println("INVALID MOVE.")
	}
}

func (g *game) isValidMove(row int, col int, num int) bool {
	return !g.isAlreadyInRow(row, num) && !g.isAlreadyInCol(col, num) && !g.isAlreadyInBox(row, col, num)
}

func (g *game) isAlreadyInRow(row int, num int) bool {
    for i := 0; i < 9; i++ {
        if g.gameBoard[row][i] == num {
            return true
        }
    }
    return false
}

func (g *game) isAlreadyInCol(col int, num int) bool {
    for row := 0; row < 9; row++ {
        if g.gameBoard[row][col] == num {
            return true
        }
    }
    return false
}

func (g *game) isAlreadyInBox(row int, col int, num int) bool {
    beginBoxRow := (row/3) * 3
    beginBoxCol := (col/3) *3
    for i := beginBoxRow; i < beginBoxRow + 3; i++ {
        for j := beginBoxCol; j < beginBoxCol + 3; j++ {
            if g.gameBoard[i][j] == num {
                return true
            }
        }
    }
    return false
}

func (g *game) removeCells() {
	numEmptyCells := 0

	if g.difficulty == "easy" {
		numEmptyCells = 2
	} else if g.difficulty == "medium" {
		numEmptyCells = 40
	} else {
		numEmptyCells = 80
	}

	emptyCellCount := 0
    for emptyCellCount < numEmptyCells {
		row := rand.Intn(9)
		col := rand.Intn(9)
		if g.gameBoard[row][col] != 0 {
		    g.gameBoard[row][col] = 0
		    emptyCellCount++
		}
	}
}

func nextCell(row int, col int) (int, int) {
    nextRow := row
    nextCol := (col +1) % 9 
    if nextCol == 0 {
        nextRow = row + 1
    }
    return nextRow, nextCol
}

func (g *game) undoLastMove() {
	// remove the latest move from the stack of moves. Format is [row, col, num]
	lastMove, err := g.myMoves.RemoveMove()
	if err != nil {
		fmt.Println(err)
		return
	}
	// set the gameboard cell to 0 
	g.gameBoard[lastMove.move[0]][lastMove.move[1]] = 0
}

func (g *game) isGameComplete() (bool) {

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col ++{
			if g.gameBoard[row][col] == 0 {
				return false
			}
		}
	}

	return true
}