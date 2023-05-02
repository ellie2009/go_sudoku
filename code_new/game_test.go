package main

import (
	"testing"
)

func BenchmarkCreateBoard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g := game {playerName: "mimi", difficulty: "hard", isComplete: false} 
        g.createBoard()
    }
}

func BenchmarkCreateBoardWithPrefilledDiagonalBoxes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g := game {playerName: "mimi", difficulty: "hard", isComplete: false} 
        g.createBoardWithPrefilledDiagonalBoxes()
    }
}