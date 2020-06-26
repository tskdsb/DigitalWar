package types

import (
	"image/color"
)

type Force struct {
	Color color.Color
}

type Team struct {
	member *Fraction
	block  *Block
}
