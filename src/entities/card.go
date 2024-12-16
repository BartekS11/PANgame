package entities

type Card struct {
	*Sprite
	suit      int8
	cardColor int8
	isFace    int8
}
