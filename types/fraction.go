package types

import (
	"github.com/tskdsb/DigitalWar/utils"
)

type Fraction struct {
	fenzi int
	fenmu int
}

func NewFraction(a int) *Fraction {
	return &Fraction{
		fenzi: a,
		fenmu: 1,
	}
}

func (f *Fraction) ToFloat() float64 {
	return float64(f.fenzi) / float64(f.fenmu)
}

func (f *Fraction) Add(f2 Fraction) {
	f.fenmu = f.fenmu * f2.fenmu
	f.fenzi = f.fenzi*f2.fenmu + f2.fenzi*f.fenmu
}

func (f *Fraction) Sub(f2 Fraction) {
	f.fenmu = f.fenmu * f2.fenmu
	f.fenzi = f.fenzi*f2.fenmu - f2.fenzi*f.fenmu
}

func (f *Fraction) YueFen() {
	gcd := utils.GCD(f.fenzi, f.fenmu)
	f.fenzi = f.fenzi / gcd
	f.fenmu = f.fenmu / gcd
}
