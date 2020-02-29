package types

type Location struct {
	x int
	y int
}

type Block struct {
	Location *Location
	Belong   *Force
	Base     *Fraction
	Growth   *Fraction
}

func (b *Block) Grow() {
	b.Base.Add(*b.Growth)
}

func (b *Block) Submission(force *Force) {
	b.Belong = force
}
