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
	CardSizeX   = (822.0 / 5.0)
	CardSizeY   = (1122.0 / 5.0)
	IsDebugMode = false
)

type Game struct {
	rectImage                     *ebiten.Image
	rectX, rectY                  float64
	gridCols, gridRows            int
	gridCellWidth, gridCellHeight float64
	cardEntities                  *entities.Card
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyF4) {
		IsDebugMode = !IsDebugMode
	}

	g.cardEntities.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for row := range g.gridRows {
		for col := range g.gridCols {
			x := float64(col) * g.gridCellWidth
			y := float64(row) * g.gridCellHeight
			rect := ebiten.NewImage(int(g.gridCellWidth)-1, int(g.gridCellHeight)-1)
			rect.Fill(color.RGBA{200, 200, 200, 255})
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(x, y)

			screen.DrawImage(rect, op)
		}
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.cardEntities.RectX, g.cardEntities.RectY)
	screen.DrawImage(g.cardEntities.RectImage, op)
	if IsDebugMode {
		debugPrints(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return constants.ScreenWidth, constants.ScreenHeight
}

func NewGame() (*Game, error) {
	width, height := int(CardSizeX), int(CardSizeY)
	rectImage := ebiten.NewImage(width, height)
	rectImage.Fill(color.RGBA{255, 0, 0, 255})
	g := &Game{
		rectImage:      rectImage,
		rectX:          CardSizeX,
		rectY:          CardSizeY,
		gridCols:       4,
		gridRows:       3,
		gridCellWidth:  100.0,
		gridCellHeight: 100.0,
		cardEntities:   nil,
	}
	g.cardEntities = &entities.Card{
		RectX:      CardSizeX,
		RectY:      CardSizeY,
		RectImage:  rectImage,
		IsDragging: false,
	}

	return g, nil
}

func debugPrints(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("FPS: %0.1f", ebiten.ActualFPS()), 100, 1)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("TPS: %0.1f", ebiten.ActualTPS()), 200, 1)
}
