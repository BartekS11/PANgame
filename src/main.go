package main

import (
	"github.com/BartekS11/PANgame/constants"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(constants.ScreenWidth, constants.ScreenHeight)
	ebiten.SetTPS(60)
	ebiten.SetWindowTitle("PAN(game) | dev")

	g, err := NewGame()
	if err != nil {
		panic(err)
	}

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
