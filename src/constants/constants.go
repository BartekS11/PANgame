package constants

// Screen
const (
	ScreenWidth  = 1280
	ScreenHeight = 720
	ScreenSize   = ScreenWidth * ScreenHeight
)

// Card suits
type CardSuits int8

const (
	Club CardSuits = iota
	Hearth
	Spade
	Diamond
)
