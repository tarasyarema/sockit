package main

// BotName is the bot name
const BotName = "Sockit"

// Game is main game struct
type Game struct {
	Size        int
	Board       []int
	Player      string
	PlayerMoves bool
	Finished    bool
	Winner      string
}
