package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"strings"
	"time"
)

func newGame(n string, s ...int) (*Game, error) {
	rand.Seed(time.Now().UnixNano())
	// moves := [2]bool{true, false}

	size := 3

	if len(s) > 0 {
		if s[0] > 10 {
			return nil, fmt.Errorf("board size is too big: %d", s)
		}

		size = s[0]
	}

	if strings.ToLower(n) == strings.ToLower(BotName) {
		return nil, fmt.Errorf("username is not valid")
	}

	bSize := size * size
	board := make([]int, bSize)

	for i := 0; i < bSize; i++ {
		board[i] = 0
	}

	return &Game{size, board, n, true, false, ""}, nil // moves[rand.Intn(len(moves))], false, ""}, nil
}

func (g *Game) print(w io.Writer) error {
	for i := 0; i < g.Size; i++ {
		t := "|"

		for j := 0; j < g.Size; j++ {
			// 0 - Empty
			// 1 - Bot
			// 2 - Player

			var move string

			if g.Board[i*g.Size+j] == 0 {
				move = "   |"
			} else if g.Board[i*g.Size+j] == 1 {
				move = " B |"
			} else if g.Board[i*g.Size+j] == 2 {
				move = " P |"
			}

			t += move
		}

		fmt.Fprintf(w, "%s\n", t)
	}

	return nil
}

// Move is the general player move method
func (g *Game) Move(i, j int) (string, error) {
	if !g.PlayerMoves {
		return "player cannot move\n", errors.New("cannot move")
	}

	if i < 0 || i >= g.Size || j < 0 || j >= g.Size {
		return fmt.Sprintf("invalid move, out of range (%d, %d)\n", i, j), errors.New("out of range")
	}

	if g.Board[i*g.Size+j] != 0 {
		return fmt.Sprintf("invalid move to position (%d, %d)\n", i, j), errors.New("invalid position")
	}

	g.Board[i*g.Size+j] = 2
	return "", nil
}
