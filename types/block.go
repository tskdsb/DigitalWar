package types

import (
	"image/color"
	"math/rand"

	"github.com/tskdsb/DigitalWar/utils"
)

type Location struct {
	x int
	y int
}

type Block struct {
	Location *Location
	Belong   *Force
	Color    color.Color
	Base     *Fraction
	Growth   *Fraction
}

func NewBlock(x, y int) *Block {
	return &Block{
		Location: &Location{x, y},
		Belong:   nil,
		// Color:    color.White,
		Color:  color.Black,
		Base:   NewFraction2(uint(utils.Abs(rand.Intn(10))), 10),
		Growth: NewFraction2(uint(utils.Abs(rand.Intn(10))), 10),
	}
}

func (b *Block) Grow() {
	b.Base.Add(*b.Growth)
}

func (b *Block) Submission(force *Force) {
	b.Belong = force
	b.Color = force.Color
}
