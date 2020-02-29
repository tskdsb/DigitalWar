package utils

import (
	"testing"
)

func TestGCD(t *testing.T) {
	list := []struct {
		a, b, c int
	}{
		{105, 252, 21},
		{120, 60, 60},
	}

	for _, l := range list {
		a := GCD(l.a, l.b)
		if a != l.c {
			t.Errorf("want %d, but got %d", l.c, a)
		}
	}

}
