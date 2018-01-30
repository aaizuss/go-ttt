package main

import (
	"fmt"
	"github.com/aaizuss/tictactoe/game"
)

func main() {
	fmt.Printf("Hello, world\n")
	game := game.New()
	game.Play()
}
