package main

import "conway/pkg/game"

func main() {
	game := game.NewGame(6, 64)
	game.Run()
}
