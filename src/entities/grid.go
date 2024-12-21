package entities

import (
	"image/color"

	"github.com/BartekS11/PANgame/constants"
	"github.com/hajimehoshi/ebiten/v2"
)

type Grid struct {
	gridCols, gridRows                 int
	gridCellWidth, gridCellHeight      float64
	Player1HandCards, player2HandCards []*Card
	grid                               [][]*ebiten.Image
}

func LoadGrid() (*Grid, error) {
	gridCols, gridRows := 4, 3
	gridCellWidth, gridCellHeight := constants.GridCellWidth, constants.GridCellHeight

	// Create empty grid
	grid := make([][]*ebiten.Image, gridRows)
	for i := range grid {
		grid[i] = make([]*ebiten.Image, gridCols)
	}

	// Fill hands with 5 (hardcoded for now) cards
	player1Hand := CreateHand()
	player2Hand := CreateHand()

	// Load 1st player hand
	for i, card := range player1Hand {
		card.RectX = float64(i)*150 + 20 // Spacing between cards
		card.RectY = 80.0                // Top of the screen
	}

	// Load 2nd player hand
	for i, card := range player2Hand {
		card.RectX = float64(i)*150 + 20 // Spacing between cards
		card.RectY = 400.0               // Bottom of the screen
	}

	return &Grid{
		gridCols:         gridCols,
		gridRows:         gridRows,
		gridCellWidth:    gridCellWidth,
		gridCellHeight:   gridCellHeight,
		Player1HandCards: player1Hand,
		player2HandCards: player2Hand,
		grid:             grid,
	}, nil
}

func (g *Grid) Draw(screen *ebiten.Image) {
	// Draw the grid
	for row := range g.gridRows {
		for col := range g.gridCols {
			x := float64(col)*g.gridCellWidth + (constants.ScreenHeight / 2) // Center grid horizontally
			y := float64(row)*g.gridCellHeight + 200                         // Center grid vertically
			cell := ebiten.NewImage(int(g.gridCellWidth), int(g.gridCellHeight))
			cell.Fill(color.RGBA{200, 255, 200, 255}) // Light gray grid cells
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(x, y)
			screen.DrawImage(cell, op)

			// Draw cards on the grid if present
			if g.grid[row][col] != nil {
				cardOp := &ebiten.DrawImageOptions{}
				cardOp.GeoM.Translate(x+10, y+10)
				screen.DrawImage(g.grid[row][col], cardOp)
			}
		}
	}

	// Draw 1st player hand
	for _, card := range g.Player1HandCards {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(card.RectX, card.RectY)
		screen.DrawImage(card.RectImage, op)
	}

	// Draw 2nd player hand
	for _, card := range g.player2HandCards {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(card.RectX, card.RectY)
		screen.DrawImage(card.RectImage, op)
	}
}

func (g *Grid) Update() {
	// Update 1st player hand
	for _, card := range g.Player1HandCards {
		card.Update()
	}

	// Update 2nd player hand
	for _, card := range g.player2HandCards {
		card.Update()
	}
}
