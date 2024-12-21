package main

import (
	"fmt"
	"image/color"

	"github.com/BartekS11/PANgame/constants"
	"github.com/BartekS11/PANgame/entities"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	// Keep real life card size reference
	// CardSizeX   = (822.0 / 5.0)
	// CardSizeY   = (1122.0 / 5.0)
	IsDebugMode = false
)

type Game struct {
	cardEntities *entities.Card
	gridEntity   *entities.Grid
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyF4) {
		IsDebugMode = !IsDebugMode
	}

	g.gridEntity.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{120, 180, 255, 255})

	g.gridEntity.Draw(screen)

	if IsDebugMode {
		debugPrints(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return constants.ScreenWidth, constants.ScreenHeight
}

func NewGame() (*Game, error) {
	grid, err := entities.LoadGrid()
	if err != nil {
		panic(err)
	}
	g := &Game{
		cardEntities: nil,
		gridEntity:   grid,
	}

	return g, nil
}

func debugPrints(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("FPS: %0.1f", ebiten.ActualFPS()), 100, 1)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("TPS: %0.1f", ebiten.ActualTPS()), 200, 1)
}
