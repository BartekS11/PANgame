package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	CardSizeX   = (822.0 / 15.0)
	CardSizeY   = (1122.0 / 15.0)
	IsDebugMode = false
)

type Game struct {
	rectImage                     *ebiten.Image
	rectX, rectY                  float64 // Rectangle position
	gridCols, gridRows            int
	gridCellWidth, gridCellHeight float64
	isDragging                    bool
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyF4) {
		IsDebugMode = !IsDebugMode
	}
	mouseX, mouseY := ebiten.CursorPosition()
	mousePressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)

	if mousePressed {
		if !g.isDragging {
			// Check if the click is inside the rectangle to start dragging
			if float64(mouseX) >= g.rectX && float64(mouseX) <= g.rectX+float64(g.rectImage.Bounds().Dx()) &&
				float64(mouseY) >= g.rectY && float64(mouseY) <= g.rectY+float64(g.rectImage.Bounds().Dy()) {
				g.isDragging = true
			}
		}
		if g.isDragging {
			// Update rectangle position to follow the mouse
			g.rectX = float64(mouseX) - float64(g.rectImage.Bounds().Dx())/2
			g.rectY = float64(mouseY) - float64(g.rectImage.Bounds().Dy())/2
		}
	} else {
		// Stop dragging when the mouse is released
		g.isDragging = false
	}

	// Snap rectangle to the nearest grid cell when not dragging
	if !g.isDragging {
		g.rectX = float64(int((g.rectX+g.gridCellWidth/2)/g.gridCellWidth) * int(g.gridCellWidth))
		g.rectY = float64(int((g.rectY+g.gridCellHeight/2)/g.gridCellHeight) * int(g.gridCellHeight))
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for row := 0; row < g.gridRows; row++ {
		for col := 0; col < g.gridCols; col++ {
			x := float64(col) * g.gridCellWidth
			y := float64(row) * g.gridCellHeight
			rect := ebiten.NewImage(int(g.gridCellWidth), int(g.gridCellHeight))
			rect.Fill(color.RGBA{200, 200, 200, 255}) // Light gray for grid cells
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(x, y)
			screen.DrawImage(rect, op)
		}
	}

	// Draw the rectangle image onto the screen
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.rectX, g.rectY) // Position the rectangle
	screen.DrawImage(g.rectImage, op)
	if IsDebugMode {
		debugPrints(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func NewGame() (*Game, error) {
	width, height := int(CardSizeX), int(CardSizeY)
	rectImage := ebiten.NewImage(width, height)
	rectImage.Fill(color.RGBA{255, 0, 0, 255})
	g := &Game{
		rectImage:      rectImage,
		rectX:          CardSizeX,
		rectY:          CardSizeY,
		gridCols:       4, // Number of columns in the grid
		gridRows:       3, // Number of rows in the grid
		gridCellWidth:  20.0,
		gridCellHeight: 20.0,
		isDragging:     false,
	}
	return g, nil
}

func debugPrints(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("FPS: %0.1f", ebiten.ActualFPS()), 100, 1)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("TPS: %0.1f", ebiten.ActualTPS()), 200, 1)
}
