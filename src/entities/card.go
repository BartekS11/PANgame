package entities

import (
	"image/color"

	"github.com/BartekS11/PANgame/constants"
	"github.com/hajimehoshi/ebiten/v2"
)

// var (
// 	CardSizeX   = (822.0 / 5.0)
// 	CardSizeY   = (1122.0 / 5.0)
// 	IsDebugMode = false
// )

type Card struct {
	RectX, RectY float64
	RectImage    *ebiten.Image
	IsDragging   bool
	sprite       *Sprite
	color        constants.CardColor
	suite        constants.CardSuits
	val          constants.CardType
}

func (c *Card) Update() error {

	// Handle grid snapping for card
	mouseX, mouseY := ebiten.CursorPosition()
	var floatmouseX float64 = float64(mouseX)
	var floatmouseY float64 = float64(mouseY)
	mousePressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)

	if mousePressed {
		if !c.IsDragging {
			if floatmouseX >= c.RectX && floatmouseX <= c.RectX+float64(c.RectImage.Bounds().Dx()) &&
				floatmouseY >= c.RectY && floatmouseY <= c.RectY+float64(c.RectImage.Bounds().Dy()) {
				c.IsDragging = true
			}
		}
		if c.IsDragging {
			c.RectX = floatmouseX - float64(c.RectImage.Bounds().Dx())/2
			c.RectY = floatmouseY - float64(c.RectImage.Bounds().Dy())/2
		}
	} else {
		c.IsDragging = false
	}

	if !c.IsDragging {
		c.RectX = float64(int((c.RectX+constants.GridCellWidth/2)/constants.GridCellWidth) * int(constants.GridCellWidth))
		c.RectY = float64(int((c.RectY+constants.GridCellHeight/2)/constants.GridCellHeight) * int(constants.GridCellHeight))
	}

	return nil
}

func (c *Card) Draw(screen *ebiten.Image) {
	// Draw Card
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(c.RectX, c.RectY)
	screen.DrawImage(c.RectImage, op)
}

func CreateHand() []*Card {
	// Create hand
	var hand []*Card
	rectImg := ebiten.NewImage(int(80+1), int(60))
	rectImg.Fill(color.RGBA{0, 0, 0, 255})
	for card := range 5 {
		hand = append(hand, &Card{
			RectX:      float64(rectImg.Bounds().Dx()),
			RectY:      float64(rectImg.Bounds().Dy() + card),
			RectImage:  rectImg,
			IsDragging: false})
	}

	return hand
}
