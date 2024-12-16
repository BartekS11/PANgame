package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	ScreenWidth  = 1280
	ScreenHeight = 720
	ScreenSize   = ScreenWidth * ScreenHeight
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("FPS: %0.1f", ebiten.ActualFPS()), 100, 1)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("TPS: %0.1f", ebiten.ActualTPS()), 200, 1)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
