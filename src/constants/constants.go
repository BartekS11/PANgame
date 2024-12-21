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

// Card color
type CardColor int8

const (
	Red CardColor = iota
	Black
)

// Card type
type CardType int8

const (
	Num9 CardType = iota
	Num10

	Jack
	Queen
	King
	Ace
)

// Grid
const (
	GridCellWidth  = 80.0
	GridCellHeight = 80.0
)
