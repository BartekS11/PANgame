package entities

import (
	"github.com/BartekS11/PANgame/constants"
	"github.com/hajimehoshi/ebiten/v2"
)

// type Card struct {
// 	*Sprite
// 	suit      int8
// 	cardColor int8
// 	isFace    int8
// }

var (
	gridCellWidth  = 100.0
	gridCellHeight = 100.0
)

type Card struct {
	RectX, RectY float64
	RectImage    *ebiten.Image
	IsDragging   bool
	color        constants.CardColor
	suite        constants.CardSuits
	val          constants.CardType
}

func (g *Card) Update() error {
	mouseX, mouseY := ebiten.CursorPosition()
	var floatmouseX float64 = float64(mouseX)
	var floatmouseY float64 = float64(mouseY)
	mousePressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)

	if mousePressed {
		if !g.IsDragging {
			if floatmouseX >= g.RectX && floatmouseX <= g.RectX+float64(g.RectImage.Bounds().Dx()) &&
				floatmouseY >= g.RectY && floatmouseY <= g.RectY+float64(g.RectImage.Bounds().Dy()) {
				g.IsDragging = true
			}
		}
		if g.IsDragging {
			g.RectX = floatmouseX - float64(g.RectImage.Bounds().Dx())/2
			g.RectY = floatmouseY - float64(g.RectImage.Bounds().Dy())/2
		}
	} else {
		g.IsDragging = false
	}

	if !g.IsDragging {
		g.RectX = float64(int((g.RectX+gridCellWidth/2)/gridCellWidth) * int(gridCellWidth))
		g.RectY = float64(int((g.RectY+gridCellHeight/2)/gridCellHeight) * int(gridCellHeight))
	}

	return nil
}
