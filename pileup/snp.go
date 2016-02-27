package pileup

import (
	"math"
)

type SNP struct {
	Reference string
	Position  int
	RefBase   byte
	Bases     []byte
	Quals     []byte
}

func (s *SNP) Pi() float64 {
	a, t, g, c := 0, 0, 0, 0
	for i := 0; i < len(s.Bases); i++ {
		switch s.Bases[i] {
		case 'A', 'a':
			a++
			break
		case 'T', 't':
			t++
			break
		case 'G', 'g':
			g++
			break
		case 'C', 'c':
			c++
			break
		}
	}

	cross := a*t + a*g + a*c + t*g + t*c + g*c
	total := a + t + g + c
	if total < 2 {
		return math.NaN()
	} else {
		return float64(cross) / float64(total*(total-1)/2)
	}
}
