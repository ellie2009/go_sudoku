package main

import (
	"errors"
)

type MoveStack []Move 

type Move struct {
	move [3] int
}

func (ms *MoveStack) AddMove (move Move) {
	*ms = append(*ms, move)
}

func (ms *MoveStack) RemoveMove() (*Move, error) {
	//check if Stack is empty
	if len(*ms) == 0 {
		return  nil, errors.New("no move to undo")
	} else {
	//get last move
		lastMoveIndex := len(*ms) - 1 // Get the index of the top move
		lastMove := (*ms)[lastMoveIndex] // Get the move using the index
		*ms = (*ms)[:lastMoveIndex] // Remove it from the stack
		return &lastMove, nil
	}
}