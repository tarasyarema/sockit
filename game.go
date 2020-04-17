package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

func newGame(n string) (*Game, error) {
	rand.Seed(time.Now().UnixNano())
	moves := [2]bool{true, false}

	s := os.Getenv("SIZE")

	if s == "" {
		s = "3"
	}

	size, _ := strconv.ParseInt(s, 10, 64)

	if strings.ToLower(n) == strings.ToLower(BotName) {
		return nil, fmt.Errorf("username is not valid")
	}

	bSize := int(size * size)
	board := make([]int, bSize)

	for i := 0; i < bSize; i++ {
		board[i] = 0
	}

	flag := os.Getenv("FLAG")

	if flag == "" {
		flag = "FLAG"
	}

	game := &Game{
		Flag:        flag,
		Board:       board,
		Size:        int(size),
		Player:      n,
		Finished:    false,
		PlayerMoves: moves[rand.Intn(len(moves))],
		Won:         false,
	}

	log.Infof("game created: %v", game)

	return game, nil
}

// Print prints the current game state
func (g *Game) Print(w io.Writer) {
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

	g.PlayerMoves = !g.PlayerMoves
	g.Board[i*g.Size+j] = 2
	return "", nil
}

func checkBoard(t, s int, a []int) bool {
	// check main diagonal
	diag := true
	for i := 0; i < s; i++ {
		if a[i*s+i] != t {
			diag = false
			break
		}
	}

	if diag {
		return true
	}

	// check main antidiagonal
	diag = true
	for i := 0; i < s; i++ {
		if a[i*s+s-i-1] != t {
			diag = false
			break
		}
	}

	if diag {
		return true
	}

	for i := 0; i < s; i++ {
		local := true

		// check horizontal
		for j := 0; j < s; j++ {
			if a[i*s+j] != t {
				local = false
				break
			}
		}

		if local {
			return true
		}

		local = true

		// check vertical
		for j := 0; j < s; j++ {
			if a[j*s+i] != t {
				local = false
				break
			}
		}

		if local {
			return true
		}
	}

	return false
}

func (g *Game) check() {
	var wg sync.WaitGroup
	bot, ply := false, false

	wg.Add(2)

	// bot check (1)
	go func() {
		defer wg.Done()
		if checkBoard(1, g.Size, g.Board) {
			bot = true
		}
	}()

	// player check (2)
	go func() {
		defer wg.Done()
		if checkBoard(2, g.Size, g.Board) {
			ply = true
		}
	}()

	wg.Wait()

	if !bot && !ply {
		return
	}

	g.Finished = true

	if ply && !bot {
		g.Won = true
	}
}

func (g *Game) counter() error {
	var moves []int

	for i := range g.Board {
		if g.Board[i] == 0 {
			moves = append(moves, i)
		}
	}

	if len(moves) == 0 {
		return fmt.Errorf("no moves available")
	}

	g.PlayerMoves = !g.PlayerMoves
	g.Board[moves[rand.Intn(len(moves))]] = 1
	return nil
}
