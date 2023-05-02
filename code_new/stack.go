package main

import (
	"errors"
)

type Stack []string

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}


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
		lastMoveIndex := len(*ms) - 1 // Get the index of the top most element.
		lastMove := (*ms)[lastMoveIndex] // Index into the slice and obtain the element.
		*ms = (*ms)[:lastMoveIndex] // Remove it from the stack by slicing it off.
//		fmt.Println(lastMove)
		return &lastMove, nil
	}
}