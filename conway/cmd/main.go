package main

import "conway/pkg/game"

func main() {
	game := game.NewGame(32, 64)
	game.Run()
}
